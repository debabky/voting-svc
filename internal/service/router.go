package service

import (
	"github.com/debabky/voting-svc/internal/data/pg"
	"github.com/debabky/voting-svc/internal/service/api/handlers"
	"github.com/go-chi/chi"
	"gitlab.com/distributed_lab/ape"
)

func (s *service) router() chi.Router {
	r := chi.NewRouter()

	r.Use(
		ape.RecoverMiddleware(s.log),
		ape.LoganMiddleware(s.log),
		ape.CtxMiddleware(
			handlers.CtxLog(s.log),
			handlers.CtxMasterQ(pg.NewMasterQ(s.cfg.DB())),
		),
	)
	r.Route("/integrations/voting-svc", func(r chi.Router) { // TODO auth middleware
		r.Route("/voting", func(r chi.Router) {
			r.Get("/", handlers.GetVotings)
			r.Post("/vote", handlers.Vote)
			r.Post("/register", handlers.Register)
		})
		r.Get("/auth-data", handlers.GetAuthData)
	})

	return r
}
