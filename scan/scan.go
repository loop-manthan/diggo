package scan

import (
	"diggo/model"
	"os"
	"path/filepath"
)

func Dir(root string) ([]model.DirInfo, error) {
	sizes := make(map[string]int64)

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		size := info.Size()
		if info.IsDir() {
			size = 0
		}
		dir := filepath.Dir(path)
		for {
			sizes[dir] += size
			if dir == root || dir == "." || dir == filepath.Dir(dir) {
				break
			}
			dir = filepath.Dir(dir)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	result := make([]model.DirInfo, 0, len(sizes))
	for path, size := range sizes {
		result = append(result, model.DirInfo{Path: path, Size: size})
	}
	return result, nil
}
