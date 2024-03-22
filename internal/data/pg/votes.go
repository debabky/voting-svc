package pg

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/debabky/voting-svc/internal/data"
	"gitlab.com/distributed_lab/kit/pgdb"
)

const votesTableName = "votes"

var (
	votesSelector = sq.Select("*").From(votesTableName)
	votesUpdate   = sq.Update(votesTableName)
)

func NewVotesQ(db *pgdb.DB) data.VotesQ {
	return &votesQ{
		db:  db,
		sql: votesSelector,
		upd: votesUpdate,
	}
}

type votesQ struct {
	db  *pgdb.DB
	sql sq.SelectBuilder
	upd sq.UpdateBuilder
}

func (q *votesQ) New() data.VotesQ {
	return NewVotesQ(q.db.Clone())
}

func (q *votesQ) FilterBy(column string, value any) data.VotesQ {
	q.sql = q.sql.Where(sq.Eq{column: value})
	return q
}

func (q *votesQ) Count() (int64, error) {
	stmt := sq.Select("COUNT(*)").FromSelect(q.sql, "s")
	var result []int64
	err := q.db.Select(&result, stmt)
	if err != nil {
		return 0, err
	}
	return result[0], nil
}
