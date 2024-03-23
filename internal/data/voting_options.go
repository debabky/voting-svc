package data

import (
	"github.com/google/uuid"
)

type VotingOptionsQ interface {
	New() VotingOptionsQ
	FilterBy(column string, value any) VotingOptionsQ
	Select() ([]VotingOption, error)
}

type VotingOption struct {
	Name        string    `db:"name" structs:"name"`
	VotingID    uuid.UUID `db:"voting_id" structs:"voting_id"`
	Description *string   `db:"description" structs:"description"`
}
