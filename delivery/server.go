package delivery

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"online-store-application/config"
	"online-store-application/delivery/controller"
	"online-store-application/delivery/middleware"
	"online-store-application/manager"
)

type Server struct {
	ucManager manager.UseCaseManager
	engine    *echo.Echo
	host      string
	log       *logrus.Logger
}

// Constructor
func NewServer() *Server {
	cfg, err := config.NewConfig()
	if err != nil {
		fmt.Println(err)
	}
	infraManager, err := manager.NewInfraManager(cfg)
	if err != nil {
		fmt.Println(err)
	}

	// Instance Repo
	rm := manager.NewRepoManager(infraManager)

	// Instance UC
	ucm := manager.NewUseCaseManager(rm)

	hostAndPort := fmt.Sprintf("%s:%s", viper.GetString("APP_API_HOST"), viper.GetString("APP_API_PORT"))
	log := logrus.New()

	engine := echo.New()
	return &Server{
		ucManager: ucm,
		engine:    engine,
		host:      hostAndPort,
		log:       log,
	}
}

func (s *Server) Run() {
	s.initMiddlewares()
	s.initControllers()
	err := s.engine.Start(s.host)
	if err != nil {
		panic(err)
	}
}

func (s *Server) initMiddlewares() {
	s.engine.Use(middleware.LogRequest(s.log))
}

func (s *Server) initControllers() {
	controller.NewUsersController(s.ucManager.UsersUC(), s.ucManager.WalletsUC(), s.engine).AuthRoute()
	controller.NewCategoriesController(s.ucManager.CategoriesUC(), s.engine).CategoriesRoute()
	controller.NewProductsController(s.ucManager.ProductsUC(), s.engine).ProductsRoute()
	controller.NewCartsController(s.ucManager.CartsUC(), s.engine).CartsRoute()
}