package server

import (
	"fmt"
	"github.com/andrewshostak/awesome-service/handler"
	"github.com/andrewshostak/awesome-service/middleware"
	"github.com/andrewshostak/awesome-service/repository"
	"github.com/andrewshostak/awesome-service/service"
	"github.com/caarlos0/env/v6"
	"github.com/gin-gonic/gin"
	timeout "github.com/vearne/gin-timeout"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

type serverConfig struct {
	Port       string `env:"PORT" envDefault:"8080"`
	PgHost     string `env:"PG_HOST" envDefault:"localhost"`
	PgUser     string `env:"PG_USER" envDefault:"postgres"`
	PgPassword string `env:"PG_PASSWORD"`
	PgPort     string `env:"PG_PORT" envDefault:"5432"`
	PgDatabase string `env:"PG_DATABASE" envDefault:"postgres"`
	Timeout    int    `env:"TIMEOUT" envDefault:"10"`
}

func StartServer() {
	config := serverConfig{}
	if err := env.Parse(&config); err != nil {
		panic(err)
	}

	r := gin.Default()
	r.Use(middleware.ErrorHandle())
	r.Use(timeout.Timeout(
		timeout.WithTimeout(time.Duration(config.Timeout)*time.Second),
		timeout.WithDefaultMsg(`{"error": "timeout error"}`),
	))

	connectionParams := fmt.Sprintf(
		"host=%s user=%s password=%s port=%s database=%s sslmode=disable",
		config.PgHost,
		config.PgUser,
		config.PgPassword,
		config.PgPort,
		config.PgDatabase,
	)
	db, err := gorm.Open(postgres.Open(connectionParams))
	if err != nil {
		panic(err)
	}

	questionRepo := repository.NewQuestionRepo(db)

	questionService := service.NewQuestionService(questionRepo)

	questionHandler := handler.NewQuestionHandler(questionService)

	r.DELETE("/questions/:id", questionHandler.Delete)
	r.PUT("/questions/:id", questionHandler.Update)
	r.GET("/questions/:id", questionHandler.One)
	r.POST("/questions", questionHandler.Create)
	r.GET("/questions", questionHandler.List)

	r.Run()
}
