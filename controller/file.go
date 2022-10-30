package controller

import (
	. "QuickShare/db"
	. "QuickShare/db/model"
	"QuickShare/utility"
	. "QuickShare/utility"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Upload(c *gin.Context) {
	temp := c.Query("temporary")
	var temporary bool
	if temp == "true" {
		temporary = true
	} else {
		temporary = false
	}
	data := File{}
	data.CreatedAt = time.Now()
	file, err := c.FormFile("file")
	if err != nil {
		log.Println(err)
		Response(c, http.StatusBadRequest, "Bad Request", nil)
		return
	}
	data.Name = file.Filename
	data.Path = GenerateFilePath(file.Filename, data.CreatedAt)
	data.Hash = GenerateFileHash(file.Filename, data.CreatedAt)
	data.Size = file.Size
	data.Type = file.Header.Get("Content-Type")
	data.Temporary = temporary
	data.ExpireAt = time.Now().AddDate(0, 0, 1)
	if err := c.SaveUploadedFile(file, data.Path); err != nil {
		log.Println(err)
		Response(c, http.StatusInternalServerError, "Internal Server Error", nil)
		return
	}
	if err := DB.Create(&data).Error; err != nil {
		log.Println(err)
		Response(c, http.StatusInternalServerError, "Internal Server Error", nil)
		return
	}
	Response(c, http.StatusOK, "OK", gin.H{
		"hash": data.Hash,
	})
}

func Download(c *gin.Context) {
	hash := c.Param("hash")
	data := File{}
	if err := DB.Where("hash = ?", hash).First(&data).Error; err != nil {
		log.Println(err)
		Response(c, http.StatusNotFound, "Not Found", nil)
		return
	}
	if data.Temporary {
		if data.ExpireAt.Before(time.Now()) {
			Response(c, http.StatusNotFound, "Not Found", nil)
			utility.DeleteFile(data.Path)
			DB.Delete(&data)
			return
		}
	}
	c.FileAttachment(data.Path, data.Name)
	data.DownloadCount++
	DB.Save(&data)
}
