package main

import (
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"hac/internal/app/ds"
	"hac/internal/app/dsn"
)

func main() {
	_ = godotenv.Load()
	db, err := gorm.Open(postgres.Open(dsn.FromEnv()), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	err = db.AutoMigrate(&ds.Users{})
	if err != nil {
		panic("cant migrate db goods")
	}

	err = db.AutoMigrate(&ds.FeedBack{})
	if err != nil {
		panic("cant migrate db goods")
	}

	err = db.AutoMigrate(&ds.Objests{})
	if err != nil {
		panic("cant migrate db goods")
	}

	err = db.AutoMigrate(&ds.Occupation{})
	if err != nil {
		panic("cant migrate db goods")
	}

	err = db.AutoMigrate(&ds.Favorites{})
	if err != nil {
		panic("cant migrate db goods")
	}

	err = db.AutoMigrate(&ds.Groups{})
	if err != nil {
		panic("cant migrate db goods")
	}

	err = db.AutoMigrate(&ds.Coordinates{})
	if err != nil {
		panic("cant migrate db goods")
	}

}
