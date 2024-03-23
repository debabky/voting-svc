package requests

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type GetAuthDataRequest struct {
	ID uuid.UUID
}

func NewGetAuthDataRequest(r *http.Request) (GetAuthDataRequest, error) {
	var request GetAuthDataRequest

	rawID := chi.URLParam(r, "id")
	id, err := uuid.Parse(rawID)
	if err != nil {
		return GetAuthDataRequest{}, errors.Wrap(err, "failed to parse UUID")
	}
	request.ID = id

	return request, nil
}
