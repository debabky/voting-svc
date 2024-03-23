package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type TokenClaims struct {
	VotingID  uuid.UUID
	Nullifier string
}

type TokenType string

func (t TokenType) String() string {
	return string(t)
}

var (
	AccessTokenType  TokenType = "access"
	RefreshTokenType TokenType = "refresh"
)

// RawJWT represents helper structure to provide setter and getter methods to work with JWT claims
type RawJWT struct {
	claims jwt.MapClaims
}

// Setters

func (r *RawJWT) SetVotingID(votingID string) *RawJWT {
	r.claims["voting_id"] = votingID
	return r
}

func (r *RawJWT) SetNullifier(nullifier string) *RawJWT {
	r.claims["nullifier"] = nullifier
	return r
}

func (r *RawJWT) SetTokenAccess() *RawJWT {
	r.claims["type"] = AccessTokenType
	return r
}

func (r *RawJWT) SetTokenRefresh() *RawJWT {
	r.claims["type"] = RefreshTokenType
	return r
}

func (r *RawJWT) SetExpirationTimestamp(expiration time.Duration) *RawJWT {
	r.claims["exp"] = jwt.NewNumericDate(time.Now().UTC().Add(expiration))
	return r
}

// Getters

func (r *RawJWT) VotingID() (id uuid.UUID, ok bool) {
	var (
		val interface{}
	)

	if val, ok = r.claims["voting_id"]; !ok {
		return
	}

	if id, ok = val.(uuid.UUID); !ok {
		return
	}

	return id, ok
}

func (r *RawJWT) Nullifier() (str string, ok bool) {
	var (
		val interface{}
	)

	if val, ok = r.claims["nullifier"]; !ok {
		return
	}

	if str, ok = val.(string); !ok {
		return
	}

	return str, ok
}
