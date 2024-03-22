/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

import "github.com/google/uuid"

type VotingOptionAttributes struct {
	Description *string   `json:"description,omitempty"`
	Name        string    `json:"name"`
	VotesNumber int32     `json:"votes_number"`
	VotingId    uuid.UUID `json:"voting_id"`
}
