package server

import (
	"fmt"
	"io"
	"strings"

	"github.com/gin-gonic/gin"
)

type ServerApi struct {
}

func (s *ServerApi) ReflectAction(c *gin.Context) {
	msg_lang := c.Request.Header.Get("Accept-Language")
	msg_lang = strings.ReplaceAll(msg_lang, "_", "-")
	if msg_lang != "en-US" {
		msg_lang = "zh-CN"
	}
	c.Request.Header.Set("Accept-Language", msg_lang)

}

func (s *ServerApi) Report(c *gin.Context) {
	body, _ := io.ReadAll(c.Request.Body)
	fmt.Println("Request :", string(body))

}
