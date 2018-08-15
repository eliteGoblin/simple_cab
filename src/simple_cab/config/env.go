package config

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

const (
	simpleCabRelativePath = "/config/simple_cab.toml"
)

func getConfigFilePath() string {
	return rootPath() + simpleCabRelativePath
}

func rootPath() string {
	if rp := os.Getenv("GOPATH"); len(rp) > 0 {
		return rp
	}

	path, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return parentDir(path)
}

func parentDir(dir string) string {
	return substr(dir, 0, strings.LastIndex(dir, "/"))
}

func substr(s string, pos, length int) string {
	runes := []rune(s)
	l := pos + length
	if l > len(runes) {
		l = len(runes)
	}
	return string(runes[pos:l])
}
