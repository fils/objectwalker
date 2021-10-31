package fswalk

import (
	"os"
	"path/filepath"
	"sync/atomic"
)

func WalkDirNames(dir string) ([]string, error) {

	m := []string{}

	filepath.Walk(dir, func(root string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			m = append(m, root)
		}

		return nil
	})

	return m, nil

}

func WalkDir(dir string) (int64, error) {
	var count int64
	filepath.Walk(dir, func(root string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		atomic.AddInt64(&count, 1)
		return nil
	})

	return count, nil

}
