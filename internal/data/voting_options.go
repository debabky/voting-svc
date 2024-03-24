package data

type VotingOptionsQ interface {
	New() VotingOptionsQ
	FilterBy(column string, value any) VotingOptionsQ
	Select() ([]VotingOption, error)
}

type VotingOption struct {
	ID          int64   `db:"id" structs:"id"`
	Name        string  `db:"name" structs:"name"`
	VotingID    string  `db:"voting_id" structs:"voting_id"`
	Description *string `db:"description" structs:"description"`
}
