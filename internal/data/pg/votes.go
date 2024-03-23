package pg

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/debabky/voting-svc/internal/data"
	"github.com/fatih/structs"
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

func (q *votesQ) Insert(value data.Vote) error {
	clauses := structs.Map(value)
	stmt := sq.Insert(votesTableName).SetMap(clauses)
	err := q.db.Exec(stmt)
	return err
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

func (q *votesQ) Select() ([]data.Vote, error) {
	var result []data.Vote
	err := q.db.Select(&result, q.sql)
	return result, err
}
