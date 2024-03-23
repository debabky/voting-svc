package data

import (
	"github.com/google/uuid"
	"time"
)

type VotingsQ interface {
	New() VotingsQ
	FilterBy(column string, value any) VotingsQ
	Select() ([]Voting, error)
	Get() (*Voting, error)
}

type Voting struct {
	ID          uuid.UUID  `db:"id" structs:"-"`
	Name        string     `db:"name" structs:"name"`
	Description string     `db:"description" structs:"description"`
	Type        VotingType `db:"voting_type" structs:"voting_type"`
	CreatedAt   time.Time  `db:"created_at" structs:"-"`
	ActiveUntil time.Time  `db:"active_until" structs:"active_until"`
}
