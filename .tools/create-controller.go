package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var log = fmt.Println

func isFileExist(path string) bool {
	_, err := os.Stat(path)
	return !(err != nil && os.IsNotExist(err))
}

func template(name, groupPath, urlPath string) string {
	return `package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {
	addController(` + name + `)
}

func ` + name + `() Controller {

	handler := func(c *gin.Context) {
		responseBody := j{"code": 0, "message": "success"}
		c.Set("response body", responseBody)
		c.JSON(http.StatusOK, responseBody)
	}

	return Controller{
		Group:   []string{"` + groupPath + `"},
		URLPath: "` + urlPath + `",
		Method:  "GET",
		Handler: handler,
	}
}	
	`
}

func main() {
	if !isFileExist("./controllers") {
		log("this dir doesn't seem to be a gat project")
		return
	}

	var name string
	flag.StringVar(&name, "n", "", "gat will use this name as file name along with function name")
	var force bool
	flag.BoolVar(&force, "f", false, "force to create controller whether there already have such controller or not")
	var groupPath string
	flag.StringVar(&groupPath, "g", "", "gat will use this group name as controller group path")
	var urlPath string
	flag.StringVar(&urlPath, "u", "", "gat will use this path as request path")
	flag.Parse()

	controllerName := strings.TrimRight(name, ".go") // rm golang suffix
	if controllerName == "" {
		log("you have to point out what controller name you wanto use \n use -h option to see what you can do")
		return
	}

	name = filepath.Join("./controllers", controllerName+".go")
	if isFileExist(name) && !force {
		log("file " + name + " already exist , you need to use another file name or use -f option to recreate this file")
		return
	}

	file, err := os.Create(name)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	file.WriteString(template(controllerName, groupPath, urlPath))
}
