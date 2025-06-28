package app

import (
	"context"
	"learning/app/controllers"
	"learning/app/models"
	"learning/config"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/roonglit/credentials/pkg/credentials"
)

type Application struct {
	Server *controllers.Server
}

func New() *Application {
	config := loadConfig()

	store := connectDb(config)

	server := controllers.New(
		config,
		store,
	)

	return &Application{
		Server: server,
	}
}

func (app *Application) Run() {
	app.Server.Run()
}

func loadConfig() *config.Config {
	reader := credentials.NewConfigReader()

	var config config.Config
	if err := reader.Read(gin.Mode(), &config); err != nil {
		panic("failed to load config: " + err.Error())
	}
	return &config
}

func connectDb(config *config.Config) models.Store {
	dbConfig, err := pgxpool.ParseConfig(config.DBUri)
	if err != nil {
		panic("failed to parse db config: " + err.Error())
	}

	connPool, err := pgxpool.New(
		context.Background(),
		dbConfig.ConnString(),
	)

	store := models.NewStore(connPool)
	return store
}
