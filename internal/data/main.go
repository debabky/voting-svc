package data

type MasterQ interface {
	New() MasterQ

	VotingsQ() VotingsQ
	VotingOptionsQ() VotingOptionsQ
	RegistrationsQ() RegistrationsQ
	VotesQ() VotesQ
	VerificationRequestsQ() VerificationRequestsQ

	Transaction(fn func(db MasterQ) error) error
}
