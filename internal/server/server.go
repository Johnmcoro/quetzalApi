package server

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/golang-migrate/migrate/v4"
	pg "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	"github.com/johnmcoro/quetzalapi/internal/storage/postgres"
	transport "github.com/johnmcoro/quetzalapi/internal/transport/http"

	"go.uber.org/zap"
)

type Server struct {
	Router *chi.Mux
	DB     *sqlx.DB
	Logger *zap.Logger
}

func New() *Server {
	router := chi.NewRouter()
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	db, err := postgres.New()
	logger.Info("Test logger")
	if err != nil {
		logger.Error("Error connecting to database")
	}
	env := transport.NewEnv(db, logger, router)
	transport.InitializeRoutes(env)
	return &Server{
		Router: router,
		Logger: logger,
		DB:     db,
	}
}

func (s *Server) Run() error {
	http.ListenAndServe(":8080", s.Router)
	s.Logger.Info("Server starting on port 8080")
	return nil
}

func (s *Server) Migrate() error {
	driver, err := pg.WithInstance(s.DB.DB, &pg.Config{})
	if err != nil {
		return err
	}
	m, err := migrate.NewWithDatabaseInstance("file://migrations", "postgres", driver)
	if err != nil {
		return err
	}
	s.Logger.Info("Running migrations...")
	if err := m.Up(); err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
