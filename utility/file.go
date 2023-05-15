package utility

import (
	. "QuickShare/config"
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
	return Config.File.AutoPath + GenerateFileHash(fileName, updateTime)
}

func RemoveFile(path string) error {
	fmt.Println("Remove file", path)
	return os.Remove(path)
}

func GenerateShareCode(hash string) string {
	// 6 digits random number
	n := 0
	for i := 0; i < len(hash); i++ {
		n += int(hash[i])
	}
	n = n % 1000000
	return fmt.Sprintf("%06d", n)
}
