package service

import (
	"github.com/go-chi/chi"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/tokend/url-shortener-svc/internal/service/handlers"
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
	r.Route("/integrations/url-shortener-svc", func(r chi.Router) {
		// configure endpoints here
		r.Post("/shorten", handlers.ShortenURLHandler)
		r.Get("/short/{code}", handlers.ResolveURLHandler)
	})

	return r
}
