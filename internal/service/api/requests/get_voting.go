package requests

import (
	"net/http"

	"github.com/go-chi/chi"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/distributed_lab/urlval"
)

type GetVotingRequest struct {
	ID             string
	IncludeOptions bool    `include:"options"`
	Nullifier      *string `url:"nullifier"`
}

func NewGetVotingRequest(r *http.Request) (GetVotingRequest, error) {
	var request GetVotingRequest

	rawID := chi.URLParam(r, "id")
	request.ID = rawID

	if err := urlval.Decode(r.URL.Query(), &request); err != nil {
		return GetVotingRequest{}, errors.Wrap(err, "failed to decode url")
	}
	return request, nil
}
