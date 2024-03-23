package pg

import (
	"database/sql"
	sq "github.com/Masterminds/squirrel"
	"github.com/debabky/voting-svc/internal/data"
	"github.com/fatih/structs"
	"gitlab.com/distributed_lab/kit/pgdb"
)

const verificationRequestsTableName = "verification_requests"

var (
	verificationRequestsSelector = sq.Select("*").From(verificationRequestsTableName)
	verificationRequestsUpdate   = sq.Update(verificationRequestsTableName)
)

func NewVerificationRequestsQ(db *pgdb.DB) data.VerificationRequestsQ {
	return &verificationRequestsQ{
		db:  db,
		sql: verificationRequestsSelector,
		upd: verificationRequestsUpdate,
	}
}

type verificationRequestsQ struct {
	db  *pgdb.DB
	sql sq.SelectBuilder
	upd sq.UpdateBuilder
}

func (q *verificationRequestsQ) New() data.VerificationRequestsQ {
	return NewVerificationRequestsQ(q.db.Clone())
}

func (q *verificationRequestsQ) Insert(value data.VerificationRequest) error {
	clauses := structs.Map(value)
	stmt := sq.Insert(verificationRequestsTableName).SetMap(clauses)
	err := q.db.Exec(stmt)
	return err
}

func (q *verificationRequestsQ) FilterBy(column string, value any) data.VerificationRequestsQ {
	q.sql = q.sql.Where(sq.Eq{column: value})
	return q
}

func (q *verificationRequestsQ) Get() (*data.VerificationRequest, error) {
	var result data.VerificationRequest
	err := q.db.Get(&result, q.sql)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &result, err
}
