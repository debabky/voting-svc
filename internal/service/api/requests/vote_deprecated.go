package requests

//type Vote struct {
//	VotingOption string `json:"voting_option"`
//	Rank         *int64 `json:"rank"`
//}
//
//type VoteRequestVotes struct {
//	Votes []Vote `json:"votes"`
//}
//
//type VoteRequestAttributes struct {
//	Attributes VoteRequestVotes `json:"attributes"`
//}
//
//type VoteRequest struct {
//Data VoteRequestAttributes `json:"data"`
//}
//
//func NewVoteRequest(r *http.Request) (VoteRequest, error) {
//	var request VoteRequest
//
//	err := json.NewDecoder(r.Body).Decode(&request)
//	if err != nil {
//		return request, errors.Wrap(err, "failed to unmarshal")
//	}
//
//	return request, nil
//}
