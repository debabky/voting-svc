package service

import (
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
		),
	)
	r.Route("/integrations/voting-svc", func(r chi.Router) { // TODO auth middleware
		r.Route("/voting", func(r chi.Router) {
			r.Get("/", handlers.GetVotings)
			r.Post("/vote", handlers.Vote)
		})
		r.Post("/verify-proof", handlers.VerifyProof)
		r.Get("/auth-data", handlers.GetAuthData)
	})

	return r
}
