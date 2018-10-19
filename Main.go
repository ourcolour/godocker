package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"iamcc.cn/godocker/entities"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	var (
		host   string = "0.0.0.0"
		port   int    = 8888
		router *gin.Engine
	)

	router = gin.New()

	router.Static("/images", "/images")
	router.Static("/css", "/css")
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World!")
	})
	router.GET("/fs", func(c *gin.Context) {
		// 参数
		rootPath := c.DefaultQuery("p", "/Users/cc/Desktop")

		var (
			html string = "<html><head>{$HEAD}</head><body>{$BODY}</body></html>"
			head string = "<meta http-equiv=\"Content-Type\" content=\"text/html; charset=utf-8\" />"
			body string = ""
		)

		err := filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {
			if nil != err {
				log.Printf("Invalid root path, %s.\n", err.Error())
				return err
			}

			var line string = ""
			if info.IsDir() {
				line += "[ DIR]"
			} else {
				line += "[FILE]"
			}
			line += fmt.Sprintf(" %d \t %s<br/>", int(info.Size()), info.Name())
			log.Printf("%s\n", line)
			body += line

			return nil
		})
		html = strings.Replace(html, "{$HEAD}", head, -1)
		html = strings.Replace(html, "{$BODY}", body, -1)

		if nil != err {
			c.Error(err)
		}

		c.Header("content-type", "text/html")
		c.String(http.StatusOK, html)
	})

	router.Run(fmt.Sprintf("%s:%d", host, port))
}

func walk(rootPath string) []*entities.FileInfo {
	var result []*entities.FileInfo = make([]*entities.FileInfo, 0)

	err := filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {
		if nil != err {
			log.Printf("Invalid root path, %s.\n", err.Error())
			return err
		}

		// Build entities
		fileInfo := &entities.FileInfo{
			Dir:    filepath.Dir(path),
			Name:   info.Name(),
			IsDir:  info.IsDir(),
			IsFile: !info.IsDir(),
		}

		result = append(result, fileInfo)

		return nil
	})

	if nil != err {
		log.Fatalf("%s\n", err.Error())
	}

	return result
}
