package tests

import (
	"errors"
	"os"
	"path/filepath"
)

func GetProjectRoot() string {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	for {
		_, err = os.Stat(filepath.Join(path, ".env"))
		if errors.Is(err, os.ErrNotExist) {
			path = filepath.Dir(path)
			continue
		}
		if err != nil {
			panic(err)
		}

		break
	}

	return path
}
