package model

// DirInfo holds the path and total size (in bytes) of a directory.
// Size includes all files and subdirectories recursively.
type DirInfo struct {
	Path string
	Size int64
}
