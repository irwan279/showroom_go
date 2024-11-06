package handler

import (
	"mncbank/config"
	"mncbank/manager"
	"mncbank/middleware"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

type Server interface {
	Run()
}

type server struct {
	usecaseManager manager.UsecaseManager

	srv  *gin.Engine
	host string
}

func (s *server) Run() {
	// session
	store := cookie.NewStore([]byte("secret"))

	s.srv.Use(middleware.LoggerMiddleware())
	s.srv.Use(sessions.Sessions("session", store))

	// handler
	NewUserHandler(s.srv, s.usecaseManager.GetUserUsecase())
	NewLoginHandler(s.srv, s.usecaseManager.GetLoginUsecase())
	NewCustomerHandler(s.srv, s.usecaseManager.GetCustomerUsecase())

	s.srv.Run(s.host)
}

func NewServer() Server {
	c := config.NewConfig()

	infra := manager.NewInfraManager(c)
	repo := manager.NewRepoManager(infra)
	usecase := manager.NewUsecaseManager(repo)

	srv := gin.Default()

	return &server{
		usecaseManager: usecase,
		srv:            srv,
		host:           c.AppPort,
	}
}
