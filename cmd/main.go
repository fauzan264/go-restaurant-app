package main

import (
	"log"
	"os"

	"github.com/fauzan264/go-restaurant-app/internal/database"
	"github.com/fauzan264/go-restaurant-app/internal/handler/restapi"
	"github.com/fauzan264/go-restaurant-app/internal/repository/menu"
	"github.com/fauzan264/go-restaurant-app/internal/usecase/resto"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

var (databaseURL, port string)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Couldn't load .env file!")
	}

	databaseURL = os.Getenv("DATABASE_URL")
	port = os.Getenv("PORT")
}

func main() {
	e := echo.New()
	db := database.GetDB(databaseURL)
	
	menuRepo := menu.GetRepository(db)

	restoUsecase := resto.GetUsecase(menuRepo)

	handler := restapi.NewHandler(restoUsecase)

	restapi.LoadRoutes(e, handler)
	// e.GET("/menu", restapi.getMenu)
	e.Logger.Fatal(e.Start(port))
}