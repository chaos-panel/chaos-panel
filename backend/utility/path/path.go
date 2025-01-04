package path

import (
	"path/filepath"
)

const (
	Separator = string(filepath.Separator)
)

// get file name from path
func GetFileName(path string) string {
	return filepath.Base(path)
}

// get file extension from path
func GetFileExt(path string) string {
	ext := filepath.Ext(path)
	// 去掉扩展名的点
	return ext[1:]
}

// get file name without extension from path
func GetFileNameWithoutExt(path string) string {
	name := filepath.Base(path)
	ext := filepath.Ext(name)
	return name[:len(name)-len(ext)]
}

// get parent directory from path
func GetParentDirectory(path string) string {
	return filepath.Dir(path)
}

// normalize path to use forward slash as separator
func NormalizePath(path string) string {
	return filepath.ToSlash(path)
}
