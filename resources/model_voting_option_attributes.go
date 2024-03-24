/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

import "github.com/google/uuid"

type VotingOptionAttributes struct {
	Description *string   `json:"description,omitempty"`
	Id          int64     `json:"id"`
	Name        string    `json:"name"`
	Points      *int64    `json:"points,omitempty"`
	Votes       *int64    `json:"votes,omitempty"`
	VotingId    uuid.UUID `json:"voting_id"`
}
