package pg

import (
	"database/sql"
	sq "github.com/Masterminds/squirrel"
	"github.com/debabky/voting-svc/internal/data"
	"github.com/fatih/structs"
	"gitlab.com/distributed_lab/kit/pgdb"
)

const registrationsTableName = "registrations"

var (
	registrationsSelector = sq.Select("*").From(registrationsTableName)
	registrationsUpdate   = sq.Update(registrationsTableName)
)

func NewRegistrationsQ(db *pgdb.DB) data.RegistrationsQ {
	return &registrationsQ{
		db:  db,
		sql: registrationsSelector,
		upd: registrationsUpdate,
	}
}

type registrationsQ struct {
	db  *pgdb.DB
	sql sq.SelectBuilder
	upd sq.UpdateBuilder
}

func (q *registrationsQ) New() data.RegistrationsQ {
	return NewRegistrationsQ(q.db.Clone())
}

func (q *registrationsQ) Insert(value data.Registration) error {
	clauses := structs.Map(value)
	stmt := sq.Insert(registrationsTableName).SetMap(clauses)
	err := q.db.Exec(stmt)
	return err
}

func (q *registrationsQ) FilterBy(column string, value any) data.RegistrationsQ {
	q.sql = q.sql.Where(sq.Eq{column: value})
	return q
}

func (q *registrationsQ) Get() (*data.Registration, error) {
	var result data.Registration
	err := q.db.Get(&result, q.sql)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &result, err
}
