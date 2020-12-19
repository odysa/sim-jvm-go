package classpath

import (
	"os"
	"strings"
)

const pathListSeparator = string(os.PathListSeparator)

type Entry interface {
	readClass(className string) ([]byte, Entry, error)
	String() string
}

func newEntry(path string) Entry {
	lowerPath := strings.ToLower(path)
	if strings.Contains(lowerPath, pathListSeparator) {
		return newCompositeEntry(path)
	}
	if strings.Contains(lowerPath, ".zip") || strings.Contains(lowerPath, ".jar") {
		return newZipEntry(path)
	}
	return newDirEntry(path)
}
