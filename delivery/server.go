package delivery

import (
	"enigma-laundry-clean-code/config"
	"enigma-laundry-clean-code/delivery/handler"
	"enigma-laundry-clean-code/manager"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

type Server struct {
	ucManager manager.UseCaseManager
	engine    *gin.Engine
	host      string
}

func (s *Server) setupHandler() {
	rg := s.engine.Group("/api/v1")
	handler.NewCustomerHandler(s.ucManager.CustomerCase(), rg).Route()
}

func (s *Server) Run() {
	s.setupHandler()
	if err := s.engine.Run(s.host); err != nil {
		panic(err)
	}
}

func NewServer() *Server {
	conf, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}
	infraManager, _ := manager.NewInfraManager(conf)
	repoManager := manager.NewRepoManager(infraManager)
	useCaseManager := manager.NewUseCaseManager(repoManager)
	engine := gin.Default()
	host := fmt.Sprintf(":%s", conf.ApiPort)

	return &Server{
		ucManager: useCaseManager,
		engine:    engine,
		host:      host,
	}
}
