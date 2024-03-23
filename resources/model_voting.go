/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type Voting struct {
	Key
	Attributes VotingAttributes `json:"attributes"`
}
type VotingResponse struct {
	Data     Voting   `json:"data"`
	Included Included `json:"included"`
}

type VotingListResponse struct {
	Data     []Voting `json:"data"`
	Included Included `json:"included"`
	Links    *Links   `json:"links"`
}

// MustVoting - returns Voting from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustVoting(key Key) *Voting {
	var voting Voting
	if c.tryFindEntry(key, &voting) {
		return &voting
	}
	return nil
}
