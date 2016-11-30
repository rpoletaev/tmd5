package main

import (
	"log"

	"net/http"

	"unicode/utf8"

	"fmt"

	"crypto/md5"

	"github.com/gin-gonic/gin"
)

// import "crypto/md5"

const (
	prefix = "md5"
)

type md5source struct {
	ID   uint64 `json:"id"`
	Text string `json:"text"`
}

func main() {
	r := getEngine()
	r.Run() // listen and server on 0.0.0.0:8080
}

func getEngine() *gin.Engine {
	r := gin.Default()
	r.POST("/md5", func(c *gin.Context) {
		var source md5source
		if err := c.BindJSON(&source); err != nil {
			log.Println("INPUT PARSING ERROR: ", err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"status": "Оишбка при отправке данных"})
			return
		}

		textLen := utf8.RuneCountInString(source.Text)
		if textLen == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"status": "Поле 'текст' не может быть пустым"})
			return
		}

		if textLen > 100 {
			c.JSON(http.StatusBadRequest, gin.H{"status": "Поле 'текст' не должно содержать более 100 символов"})
			return
		}

		sh := fmt.Sprint(source.ID, source.Text, source.ID%2)
		res := md5.Sum([]byte(sh))
		c.JSON(http.StatusOK, fmt.Sprintf("%s%x", prefix, res))
	})

	return r
}
