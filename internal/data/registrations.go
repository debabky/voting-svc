package data

type RegistrationsQ interface {
	New() RegistrationsQ
	Insert(value Registration) error
	FilterBy(column string, value any) RegistrationsQ
	Get() (*Registration, error)
}

type Registration struct {
	VotingID  string `db:"voting_id" structs:"voting_id"`
	Nullifier string `db:"nullifier" structs:"nullifier"`
}
