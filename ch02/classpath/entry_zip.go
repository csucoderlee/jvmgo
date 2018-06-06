package classpath

import (
	"archive/zip"
	"errors"
	"io/ioutil"
	"path/filepath"
)

/*
	读取zip文件
 */
type ZipEntry struct {
	absPath string
}

func newZipEntry(path string) *ZipEntry {
	absPath, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &ZipEntry{absPath}
}

func (self *ZipEntry) String() string {
	return self.absPath
}

func (self *ZipEntry) readClass(className string) ([]byte, Entry, error) {
	r, err := zip.OpenReader(self.absPath)
	if err != nil {
		return nil, nil, err
	}

	defer r.Close()
	for _, f:= range r.File {
		if f.Name == className {
			readCloser, error := f.Open()
			if error != nil {
				return nil, nil, error
			}

			defer readCloser.Close()
			data, error := ioutil.ReadAll(readCloser)
			if error != nil {
				return nil, nil, error
			}
			return data, self, nil
		}
	}
	return nil, nil, errors.New("class not fount: " + className)
}
