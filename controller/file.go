package controller

import (
	"QuickShare/db/model"
	. "QuickShare/db/model"
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
	if err := model.CreateFile(&data); err != nil {
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
	data, err := model.GetFileByHash(hash)

	if (err != nil || data == File{}) {
		Response(c, http.StatusNotFound, "Not Found", nil)
		return
	}
	c.FileAttachment(data.Path, data.Name)
	if err := model.FileDownloadCountIncrement(hash); err != nil {
		log.Println(err)
	}
}
