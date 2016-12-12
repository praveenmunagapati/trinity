package intake

import (
	"os"
	"path/filepath"
)

// FileList is a list of file paths
type FileList []string

/*
Add appends a path to the FileList
*/
func (list *FileList) Add(file string) {
	*list = append(*list, file)
}

/*
GatherFiletype walks a root directory and collects all the
paths of files of a given filetype. For example, to gather
all HTML files, the function would be called:

fileList, err := GatherFiletype("some/path", ".html")
*/
func GatherFiletype(root, filetype string) (*FileList, error) {
	list := &FileList{}
	gather := func(path string, file os.FileInfo, err error) error {
		if filepath.Ext(path) == filetype {
			list.Add(path)
		}
		return nil
	}

	err := filepath.Walk(root, gather)
	if err != nil {
		return list, err
	}
	return list, nil
}
