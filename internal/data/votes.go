package data

import (
	"github.com/google/uuid"
)

type VotesQ interface {
	New() VotesQ
	FilterBy(column string, value any) VotesQ
	Count() (int64, error)
}

type Vote struct {
	VotingID     uuid.UUID `db:"voting_id" structs:"voting_id"`
	VotingOption string    `db:"voting_option" structs:"voting_option"`
	Nullifier    string    `db:"nullifier" structs:"nullifier"`
}
