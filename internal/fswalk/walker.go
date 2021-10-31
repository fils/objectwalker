package fswalk

import (
	"os"
	"path/filepath"
	"sync/atomic"
)

func Walkdir(dir string) (int64, error) {
	var count int64
	filepath.Walk("./", func(root string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		atomic.AddInt64(&count, 1)
		return nil
	})

	return count, nil

}
