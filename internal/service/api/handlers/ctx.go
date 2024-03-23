package handlers

import (
	"context"
	"github.com/debabky/voting-svc/internal/data"
	"github.com/debabky/voting-svc/internal/service/cookies"
	"github.com/debabky/voting-svc/internal/service/jwt"
	"gitlab.com/distributed_lab/logan/v3"
	"net/http"
)

type ctxKey int

const (
	logCtxKey ctxKey = iota
	masterQCtxKey
	jwtIssuerCtxKey
	tokenClaimsCtxKey
	cookiesCtxKey
)

func CtxLog(entry *logan.Entry) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, logCtxKey, entry)
	}
}

func Log(r *http.Request) *logan.Entry {
	return r.Context().Value(logCtxKey).(*logan.Entry)
}

func CtxMasterQ(entry data.MasterQ) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, masterQCtxKey, entry)
	}
}

func MasterQ(r *http.Request) data.MasterQ {
	return r.Context().Value(masterQCtxKey).(data.MasterQ).New()
}

func CtxJWTIssuer(jwtIssuer *jwt.JWTIssuer) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, jwtIssuerCtxKey, jwtIssuer)
	}
}

func JWTIssuer(r *http.Request) *jwt.JWTIssuer {
	return r.Context().Value(jwtIssuerCtxKey).(*jwt.JWTIssuer)
}

func CtxTokenClaims(claims *jwt.TokenClaims) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, tokenClaimsCtxKey, claims)
	}
}

func TokenClaims(r *http.Request) *jwt.TokenClaims {
	return r.Context().Value(tokenClaimsCtxKey).(*jwt.TokenClaims)
}

func CtxCookies(cookies *cookies.Cookies) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, cookiesCtxKey, cookies)
	}
}

func Cookies(r *http.Request) *cookies.Cookies {
	return r.Context().Value(cookiesCtxKey).(*cookies.Cookies)
}
