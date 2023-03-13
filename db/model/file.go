package model

import (
	. "QuickShare/db"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type File struct {
	gorm.Model
	Name          string    `json:"name"`
	Path          string    `json:"path"`
	Size          int64     `json:"size"`
	Type          string    `json:"type"`
	Hash          string    `json:"hash"`
	DownloadCount int64     `json:"download_count"`
	Temporary     bool      `json:"temporary"`
	ExpireAt      time.Time `json:"expire_at"`
}

func CreateFile(file *File) error {
	return DB.Create(file).Error
}

func GetFileByHash(hash string) (File, error) {
	var file File
	err := DB.Where("hash = ?", hash).First(&file).Error
	return file, err
}

func FileDownloadCountIncrement(hash string) error {
	return DB.Model(&File{}).Where("hash = ?", hash).Update("download_count", gorm.Expr("download_count + ?", 1)).Error
}

func GetExpiredFiles() []File {
	var files []File
	DB.Where("temporary = ? AND expire_at < ?", true, time.Now()).Find(&files)
	fmt.Println(files)
	DB.Delete(&files)
	return files
}

func GetAllFiles() ([]File, error) {
	var files []File
	err := DB.Find(&files).Error
	return files, err
}

func DeleteFileByHash(hash string) error {
	return DB.Where("hash = ?", hash).Delete(&File{}).Error
}
