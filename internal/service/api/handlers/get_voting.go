package handlers

import (
	"fmt"
	"net/http"

	"github.com/debabky/voting-svc/internal/data"
	"github.com/debabky/voting-svc/internal/service/api/requests"
	"github.com/debabky/voting-svc/resources"
	"github.com/google/uuid"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func GetVoting(w http.ResponseWriter, r *http.Request) {
	req, err := requests.NewGetVotingRequest(r)
	if err != nil {
		Log(r).WithError(err).Error("failed to get voting request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	voting, err := MasterQ(r).VotingsQ().New().
		FilterBy("id", req.ID).
		Get()
	if err != nil {
		Log(r).WithError(err).Error("failed to get voting")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	if voting == nil {
		ape.RenderErr(w, problems.NotFound())
		return
	}

	response := resources.VotingResponse{
		Data: resources.Voting{
			Key: resources.Key{
				ID:   voting.ID.String(),
				Type: resources.VOTINGS,
			},
			Attributes: resources.VotingAttributes{
				Name:        voting.Name,
				Description: voting.Description,
				CreatedAt:   voting.CreatedAt,
				ActiveUntil: voting.ActiveUntil,
			},
		},
	}

	if req.IncludeOptions {
		options, err := MasterQ(r).VotingOptionsQ().New().
			FilterBy("voting_id", voting.ID).
			Select()
		if err != nil {
			Log(r).WithError(err).Error("failed to select voting options")
			ape.RenderErr(w, problems.InternalError())
			return
		}

		optionsToAdd := make([]resources.VotingOption, 0)
		for _, option := range options {
			votes, err := MasterQ(r).VotesQ().New().
				FilterBy("voting_id", option.VotingID).
				FilterBy("voting_option", option.Name).
				Select()
			if err != nil {
				Log(r).WithError(err).Error("failed to get option votes")
				return
			}

			var ( // different vars are used for different voting types
				points   *int64
				votesNum *int64
			)

			switch voting.Type {
			case data.RankedVoting:
				res := calculateRankedVotingOptionPoints(int64(len(options)), votes)
				points = &res
			default:
				Log(r).Error(fmt.Sprintf("%d is not a valid voting type", voting.Type))
				ape.RenderErr(w, problems.InternalError())
				return
			}

			optionsToAdd = append(optionsToAdd, resources.VotingOption{
				Key: resources.Key{
					ID:   option.Name,
					Type: resources.VOTING_OPTIONS,
				},
				Attributes: resources.VotingOptionAttributes{
					Id:          option.ID,
					Name:        option.Name,
					VotingId:    voting.ID,
					Description: option.Description,
					Points:      points,
					Votes:       votesNum,
				},
			})
		}
		response.Data.Attributes.Options = &optionsToAdd
	}

	if req.Nullifier != nil && TokenClaims(r).Nullifier != "" && TokenClaims(r).VotingID != uuid.Nil {
		votes, err := MasterQ(r).VotesQ().New().
			FilterBy("nullifier", TokenClaims(r).Nullifier).
			FilterBy("voting_id", TokenClaims(r).VotingID).
			Select()
		if err != nil {
			Log(r).WithError(err).Error("failed to get votes")
			return
		}

		votesToAdd := make([]resources.Vote, 0)
		for _, vote := range votes {
			votesToAdd = append(votesToAdd, resources.Vote{
				VotingOption: vote.VotingOption,
				Rank:         vote.Rank,
			})
		}
		response.Data.Attributes.Votes = &votesToAdd
	}

	ape.Render(w, response)
}

func calculateRankedVotingOptionPoints(optionsAmount int64, votes []data.Vote) int64 {
	var points int64
	for _, vote := range votes {
		points += optionsAmount - *vote.Rank
	}
	return points
}
