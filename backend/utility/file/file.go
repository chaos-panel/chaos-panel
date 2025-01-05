package file

import (
	"os"

	"github.com/chaos-plus/chaos-plus/utility/path"
)

// get file size in bytes from path
func GetFileSize(path string) int64 {
	file, err := os.Open(path)
	if err != nil {
		return 0
	}
	defer file.Close()
	fileInfo, _ := file.Stat()
	return fileInfo.Size()
}

// get executable path
func GetExecutablePath() string {
	path, _ := os.Executable()
	return path
}

// get executable name
func GetExecutableName() string {
	return path.GetFileName(GetExecutablePath())
}

// get executable name without extension
func GetExecutableNameWithoutExt() string {
	return path.GetFileNameWithoutExt(GetExecutablePath())
}

// get executable directory
func GetExecutableDirectory() string {
	return path.GetParentDirectory(GetExecutablePath())
}

// is directory
func IsDirectory(path string) bool {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false
	}
	return fileInfo.IsDir()
}

// is file
func IsFile(path string) bool {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false
	}
	return !fileInfo.IsDir()
}

// mkdir all
func MkdirAll(path string) error {
	return os.MkdirAll(path, os.ModePerm)
}

// remove all
func RemoveAll(path string) error {
	return os.RemoveAll(path)
}

// remove file
func RemoveFile(path string) error {
	return os.Remove(path)
}
