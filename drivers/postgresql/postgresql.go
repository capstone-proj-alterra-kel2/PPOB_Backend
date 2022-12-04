package postgres_driver

import (
	"PPOB_BACKEND/drivers/postgresql/products"
	"PPOB_BACKEND/drivers/postgresql/producttypes"
	"PPOB_BACKEND/drivers/postgresql/providers"
	"PPOB_BACKEND/drivers/postgresql/users"
	"fmt"

	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type ConfigDB struct {
	DB_USERNAME string
	DB_PASSWORD string
	DB_NAME     string
	DB_HOST     string
	DB_PORT     string
}

// configure and connecting database
func (config *ConfigDB) InitDB() *gorm.DB {
	connectionString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		config.DB_HOST,
		config.DB_USERNAME,
		config.DB_PASSWORD,
		config.DB_NAME,
		config.DB_PORT,
	)

	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatalf("error when connecting to a database server: %s", err)
	}

	log.Println("connected to a database server")
	return db
}

// Migrating Struct into Table in Database
func DBMigrate(db *gorm.DB) {
	db.AutoMigrate(
		&users.User{},               // User
		&producttypes.ProductType{}, // ProductType
		&providers.Provider{},       // Provider
		&products.Product{},         // Product
	)
}

// Closing Database
func CloseDB(db *gorm.DB) error {
	database, err := db.DB()
	if err != nil {
		log.Printf("error when getting the database instance : %v", err)
		return err
	}

	if err := database.Close(); err != nil {
		log.Printf("error when closing the database connection : %v", err)
		return err
	}
	log.Println("database connection is closed")
	return nil
}
