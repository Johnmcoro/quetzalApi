package http

import (
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
	"github.com/johnmcoro/quetzalapi/internal/service"
	storage "github.com/johnmcoro/quetzalapi/internal/storage/postgres"
	"go.uber.org/zap"
)

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
	loggerMiddleware := NewLoggerMiddleware("zap", env.Logger)
	env.Router.Use(middleware.RealIP)
	env.Router.Use(middleware.RequestID)
	env.Router.Use(loggerMiddleware)
	env.Router.Get("/api/healthcheck", HealthCheck)
	env.Router.Mount("/api/v1/users", userRouter(env))
}

func userRouter(env *Env) chi.Router {
	r := chi.NewRouter()
	userStorage := storage.NewUserStorage(env.DB)
	userService := service.NewUserService(userStorage)
	userHandler := NewUserHandler(userService)
	r.Get("/", userHandler.GetUsers)
	return r
}
