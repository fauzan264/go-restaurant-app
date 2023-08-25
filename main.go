package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type MenuItem struct {
	Name 		string
	OrderCode 	string
	Price 		int
}

func getFoodMenu(c echo.Context) error {
	foodMenu := []MenuItem{
		{
			Name: "Bakmie",
			OrderCode: "bakmie",
			Price: 37500,
		},
		{
			Name: "Ayam Rica-Rica",
			OrderCode: "ayam_rica_rica",
			Price: 41250,
		},
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": foodMenu,
	})
}

func getDrinkMenu(c echo.Context) error {
	drinkMenu := []MenuItem{
		{
			Name: "Es Jeruk",
			OrderCode: "es_jeruk",
			Price: 8000,
		},
		{
			Name: "Es Teh",
			OrderCode: "es_teh",
			Price: 5000,
		},
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": drinkMenu,
	})
}

func main() {
	e := echo.New()
	e.GET("/menu/food", getFoodMenu)
	e.GET("/menu/drink", getDrinkMenu)
	e.Logger.Fatal(e.Start(":14045"))
}