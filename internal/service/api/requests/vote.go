package requests

import (
	"encoding/json"
	snark "github.com/iden3/go-rapidsnark/types"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"net/http"
)

type VoteRequestData struct {
	InternalPublicKey string `json:"internal_public_key"`
	Signature         struct {
		S string `json:"s"`
		N string `json:"n"`
	} `json:"signature"`
	Proof     snark.ZKProof `json:"proof"`
	Timestamp int64         `json:"timestamp"`
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
