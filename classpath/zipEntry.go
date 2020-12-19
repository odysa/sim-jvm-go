package classpath

import (
	"archive/zip"
	"io/ioutil"
	"path/filepath"
)

type ZipEntry struct {
	absDir string
}

func newZipEntry(path string) *ZipEntry {
	absDir, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &ZipEntry{
		absDir,
	}
}
func (z ZipEntry) String() string {
	return z.absDir
}
/*
	Read from zip file
 */
func (z ZipEntry) readClass(className string) ([]byte, Entry, error) {
	f, err := zip.OpenReader(className)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	// iterate through files in the zip
	for _, f := range f.File {
		if f.Name == className {
			res, err := f.Open()
			if err != nil {
				panic(err)
			}
			defer res.Close()
			data, err := ioutil.ReadAll(res)
			if err != nil {
				panic(err)
			}
			return data, z, nil
		}
	}
	return nil, nil, err
}
