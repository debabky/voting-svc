package pg

import (
	"github.com/debabky/voting-svc/internal/data"
	"gitlab.com/distributed_lab/kit/pgdb"
)

func NewMasterQ(db *pgdb.DB) data.MasterQ {
	return &masterQ{
		db: db.Clone(),
	}
}

type masterQ struct {
	db *pgdb.DB
}

func (m *masterQ) New() data.MasterQ {
	return NewMasterQ(m.db)
}

func (m *masterQ) Transaction(fn func(q data.MasterQ) error) error {
	return m.db.Transaction(func() error {
		return fn(m)
	})
}

func (m *masterQ) VotingsQ() data.VotingsQ {
	return NewVotingsQ(m.db)
}

func (m *masterQ) VotingOptionsQ() data.VotingOptionsQ {
	return NewVotingOptionsQ(m.db)
}

func (m *masterQ) RegistrationsQ() data.RegistrationsQ {
	return NewRegistrationsQ(m.db)
}

func (m *masterQ) VotesQ() data.VotesQ {
	return NewVotesQ(m.db)
}

func (m *masterQ) VerificationRequestsQ() data.VerificationRequestsQ {
	return NewVerificationRequestsQ(m.db)
}
