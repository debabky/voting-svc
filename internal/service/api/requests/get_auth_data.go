package requests

import (
	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"net/http"
)

type GetAuthDataRequest struct {
	ID uuid.UUID
}

func NewGetAuthDataRequest(r *http.Request) (GetVotingRequest, error) {
	var request GetVotingRequest

	rawID := chi.URLParam(r, "id")
	id, err := uuid.Parse(rawID)
	if err != nil {
		return GetVotingRequest{}, errors.Wrap(err, "failed to parse UUID")
	}
	request.ID = id

	return request, nil
}
