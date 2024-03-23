package handlers

import (
	"net/http"

	"github.com/debabky/voting-svc/internal/service/jwt"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/distributed_lab/logan/v3"
)

func AuthMiddleware(issuer *jwt.JWTIssuer, log *logan.Entry, tokenType jwt.TokenType) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			var (
				token string
				err   error
			)

			if cookie, err := r.Cookie(tokenType.String()); err == nil {
				token = cookie.Value
			} else {
				token, err = jwt.GetBearer(r)
				if err != nil {
					log.WithError(err).Debug("failed to get bearer token")
					ape.RenderErr(w, problems.Unauthorized())
					return
				}
			}

			claims, err := issuer.ValidateJWT(token)
			if err != nil {
				log.WithError(err).Debug("failed validate bearer token")
				ape.RenderErr(w, problems.Unauthorized())
				return
			}

			ctx := CtxTokenClaims(&claims)(r.Context())
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
