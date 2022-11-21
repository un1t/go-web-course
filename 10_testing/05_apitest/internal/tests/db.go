package tests

import (
	"io/ioutil"
	"path/filepath"
	"strings"

	"gorm.io/gorm"
)

func MigrateUp(db *gorm.DB) error {
	pattern := filepath.Join(GetProjectRoot(), "migrations", "*.up.sql")
	sql, err := ConcatMigrations(pattern)
	if err != nil {
		return err
	}

	err = db.Exec(sql).Error
	if err != nil {
		MigrateDown(db)
		err = db.Exec(sql).Error
		if err != nil {
			return err
		}
	}

	return nil
}

func MigrateDown(db *gorm.DB) error {
	pattern := filepath.Join(GetProjectRoot(), "migrations", "*.down.sql")
	sql, err := ConcatMigrations(pattern)
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
		contents = append(contents, string(bytes))
	}
	return strings.Join(contents, "\n\n"), nil
}
