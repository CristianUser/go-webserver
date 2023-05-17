package model

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Database() (*gorm.DB, error) {

	db, err := gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	if err = db.AutoMigrate(&Todo{}, &User{}, &Session{}); err != nil {
		log.Println(err)
	}

	return db, err

}
