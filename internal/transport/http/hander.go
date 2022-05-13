package http

import (
	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

// Handler to attach to router
type Env struct {
	Router *chi.Mux
	DB     *sqlx.DB
	Logger *zap.Logger
}

func NewEnv(db *sqlx.DB, logger *zap.Logger, router *chi.Mux) *Env {
	return &Env{
		Router: router,
		DB:     db,
		Logger: logger,
	}
}

func InitializeRoutes(env *Env) {
	env.Router.Get("/api/healthcheck", HealthCheck)
	env.Router.Mount("/api/user", userRouter(env))
}

func userRouter(env *Env) chi.Router {
	r := chi.NewRouter()

	return r
}
