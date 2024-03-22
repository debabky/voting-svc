package data

import (
	"github.com/google/uuid"
	"time"
)

type VotingsQ interface {
	New() VotingsQ
	Select() ([]Voting, error)
}

type Voting struct {
	ID          uuid.UUID `db:"id" structs:"-"`
	Name        string    `db:"name" structs:"name"`
	Description string    `db:"description" structs:"description"`
	CreatedAt   time.Time `db:"created_at" structs:"-"`
}
