package data

type MasterQ interface {
	New() MasterQ

	VotingsQ() VotingsQ
	VotingOptionsQ() VotingOptionsQ
	RegistrationsQ() RegistrationsQ
	VotesQ() VotesQ

	Transaction(fn func(db MasterQ) error) error
}
