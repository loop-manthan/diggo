package scan

import (
	"diggo/model"
	"os"
	"path/filepath"
)

// Dir scans the given root path and returns a slice of DirInfo,
// one entry per directory (including root), with sizes in bytes.
func Dir(root string) ([]model.DirInfo, error) {
	// Map from directory path to total size (files + subdirs)
	sizes := make(map[string]int64)

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		size := info.Size()
		if info.IsDir() {
			// Don't count directory's own metadata as content
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

	// Convert map to slice of DirInfo
	result := make([]model.DirInfo, 0, len(sizes))
	for path, size := range sizes {
		result = append(result, model.DirInfo{Path: path, Size: size})
	}
	return result, nil
}
