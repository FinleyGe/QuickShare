package utility

import (
	. "QuickShare/config"
	"QuickShare/db/model"
	"crypto/md5"
	"fmt"
	"os"
	"time"
)

func GenerateFileHash(fileName string, updateTime time.Time) string {
	data := []byte(fileName + updateTime.String())
	return fmt.Sprintf("%x", md5.Sum(data))
}

func GenerateFilePath(fileName string, updateTime time.Time) string {
	return "./upload/" + GenerateFileHash(fileName, updateTime)
}

func DeleteFile(path string) error {
	return os.Remove(path)
}

func cleanTheFiles() {
	files := model.GetExpiredFiles()
	for _, file := range files {
		DeleteFile(file.Path)
	}
}

func CleanTheFiles() {
	if Config.Dev {
		go cleanTheFiles()
	}
}
