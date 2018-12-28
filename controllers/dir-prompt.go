package controllers

import (
	"mime"
	"os"
	"os/user"
	"path/filepath"
	"strings"
	"visual-file-server/rules"

	"github.com/gin-gonic/gin"
)

func init() {
	addController(dirPrompt)
}

type fileInfo struct {
	Name    string `json:"name"`
	Size    int64  `json:"size"`
	IsDir   bool   `json:"isDir"`
	ModTime string `json:"modTime"`
	Type    string `json:"type"`
}

func dirPrompt() Controller {
	// gather file info
	readFileInfos := func(path string, names []string) []fileInfo {
		var files = make([]fileInfo, 0, len(names))
		for _, name := range names {
			stats, err := os.Stat(filepath.Join(path, name))
			if err != nil {
				continue
			}

			isDir := stats.IsDir()

			var fType string
			if isDir {
				fType = "dir"
			} else {
				fType = mime.TypeByExtension(filepath.Ext(name))
			}

			files = append(files, fileInfo{
				Name:    name,
				Size:    stats.Size(),
				IsDir:   isDir,
				ModTime: stats.ModTime().String(),
				Type:    fType,
			})
		}
		return files
	}

	handler := func(c *gin.Context) {
		var responseBody j
		defer c.Set("response", &responseBody)
		var err error

		path := c.Query("path")
		if path[0] == '~' {
			cUser, err := user.Current()
			if err != nil {
				responseBody = rules.Error(c)
				return
			}
			path = strings.Replace(path, "~", cUser.HomeDir, 1)
		}

		var dir *os.File
		if dir, err = os.Open(path); err != nil {
			responseBody = rules.ErrorQuery(c)
			return
		}
		defer dir.Close()

		if names, err := dir.Readdirnames(-1); err != nil {
			responseBody = rules.ErrorQuery(c)
		} else {
			responseBody = rules.Success(c, j{"path": path, "files": readFileInfos(path, names)})
		}
	}

	return Controller{
		Group:   []string{"/dir/"},
		URLPath: "/prompt/",
		Method:  "GET",
		Handler: handler,
	}
}
