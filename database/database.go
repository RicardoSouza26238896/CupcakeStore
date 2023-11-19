package database

import (
	"github.com/bitebait/cupcakestore/config"
	"github.com/bitebait/cupcakestore/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func SetupDatabase() {
	var err error
	DB, err = gorm.Open(sqlite.Open(config.GetEnv("DB_PATH", "database.db")), &gorm.Config{})
	if err != nil {
		panic("Falha ao conectar ao banco de dados: " + err.Error())
	}

	migrateModels(DB)
	seedStoreConfig(DB)
}

func migrateModels(db *gorm.DB) {
	err := db.AutoMigrate(
		&models.User{},
		&models.Profile{},
		&models.Product{},
		&models.Stock{},
		&models.StoreConfig{},
		&models.ShoppingCart{},
	)
	if err != nil {
		log.Panic("erro ao migrar os modelos")
	}
}

func seedStoreConfig(db *gorm.DB) {
	var count int64

	db.Model(&models.StoreConfig{}).Count(&count)
	if count == 0 {
		// Criar uma nova StoreConfig
		storeConfig := models.StoreConfig{}

		if err := db.Create(&storeConfig).Error; err != nil {
			log.Fatalf("Falha ao criar StoreConfig: %v", err)
		}

	}
}
