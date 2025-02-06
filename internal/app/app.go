package app

import (
	"log"

	"github.com/sonjinho/otel-deploy-backend/internal/config"
	"github.com/sonjinho/otel-deploy-backend/internal/database"
	"github.com/sonjinho/otel-deploy-backend/internal/services"
)

type Application struct {
	Config        *config.Config
	DB            *database.BadgerDB
	ConfigService *services.ConfigService
}

func NewApplication() *Application {
	cfg := config.LoadConfig()

	db, err := database.NewBadgerDB(cfg.DBPath)
	if err != nil {
		log.Fatalf("Failed to initialize BadgerDB: %v", err)
	}

	configService := services.NewConfigService(db)

	return &Application{Config: cfg, DB: db, ConfigService: configService}
}

func (app *Application) Run() {
	log.Println("Starting application...")
	// TODO: Add Http server init
	server := SetupServer(app)

	server.Listen(app.Config.ServerAddress + ":" + app.Config.Port)

}
