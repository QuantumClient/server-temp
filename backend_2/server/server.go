package server

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gookit/rux"
	"github.com/gookit/slog"
	"net/http"
	"quantumclient.org/backend/v2/controller"
	"quantumclient.org/backend/v2/middleware"
	"quantumclient.org/backend/v2/models"
	"quantumclient.org/backend/v2/repository"
	"quantumclient.org/backend/v2/services"
	"time"
)

type Server struct {
	Config *models.Config
	controllers map[string]controller.WebController
}

func NewServer(config *models.Config) *Server {
	return &Server{Config: config}
}

func (s *Server) Init() {
	db := getDB(s.Config)

	userRepo := repository.NewUserRepository(db)
	authRepo := repository.NewAuthRepository(db, userRepo)

	jwtService := services.NewJwtService(s.Config, authRepo)

	middleware.NewAuthMiddleware(jwtService)

	capeRepo := repository.NewCapeRepository(db)
	capeService := services.NewCapeService(capeRepo)


	s.controllers = map[string]controller.WebController{
		"online": controller.NewOnlineController(),
		"capes": controller.NewCapeController(capeService),
	}
}

func (s *Server) Start() {

	r := rux.New()

	for name, webController := range s.controllers {
		r.Controller(name, webController)
		slog.Infof("Controller %s registered", name)
	}

	slog.Infof("Server started on http://%s:%s", s.Config.Server.Host, s.Config.Server.Port)
	err := http.ListenAndServe(s.Config.Server.Host + ":" + s.Config.Server.Port, r)
	if err != nil {
		slog.Fatal(err)
	}
}

func getDB(config *models.Config) *sql.DB {
	db, err := sql.Open("mysql", config.Database.User+":"+config.Database.Password+"@tcp("+config.Database.Host+":"+config.Database.Port+")/"+config.Database.Name+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		slog.Fatal(err)
	}
	db.SetConnMaxLifetime(time.Minute * 5)
	slog.Info("Connected to database")
	return db
}