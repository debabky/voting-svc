package jwt

import (
	"encoding/hex"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/debabky/voting-svc/internal/config"
	"github.com/golang-jwt/jwt/v5"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

const (
	AuthorizationHeaderName = "Authorization"
	BearerTokenPrefix       = "Bearer "
)

type JWTIssuer struct {
	key             []byte
	accessDuration  time.Duration
	refreshDuration time.Duration
}

func NewJWTIssuer(cfg *config.JWTConfig) *JWTIssuer {
	key, err := hex.DecodeString(cfg.SecretKey)
	if err != nil {
		return nil
	}

	return &JWTIssuer{
		key:             key,
		accessDuration:  cfg.AccessExpirationTime,
		refreshDuration: cfg.RefreshExpirationTime,
	}
}

func (j *JWTIssuer) IssueJWT(votingID, nullifier string, tokenType TokenType) (string, error) {
	raw := (&RawJWT{make(jwt.MapClaims)}).
		SetVotingID(votingID).
		SetNullifier(nullifier)

	switch tokenType {
	case AccessTokenType:
		raw.
			SetTokenAccess().
			SetExpirationTimestamp(j.accessDuration)
	case RefreshTokenType:
		raw.
			SetTokenRefresh().
			SetExpirationTimestamp(j.refreshDuration)
	}

	return jwt.NewWithClaims(jwt.SigningMethodHS256, raw.claims).SignedString(j.key)
}

func (j *JWTIssuer) ValidateJWT(tokenStr string) (TokenClaims, error) {
	var token *jwt.Token

	key := func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return j.key, nil
	}

	token, err := jwt.Parse(tokenStr, key, jwt.WithExpirationRequired())
	if err != nil {
		return TokenClaims{}, errors.Wrap(err, "failed to parse JWT")
	}

	var (
		raw RawJWT
		ok  bool
	)
	if raw.claims, ok = token.Claims.(jwt.MapClaims); !ok {
		return TokenClaims{}, nil
	}

	votingID, ok := raw.VotingID()
	if !ok {
		return TokenClaims{}, errors.New("voting_id is invalid")
	}

	nullifier, ok := raw.Nullifier()
	if !ok {
		return TokenClaims{}, errors.New("nullifier is invalid")
	}

	return TokenClaims{
		VotingID:  votingID,
		Nullifier: nullifier,
	}, nil
}

func GetBearer(r *http.Request) (string, error) {
	authHeader := r.Header.Get(AuthorizationHeaderName)
	authHeaderSplit := strings.Split(authHeader, BearerTokenPrefix)

	if len(authHeaderSplit) != 2 {
		return "", errors.New("invalid token")
	}

	return authHeaderSplit[1], nil
}
