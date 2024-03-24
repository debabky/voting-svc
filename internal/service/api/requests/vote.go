package requests

import (
	"encoding/json"
	snark "github.com/iden3/go-rapidsnark/types"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"net/http"
)

type VoteRequestData struct {
	Candidates []int64       `json:"candidates"`
	Proof      snark.ZKProof `json:"proof"`
	Nullifier  string        `json:"nullifier"`
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
