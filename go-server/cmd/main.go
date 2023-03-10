/*
 * Digimon Service API
 *
 * 提供孵化數碼蛋與培育等數碼寶貝養成服務
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package main

import (
	"database/sql"
	"fmt"

	_dietRepo "github.com/Jimmyweng006/clean-architecture/go-server/diet/repository/postgresql"
	_dietUsecase "github.com/Jimmyweng006/clean-architecture/go-server/diet/usecase"
	_digimonHandlerHttpDelivery "github.com/Jimmyweng006/clean-architecture/go-server/digimon/delivery/http"
	_digmonRepo "github.com/Jimmyweng006/clean-architecture/go-server/digimon/repository/postgresql"
	_digimonUsecase "github.com/Jimmyweng006/clean-architecture/go-server/digimon/usecase"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	_ "github.com/lib/pq"
)

func init() {
	viper.SetConfigFile(".env")
	viper.SetConfigType("dotenv")

	if err := viper.ReadInConfig(); err != nil {
		logrus.Fatal("Fatal error config file: %v \n", err)
	}
}

func main() {
	logrus.Info("HTTP server started")

	restfulHost := viper.GetString("RESTFUL_HOST")
	restfulPort := viper.GetString("RESTFUL_PORT")
	dbHost := viper.GetString("DB_HOST")
	dbDatabase := viper.GetString("DB_DATABASE")
	dbUser := viper.GetString("DB_USER")
	dbPassword := viper.GetString("DB_PASSWORD")

	// restfulHost := "0.0.0.0"
	// restfulPort := "5000"
	// dbHost := "db"
	// dbDatabase := "postgres"
	// dbUser := "root"
	// dbPassword := "root"

	db, err := sql.Open(
		"postgres",
		fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbUser, dbPassword, dbDatabase),
	)
	if err != nil {
		logrus.Fatal(err)
	}
	if err = db.Ping(); err != nil {
		logrus.Fatal(err)
	}

	r := gin.Default()

	digimonRepo := _digmonRepo.NewpostgresqlDigimonRepository(db)
	dietRepo := _dietRepo.NewPostgresqlDietRepository(db)

	digimonUsecase := _digimonUsecase.NewDigimonUsecase(digimonRepo)
	dietUsecase := _dietUsecase.NewDietUsecase(dietRepo)

	_digimonHandlerHttpDelivery.NewDigimonHandler(r, digimonUsecase, dietUsecase)

	logrus.Fatal(r.Run(restfulHost + ":" + restfulPort))
}
