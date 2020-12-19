package classpath

import (
	"os"
	"path/filepath"
	"strings"
)

const pathListSeparator = string(os.PathListSeparator)

type Classpath struct {
	bootClasspath Entry
	extClasspath  Entry
	userClasspath Entry
}
type Entry interface {
	readClass(className string) ([]byte, Entry, error)
	String() string
}

func newEntry(path string) Entry {
	lowerPath := strings.ToLower(path)
	if strings.Contains(lowerPath, pathListSeparator) {
		return newCompositeEntry(path)
	}
	if strings.HasSuffix(lowerPath, ".zip") || strings.HasSuffix(lowerPath, ".jar") {
		return newZipEntry(path)
	}
	return newDirEntry(path)
}
func Parse(jreOption, cpOption string) *Classpath {
	cp := &Classpath{}
	cp.parseBootAndExtClasspath(jreOption)
	cp.parseUserClasspath(jreOption)
	return cp
}
func (c Classpath) parseBootAndExtClasspath(jreOption string) {
	jreDir := getJreDir(jreOption)
	//jre/lib/*
	jreLibPath := filepath.Join(jreDir, "lib", "*")
	c.bootClasspath = newWildCardEntry(jreLibPath)
	//jre/lib/exy/*
	jreExtPath := filepath.Join(jreDir, "lib", "ext", "*")
	c.extClasspath = newWildCardEntry(jreExtPath)

}
func (c *Classpath) parseUserClasspath(cpOption string) {
	if cpOption == "" {
		cpOption = "."
	}
	c.userClasspath = newEntry(cpOption)
}
func (c Classpath) ReadClass(className string) ([]byte, Entry, error) {
	className = className + ".class"
	if data, entry, err := c.bootClasspath.readClass(className); err == nil {
		return data, entry, err
	}
	if data, entry, err := c.extClasspath.readClass(className); err == nil {
		return data, entry, err
	}
	return c.userClasspath.readClass(className)
}
func (c Classpath) String() string {
	return c.userClasspath.String()
}
func getJreDir(jreOption string) string {
	if jreOption != "" && dirExist(jreOption) {
		return jreOption
	}
	if dirExist("./jre") {
		return "./jre"
	}
	if jh := os.Getenv("JAVA_HOME"); jh != "" {
		return filepath.Join(jh, "jre")
	}
	panic("cannot find jre folder")
}
func dirExist(path string) bool {
	if _, err := os.Stat(path); err != nil {
		return false
	}
	return true
}
