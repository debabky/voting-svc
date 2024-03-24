package service

import (
	"github.com/debabky/voting-svc/internal/contracts"
	"github.com/debabky/voting-svc/internal/data/pg"
	"github.com/debabky/voting-svc/internal/service/api/handlers"
	"github.com/debabky/voting-svc/internal/service/cookies"
	"github.com/debabky/voting-svc/internal/service/jwt"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

func (s *service) router() chi.Router {
	r := chi.NewRouter()

	jwtIssuer := jwt.NewJWTIssuer(s.cfg.JWTConfig())
	if jwtIssuer == nil {
		panic(errors.New("failed to initialize JWT issuer"))
	}

	ethClient, err := ethclient.Dial(s.cfg.NetworkConfig().RPC)
	if err != nil {
		panic(errors.Wrap(err, "failed to dial connect"))
	}

	registrationContract, err := contracts.NewRegistration(s.cfg.NetworkConfig().VotingAddress, ethClient)
	if err != nil {
		panic(errors.Wrap(err, "failed to init new registration contract"))
	}

	r.Use(
		cors.Handler(cors.Options{
			AllowedOrigins:   []string{"*"},
			AllowedMethods:   []string{"GET", "POST"},
			AllowedHeaders:   []string{"Content-Type", "Authorization"},
			AllowCredentials: false,
			MaxAge:           300,
		}),
		ape.RecoverMiddleware(s.log),
		ape.LoganMiddleware(s.log),
		ape.CtxMiddleware(
			handlers.CtxLog(s.log),
			handlers.CtxMasterQ(pg.NewMasterQ(s.cfg.DB())),
			handlers.CtxJWTIssuer(jwtIssuer),
			handlers.CtxCookies(cookies.NewCookies(s.cfg.CookiesConfig())),
			handlers.CtxNetworkConfig(s.cfg.NetworkConfig()),
			handlers.CtxEthClient(ethClient),
			handlers.CtxRegistrationContract(registrationContract),
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
