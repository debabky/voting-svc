package data

type VotesQ interface {
	New() VotesQ
	Insert(vote Vote) error
	FilterBy(column string, value any) VotesQ
	Count() (int64, error)
	Select() ([]Vote, error)
}

type Vote struct {
	VotingID     string `db:"voting_id" structs:"voting_id"`
	VotingOption string `db:"voting_option" structs:"voting_option"`
	Rank         *int64 `db:"rank" structs:"rank"`
	Nullifier    string `db:"nullifier" structs:"nullifier"`
}
