package requests

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/distributed_lab/urlval"
)

type GetVotingRequest struct {
	ID             uuid.UUID
	IncludeOptions bool    `include:"options"`
	Nullifier      *string `url:"nullifier"`
}

func NewGetVotingRequest(r *http.Request) (GetVotingRequest, error) {
	var request GetVotingRequest

	rawID := chi.URLParam(r, "id")
	id, err := uuid.Parse(rawID)
	if err != nil {
		return GetVotingRequest{}, errors.Wrap(err, "failed to parse UUID")
	}
	request.ID = id

	if err := urlval.Decode(r.URL.Query(), &request); err != nil {
		return GetVotingRequest{}, errors.Wrap(err, "failed to decode url")
	}
	return request, nil
}
