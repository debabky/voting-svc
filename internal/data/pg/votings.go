package pg

import (
	"database/sql"
	sq "github.com/Masterminds/squirrel"
	"github.com/debabky/voting-svc/internal/data"
	"gitlab.com/distributed_lab/kit/pgdb"
)

const votingsTableName = "votings"

var (
	votingsSelector = sq.Select("*").From(votingsTableName)
	votingsUpdate   = sq.Update(votingsTableName)
)

func NewVotingsQ(db *pgdb.DB) data.VotingsQ {
	return &votingsQ{
		db:  db,
		sql: votingsSelector,
		upd: votingsUpdate,
	}
}

type votingsQ struct {
	db  *pgdb.DB
	sql sq.SelectBuilder
	upd sq.UpdateBuilder
}

func (q *votingsQ) New() data.VotingsQ {
	return NewVotingsQ(q.db.Clone())
}

func (q *votingsQ) FilterBy(column string, value any) data.VotingsQ {
	q.sql = q.sql.Where(sq.Eq{column: value})
	return q
}

func (q *votingsQ) Get() (*data.Voting, error) {
	var result data.Voting
	err := q.db.Get(&result, q.sql)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &result, err
}

func (q *votingsQ) Select() ([]data.Voting, error) {
	var result []data.Voting
	err := q.db.Select(&result, q.sql)
	return result, err
}
