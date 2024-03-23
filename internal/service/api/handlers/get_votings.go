package handlers

import (
	"github.com/debabky/voting-svc/resources"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"net/http"
)

func GetVotings(w http.ResponseWriter, r *http.Request) {
	votings, err := MasterQ(r).VotingsQ().Select()
	if err != nil {
		Log(r).WithError(err).Error("failed to select votings")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	responseData := make([]resources.Voting, len(votings))
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
	}

	response := resources.VotingListResponse{
		Data: responseData,
	}

	ape.Render(w, response)
}
