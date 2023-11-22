package delivery

import (
	"api-with-sql-nativ/config"
	"api-with-sql-nativ/delivery/controller"
	inframanager "api-with-sql-nativ/manager/infra_manager"
	repomanager "api-with-sql-nativ/manager/repo_manager"
	usecasemanager "api-with-sql-nativ/manager/usecase_manager"

	"github.com/gin-gonic/gin"
)

type serverRequirement struct {
	usecase usecasemanager.UsecaseManager
	engine  *gin.Engine
	cfg     *config.Config
}

func (s *serverRequirement) setupRouter() {
	v1 := s.engine.Group("api/v1")
	controller.NewControllerExample(v1, s.usecase.UsecaseExampleManager()).Router()
}

func (s *serverRequirement) RUN() {
	s.setupRouter()

	if err := s.engine.Run(s.cfg.API_PORT); err != nil {
		panic(err)
	}
}

func Server() *serverRequirement {
	cfg := config.NewConfig()
	engine := gin.Default()
	infra := inframanager.NewInfraManager(cfg)
	repo := repomanager.NewRepoManager(infra)
	usecaseExample := usecasemanager.NewUsecaseManager(repo)

	return &serverRequirement{
		usecase: usecaseExample,
		engine:  engine,
		cfg:     cfg,
	}
}
