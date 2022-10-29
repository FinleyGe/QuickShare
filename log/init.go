package log

import (
	"io"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

var LogMultiWriter io.Writer

func init() {
	fileName := "/var/log/QuickShare.log"
	f, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)
	if err != nil {
		return
	}
	LogMultiWriter = io.MultiWriter(os.Stdout, f)
	log.SetOutput(LogMultiWriter)
	gin.DefaultWriter = LogMultiWriter
}
