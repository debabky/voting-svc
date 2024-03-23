package requests

import (
	"encoding/json"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"net/http"
)

type Vote struct {
	VotingOption string `json:"voting_option"`
	Rank         *int64 `json:"rank"`
}

type VoteRequestData struct {
	Votes []Vote `json:"votes"`
}

type VoteRequest struct {
	Data VoteRequestData `json:"data"`
}

func NewVoteRequest(r *http.Request) (VoteRequest, error) {
	var request VoteRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		return request, errors.Wrap(err, "failed to unmarshal")
	}

	return request, nil
}
