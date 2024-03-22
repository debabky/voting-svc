package pg

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/debabky/voting-svc/internal/data"
	"gitlab.com/distributed_lab/kit/pgdb"
)

const votingOptionsTableName = "voting_options"

var (
	votingOptionsSelector = sq.Select("*").From(votingOptionsTableName)
	votingOptionsUpdate   = sq.Update(votingOptionsTableName)
)

func NewVotingOptionsQ(db *pgdb.DB) data.VotingOptionsQ {
	return &votingOptionsQ{
		db:  db,
		sql: votingOptionsSelector,
		upd: votingOptionsUpdate,
	}
}

type votingOptionsQ struct {
	db  *pgdb.DB
	sql sq.SelectBuilder
	upd sq.UpdateBuilder
}

func (q *votingOptionsQ) New() data.VotingOptionsQ {
	return NewVotingOptionsQ(q.db.Clone())
}

func (q *votingOptionsQ) FilterBy(column string, value any) data.VotingOptionsQ {
	q.sql = q.sql.Where(sq.Eq{column: value})
	return q
}

func (q *votingOptionsQ) Select() ([]data.VotingOption, error) {
	var result []data.VotingOption
	err := q.db.Select(&result, q.sql)
	return result, err
}
