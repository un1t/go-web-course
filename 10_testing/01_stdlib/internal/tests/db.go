package tests

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupDB() *gorm.DB {
	err := godotenv.Load("../../.env.test")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	databaseUrl := os.Getenv("DATABASE_URL")
	if databaseUrl == "" {
		panic("missing DATABASE_URL")
	}

	db, err := gorm.Open(postgres.Open(databaseUrl), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	bytes, err := ioutil.ReadFile("../../migrations/01_init.up.sql")
	if err != nil {
		panic(err)
	}

	sql := string(bytes)

	err = db.Exec(sql).Error
	if err != nil {
		panic(err)
	}

	return db
}

func TeardownDB(db *gorm.DB) {
	bytes, err := ioutil.ReadFile("../../migrations/01_init.down.sql")
	if err != nil {
		panic(err)
	}
	sql := string(bytes)

	err = db.Exec(sql).Error
	if err != nil {
		panic(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		// log.Fatalln(err)
	}
	sqlDB.Close()
}
