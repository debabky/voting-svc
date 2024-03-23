package data

import (
	"github.com/google/uuid"
)

type RegistrationsQ interface {
	New() RegistrationsQ
	Insert(value Registration) error
	FilterBy(column string, value any) RegistrationsQ
	Get() (*Registration, error)
}

type Registration struct {
	VotingID  uuid.UUID `db:"voting_id" structs:"voting_id"`
	Nullifier string    `db:"nullifier" structs:"nullifier"`
}
