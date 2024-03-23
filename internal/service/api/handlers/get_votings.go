package handlers

import (
	"github.com/debabky/voting-svc/resources"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"net/http"
)

func GetVotings(w http.ResponseWriter, r *http.Request) {
	// TODO request with includes
	votings, err := MasterQ(r).VotingsQ().Select()
	if err != nil {
		Log(r).WithError(err).Error("failed to select votings")
		return
	}

	responseData := make([]resources.Voting, len(votings))
	included := resources.Included{}
	for i, voting := range votings {
		responseData[i] = resources.Voting{
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
		}

		if true { // FIXME
			options, err := MasterQ(r).VotingOptionsQ().New().
				FilterBy("voting_id", voting.ID).
				Select()
			if err != nil {
				Log(r).WithError(err).Error("failed to select voting options")
				ape.RenderErr(w, problems.InternalError())
				return
			}
			for _, option := range options {
				included.Add(resources.VotingOption{
					Attributes: resources.VotingOptionAttributes{
						Name:        option.Name,
						VotingId:    voting.ID,
						Description: option.Description,
					},
				}.GetKeyP())
			}
		}
	}

	response := resources.VotingListResponse{
		Data:     responseData,
		Included: included,
	}

	ape.Render(w, response)
}
