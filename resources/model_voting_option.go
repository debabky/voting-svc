/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type VotingOption struct {
	Key
	Attributes VotingOptionAttributes `json:"attributes"`
}
type VotingOptionResponse struct {
	Data     VotingOption `json:"data"`
	Included Included     `json:"included"`
}

type VotingOptionListResponse struct {
	Data     []VotingOption `json:"data"`
	Included Included       `json:"included"`
	Links    *Links         `json:"links"`
}

// MustVotingOption - returns VotingOption from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustVotingOption(key Key) *VotingOption {
	var votingOption VotingOption
	if c.tryFindEntry(key, &votingOption) {
		return &votingOption
	}
	return nil
}
