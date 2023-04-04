package controller

import (
	"QuickShare/db/model"
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
	data := model.File{}
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

	if (err != nil || data == model.File{}) {
		Response(c, http.StatusNotFound, "Not Found", nil)
		return
	}
	c.FileAttachment(data.Path, data.Name)
	if err := model.FileDownloadCountIncrement(hash); err != nil {
		log.Println(err)
	}
}

func GetFileInfo(c *gin.Context) {
	hash := c.Param("hash")
	data, err := model.GetFileByHash(hash)
	if (err != nil || data == model.File{}) {
		Response(c, http.StatusNotFound, "Not Found", nil)
		return
	}
	Response(c, http.StatusOK, "OK", gin.H{
		"data": data,
	})
}

func GetAllInfo(c *gin.Context) {
	files, err := model.GetAllFiles()
	if err != nil {
		Response(c, http.StatusInternalServerError, "Internal Server Error", nil)
		return
	}
	Response(c, http.StatusOK, "OK", gin.H{
		"data": files,
	})
}

func DeleteFile(c *gin.Context) {
	hash := c.Param("hash")
	data, err := model.GetFileByHash(hash)
	if (err != nil || data == model.File{}) {
		Response(c, http.StatusNotFound, "Not Found", nil)
		return
	}
	if err := model.DeleteFileByHash(hash); err != nil {
		Response(c, http.StatusInternalServerError, "Internal Server Error", nil)
		return
	}
	err = RemoveFile(data.Path)
	if err != nil {
		log.Panicln(err)
	}

	Response(c, http.StatusOK, "OK", nil)
}

func GetAllInfoByType(c *gin.Context) {
	fileType := c.Query("type")
	files, err := model.GetAllInfoByType(fileType)
	if err != nil {
		Response(c, http.StatusInternalServerError, "Internal Server Error", nil)
		return
	}
	Response(c, http.StatusOK, "OK", gin.H{
		"data": files,
	})
}

func SearchFile(c *gin.Context) {
	data := c.Query("data")
	files, err := model.SearchFile(data)
	if err != nil {
		Response(c, http.StatusInternalServerError, "Internal Server Error", nil)
		return
	}
	Response(c, http.StatusOK, "OK", gin.H{
		"data": files,
	})
}
