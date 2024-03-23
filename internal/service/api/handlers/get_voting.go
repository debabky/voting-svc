package handlers

import (
	"github.com/debabky/voting-svc/internal/service/api/requests"
	"github.com/debabky/voting-svc/resources"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"net/http"
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

	included := resources.Included{}
	if req.IncludeOptions {
		options, err := MasterQ(r).VotingOptionsQ().New().
			FilterBy("voting_id", voting.ID).
			Select()
		if err != nil {
			Log(r).WithError(err).Error("failed to select voting options")
			ape.RenderErr(w, problems.InternalError())
			return
		}
		for _, option := range options {
			included.Add(&resources.VotingOption{
				Key: resources.Key{
					ID:   option.Name,
					Type: resources.VOTING_OPTIONS,
				},
				Attributes: resources.VotingOptionAttributes{
					Name:        option.Name,
					VotingId:    voting.ID,
					Description: option.Description,
				},
			})
		}

		response.Included = included
	}

	ape.Render(w, response)
}
