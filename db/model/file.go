package model

import (
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
