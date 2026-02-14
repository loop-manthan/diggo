package format

import (
	"diggo/internal/model"
	"fmt"
	"sort"
	"strings"
)

// HumanSize converts bytes to a human-readable string (e.g. "1.2 MB").
func HumanSize(bytes int64) string {
	const unit = 1024
	if bytes < unit {
		return fmt.Sprintf("%d B", bytes)
	}
	div, exp := int64(unit), 0
	for n := bytes / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB", float64(bytes)/float64(div), "KMGTPE"[exp])
}

// DirList formats a slice of DirInfo as lines of "size  path", sorted by path.
func DirList(dirs []model.DirInfo) string {
	sort.Slice(dirs, func(i, j int) bool {
		return strings.Compare(dirs[i].Path, dirs[j].Path) < 0
	})
	var b strings.Builder
	for _, d := range dirs {
		b.WriteString(HumanSize(d.Size))
		b.WriteString("\t")
		b.WriteString(d.Path)
		b.WriteString("\n")
	}
	return b.String()
}
