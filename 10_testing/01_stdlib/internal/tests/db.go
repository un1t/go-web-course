package tests

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"

	"gorm.io/gorm"
)

func SetupDB(db *gorm.DB) error {
	sql, err := ConcatMigrations("../../migrations/*.up.sql")
	if err != nil {
		return err
	}

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
	sql, err := ConcatMigrations("../../migrations/*.down.sql")
	if err != nil {
		return err
	}

	err = db.Exec(sql).Error
	if err != nil {
		return err
	}

	return nil
}

func ConcatMigrations(pattern string) (string, error) {
	filenames, err := filepath.Glob(pattern)
	if err != nil {
		return "", err
	}
	var contents []string
	for _, filename := range filenames {
		bytes, err := ioutil.ReadFile(filename)
		if err != nil {
			return "", err
		}
		contents = append(
			contents,
			fmt.Sprintf("-- %s\n\n%s", filename, string(bytes)),
		)
	}
	return strings.Join(contents, "\n\n"), nil
}
