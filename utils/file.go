package utils

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

func OpenOrCreateFile(filename string) *os.File {
	if checkFileIsExist(filename) {
		f, _ := os.OpenFile(filename, os.O_APPEND, 0666)
		return f
	} else {
		f, _ := os.Create(filename)
		return f
	}
}

func GetAppPath() string {
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	index := strings.LastIndex(path, string(os.PathSeparator))

	return path[:index+1]
}
func GetAppPathUpper() string {
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	index := strings.LastIndex(path, string(os.PathSeparator))
	index = strings.LastIndex(path[:index], string(os.PathSeparator))
	return path[:index+1]
}
