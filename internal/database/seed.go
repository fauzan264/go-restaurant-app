package database

import (
	"log"

	"github.com/fauzan264/go-restaurant-app/internal/model"
	"github.com/fauzan264/go-restaurant-app/internal/model/constant"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SeedDB(databaseUrl string) {
	foodMenu := []model.MenuItem{
		{
			Name: "Bakmie",
			OrderCode: "bakmie",
			Price: 37500,
			Type: constant.MenuTypeFood,
		},
		{
			Name: "Ayam Rica-Rica",
			OrderCode: "ayam_rica_rica",
			Price: 41250,
			Type: constant.MenuTypeFood,
		},
	}

	drinkMenu := []model.MenuItem{
		{
			Name: "Es Jeruk",
			OrderCode: "es_jeruk",
			Price: 8000,
			Type: constant.MenuTypeDrink,
		},
		{
			Name: "Es Teh",
			OrderCode: "es_teh",
			Price: 5000,
			Type: constant.MenuTypeDrink,
		},
	}

	db, err := gorm.Open(postgres.Open(databaseUrl))
	if err != nil {
		log.Fatal(err)
	}

	// db.AutoMigrate(&model.MenuItem{})
	if err := db.First(&model.MenuItem{}).Error; err == gorm.ErrRecordNotFound {
		db.Create(&foodMenu)
		db.Create(&drinkMenu)
	}
}


