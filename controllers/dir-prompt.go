package controllers

import (
	"os"
	"os/user"
	"visual-file-server/rules"

	"github.com/gin-gonic/gin"
)

func init() {
	addController(dirPrompt)
}

func dirPrompt() Controller {
	handler := func(c *gin.Context) {
		var responseBody j
		defer c.Set("response", responseBody)
		var err error

		path := c.Query("path")
		if path == "" {
			cUser, err := user.Current()
			if err != nil {
				responseBody = rules.Error(c)
				return
			}
			path = cUser.HomeDir
		}

		var dir *os.File
		if dir, err = os.Open(path); err != nil {
			responseBody = rules.ErrorQuery(c)
			return
		}

		if names, err := dir.Readdirnames(-1); err != nil {
			responseBody = rules.ErrorQuery(c)
		} else {
			responseBody = rules.Success(c, j{"dir": path, "files": names})
		}
	}

	return Controller{
		Group:   []string{"/dir/"},
		URLPath: "/prompt/",
		Method:  "GET",
		Handler: handler,
	}
}
