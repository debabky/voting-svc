package service

import (
	"github.com/debabky/voting-svc/internal/data/pg"
	"github.com/debabky/voting-svc/internal/service/api/handlers"
	"github.com/debabky/voting-svc/internal/service/cookies"
	"github.com/debabky/voting-svc/internal/service/jwt"
	"github.com/go-chi/chi"
	"gitlab.com/distributed_lab/ape"
)

func (s *service) router() chi.Router {
	r := chi.NewRouter()

	jwtIssuer := jwt.NewJWTIssuer(s.cfg.JWTConfig())

	r.Use(
		ape.RecoverMiddleware(s.log),
		ape.LoganMiddleware(s.log),
		ape.CtxMiddleware(
			handlers.CtxLog(s.log),
			handlers.CtxMasterQ(pg.NewMasterQ(s.cfg.DB())),
			handlers.CtxJWTIssuer(jwtIssuer),
			handlers.CtxCookies(cookies.NewCookies(s.cfg.CookiesConfig())),
		),
	)
	r.Route("/integrations/voting-svc", func(r chi.Router) {
		r.Route("/voting", func(r chi.Router) {
			r.With(handlers.AuthMiddleware(jwtIssuer, s.log, jwt.AccessTokenType)).Get("/{id}", handlers.GetVoting)
			r.Get("/list", handlers.GetVotings)
			r.Post("/vote", handlers.Vote)
		})
		r.Get("/auth-data/{id}", handlers.GetAuthData)
	})

	return r
}
