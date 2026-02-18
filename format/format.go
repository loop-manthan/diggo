package format

import (
	"diggo/model"
	"fmt"
	"path/filepath"
	"sort"
	"strings"
)

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

func depthFromRoot(root, path string) int {
	rel, err := filepath.Rel(root, path)
	if err != nil || rel == ".." || strings.HasPrefix(rel, "..") {
		return -1
	}
	if rel == "." {
		return 0
	}
	return 1 + strings.Count(filepath.ToSlash(rel), "/")
}

func Tree(dirs []model.DirInfo, root string, maxDepth int) string {
	root = filepath.Clean(root)

	var filtered []model.DirInfo
	for _, d := range dirs {
		depth := depthFromRoot(root, d.Path)
		if depth >= 0 && (maxDepth < 0 || depth <= maxDepth) {
			filtered = append(filtered, d)
		}
	}

	sort.Slice(filtered, func(i, j int) bool {
		return strings.Compare(filtered[i].Path, filtered[j].Path) < 0
	})

	lastChild := make(map[string]string)
	for _, d := range filtered {
		parent := filepath.Dir(d.Path)
		if d.Path != parent {
			lastChild[parent] = d.Path
		}
	}

	var b strings.Builder
	for _, d := range filtered {
		depth := depthFromRoot(root, d.Path)
		for i := 0; i < depth; i++ {
			b.WriteString("  ")
		}
		if depth == 0 {
			if d.Path == "" || d.Path == "." {
				b.WriteString(".")
			} else {
				b.WriteString(d.Path)
			}
		} else {
			parent := filepath.Dir(d.Path)
			isLast := lastChild[parent] == d.Path
			if isLast {
				b.WriteString("└── ")
			} else {
				b.WriteString("├── ")
			}
			b.WriteString(filepath.Base(d.Path))
		}
		b.WriteString("\t")
		b.WriteString(HumanSize(d.Size))
		b.WriteString("\n")
	}
	return b.String()
}
