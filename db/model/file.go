package model

import (
	"time"

	"gorm.io/gorm"
)

type FileType uint8

const (
	Image FileType = iota
	Video
	Audio
	Document
	Other
)

type File struct {
	gorm.Model
	Name          string
	Path          string
	Size          int64
	Type          FileType
	Hash          string
	DownloadCount int64
	ExpireAt      time.Time
}
