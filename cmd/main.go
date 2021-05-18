package main

import (
	"log"

	_ "github.com/lib/pq"
	"github.com/myownhatred/botAPI"
	"github.com/myownhatred/botAPI/pkg/handler"
	"github.com/myownhatred/botAPI/pkg/repository"
	"github.com/myownhatred/botAPI/pkg/service"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error loading config: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: viper.GetString("db.password"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		log.Fatalf("failed to initialize db: %s", err.Error())
	}

	rep := repository.NewRepository(db)
	repG := repository.NewGoogleRepository()
	serv := service.NewService(rep)
	servG := service.NewGoogleService(repG)
	handlers := handler.NewHandler(serv, servG)
	srv := new(botAPI.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("Server run error: %s", err.Error())
	}
}

// reads from /configs/config.yaml next fields:
// port
func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
