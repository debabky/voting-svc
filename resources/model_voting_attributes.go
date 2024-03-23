/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

import "time"

type VotingAttributes struct {
	ActiveUntil time.Time       `json:"active_until"`
	CreatedAt   time.Time       `json:"created_at"`
	Description string          `json:"description"`
	Name        string          `json:"name"`
	Options     *[]VotingOption `json:"options,omitempty"`
	Votes       *[]Vote         `json:"votes,omitempty"`
}
