package data

import (
	"github.com/google/uuid"
)

type VerificationRequestsQ interface {
	New() VerificationRequestsQ
	Insert(value VerificationRequest) error
	FilterBy(column string, value any) VerificationRequestsQ
	Get() (*VerificationRequest, error)
}

type VerificationRequest struct {
	ID        uuid.UUID `db:"id" structs:"id"`
	VotingID  uuid.UUID `db:"voting_id" structs:"voting_id"`
	Nullifier string    `db:"nullifier" structs:"nullifier"`
}
