package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/versegeek/go-skeleton/config"
	"github.com/versegeek/go-skeleton/internal/handler"
	"github.com/versegeek/go-skeleton/internal/service"
	"github.com/versegeek/toolkit/pkg/provider"
)

type Server struct {
	provider.AbstractRunProvider

	// mysqlProvider *mysql.MySQL
	router  *gin.Engine
	handler handler.Handler
}

func New() *Server {
	return &Server{
		// mysqlProvider: mysqlProvider,
	}
}

func (s *Server) Init() error {
	svc := service.New()

	s.router = gin.Default()
	// gin.New()
	s.handler = handler.New(svc)

	// authv1.RegisterClientAPIServer(nil, s.handler)
	s.RegisterRoute()

	return nil
}

func (s *Server) Run() error {
	srv := &http.Server{
		Addr:        ":3846",
		Handler:     s.router,
		ReadTimeout: 3 * time.Second,
	}

	// authServerv1.RegisterHandlers(s.router, s.handler)

	message := fmt.Sprintf("[Server] 🚀 Gin Server is running, port: [%s]", config.GetConfig().ServerPort)
	logrus.Info(message)
	s.SetRunning(true)

	return srv.ListenAndServe()
}

func (s *Server) Close() error {
	s.SetRunning(false)
	return nil
}

func (s *Server) RegisterRoute() {
	gin.SetMode(gin.DebugMode)
	s.router.Use(gin.Recovery())
	s.router.GET("/", s.handler.HomeEndpoint)
	s.router.GET("/health", s.handler.HomeEndpoint)
	s.router.GET("/version", s.handler.VersionEndpoint)

	// v1 := s.router.Group("/v1")
	s.router.GET("/auth", s.handler.AuthEndpoint)
	s.router.GET("/callback", s.handler.CallbackEndpoint)
	s.router.POST("/token", s.handler.TokenEndpoint)
	s.router.GET("/refresh", s.handler.RefreshTokenEndpoint)
	// s.router.Run()
}
