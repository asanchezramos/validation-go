package main

import (
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/goandreus/validation-go/config"
	"github.com/goandreus/validation-go/datalayer"
	"github.com/goandreus/validation-go/handler"
	"github.com/goandreus/validation-go/model"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func init() {
	if err := config.InitViper(); err != nil {
		log.WithError(err).Info("config not loaded correctly")
	}

	if err := ensureDir("./public"); err != nil {
		fmt.Println("Directory creation failed with error: " + err.Error())
		os.Exit(1)
	}
}

func main() {
	datastore := datalayer.NewDatastore()
	db, err := datastore.Open()
	if err != nil {
		panic(err)
	}

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())
	// Configure middleware with the custom claims type
	configJwt := middleware.JWTConfig{
		Claims:     &model.JwtCustomClaims{},
		SigningKey: []byte(fmt.Sprintf("%s", viper.Get("VALIDATION_API_SECRET"))),
	}
	e.Static("/public", "public")
	handler.NewAuthHandler(e, db)
	r := e.Group("")
	r.Use(middleware.JWTWithConfig(configJwt))
	handler.NewMobileHandler(r, db)

	e.Logger.Fatal(e.Start(":4000"))
}

func ensureDir(dirName string) error {

	err := os.MkdirAll(dirName, os.ModePerm)

	if err == nil || os.IsExist(err) {
		return nil
	} else {
		return err
	}
}
