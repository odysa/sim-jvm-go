package classpath

import (
	"errors"
	"strings"
)

type CompositeEntry []Entry

func newCompositeEntry(pathList string) CompositeEntry {
	var compositeEntry []Entry

	for _, path := range strings.Split(pathList, pathListSeparator) {
		compositeEntry = append(compositeEntry, newEntry(path))
	}
	return compositeEntry
}

func (c CompositeEntry) readClass(className string) ([]byte, Entry, error) {
	for _, entry := range c {
		data, res, err := entry.readClass(className)
		if err == nil {
			return data, res, nil
		}
	}
	return nil, nil, errors.New("target class not found")
}
func (c CompositeEntry) String() string {
	var res []string
	for _, entry := range c {
		res = append(res, entry.String())
	}
	return strings.Join(res, pathListSeparator)
}
