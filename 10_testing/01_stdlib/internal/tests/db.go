package tests

import (
	"io/ioutil"

	"gorm.io/gorm"
)

func SetupDB(db *gorm.DB) error {
	bytes, err := ioutil.ReadFile("../../migrations/01_init.up.sql")
	if err != nil {
		return err
	}

	sql := string(bytes)

	err = db.Exec(sql).Error
	if err != nil {
		TeardownDB(db)
		err = db.Exec(sql).Error
		if err != nil {
			return err
		}
	}

	return nil
}

func TeardownDB(db *gorm.DB) error {
	bytes, err := ioutil.ReadFile("../../migrations/01_init.down.sql")
	if err != nil {
		return err
	}
	sql := string(bytes)

	err = db.Exec(sql).Error
	if err != nil {
		return err
	}

	return nil
}
