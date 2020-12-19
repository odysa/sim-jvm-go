package classpath

import (
	"os"
	"path/filepath"
	"strings"
)

func newWildCardEntry(path string) CompositeEntry {
	// remove last *
	baseDir := path[:len(path)-1]
	compositeEntry := CompositeEntry{}

	//iterate through files and dir
	filepath.Walk(baseDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() && path != baseDir {
			return filepath.SkipDir
		}
		if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") {
			jarEntry := newZipEntry(path)
			compositeEntry = append(compositeEntry, jarEntry)

		}
		return nil
	})
	return compositeEntry
}

