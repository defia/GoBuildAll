package main

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

func GetParentFolder(folder string) string {
	parent := filepath.Join(folder, "..")
	if parent == folder {
		return ""
	}
	return parent
}

func IsFileExist(filename string) bool {

	stat, err := os.Stat(filename)
	if err != nil {
		return false
	}
	return !stat.IsDir()
}

func IsFolderExist(filename string) bool {

	stat, err := os.Stat(filename)
	if err != nil {
		return false
	}
	return stat.IsDir()
}

func FindFile(filename string) (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	gopath := GetGOPATH()

	for dir != "" && !gopath.Include(dir) {
		file := filepath.Join(dir, filename)
		if IsFileExist(file) {
			return file, nil
		}
		dir = GetParentFolder(dir)
	}
	return "", os.ErrNotExist

}

type PATHS []string

func GetGOPATH() PATHS {
	split := ""
	if runtime.GOOS == "windows" {
		split = ";"
	} else {
		split = ":"
	}
	strs := strings.Split(os.Getenv("GOPATH"), split)
	for i := range strs {
		strs[i] = filepath.Join(strs[i], "src")
	}
	return PATHS(strs)

}
func (p *PATHS) Include(dir string) bool {
	strs := []string(*p)

	for i := range strs {
		if strs[i] == dir {
			return true
		}
	}
	return false
}
