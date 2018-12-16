package pack

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"visual-file-server/routine"
)

// File define the file struct with file data
type File struct {
	Name string `json:"file-name"`
	Data []byte `json:"content"`
}

// Dir define the tree struct for Dir
type Dir struct {
	Name   string  `json:"name"`
	SubDir []*Dir  `json:"dirs"`
	Files  []*File `json:"files"`
}

func (point *Dir) insertSubDir(arg ...interface{}) {
	path := arg[0].(string)
	subDir, err := GetDirTree(path)
	if err != nil {
		return
	}
	point.SubDir = append(point.SubDir, subDir)
}

func (point *Dir) insertFile(arg ...interface{}) {
	path := arg[0].(string)
	file := &File{Name: filepath.Base(path)}
	fd, err := os.Open(path)
	if err != nil {
		return
	}
	defer fd.Close()
	file.Data, err = ioutil.ReadAll(fd)
	point.Files = append(point.Files, file)
}

// GetDirTree return tree struct for dir
func GetDirTree(path string) (*Dir, error) {
	// path must be abs
	if !filepath.IsAbs(path) {
		var err error
		path, err = filepath.Abs(path)
		if err != nil {
			return nil, err
		}
	}

	dir, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	root := &Dir{Name: filepath.Base(path)}

	names, err := dir.Readdirnames(-1)
	if err != nil {
		return nil, err
	}

	var wg routine.WaitGroup
	for _, name := range names {
		if strings.HasPrefix(name, ".") {
			continue
		}

		p := filepath.Join(path, name)

		var stats os.FileInfo
		if stats, err = os.Stat(p); err != nil {
			continue
		}

		if stats.IsDir() {
			wg.Go(root.insertSubDir, p)
		} else {
			wg.Go(root.insertFile, p)
		}
	}
	wg.Wait()
	return root, nil
}
