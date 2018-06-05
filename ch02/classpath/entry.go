package classpath

import (
	"os"
	"strings"
)

const pathListSeparator string = string(os.PathListSeparator)

/*
	定义Entry接口来表示类路径项
 */
type Entry interface {
	readClass(className string) ([]byte, Entry, error)
	String() string
}

func newEntry(path string) Entry {
	if strings.Contains(path, pathListSeparator) {
		return
	}
	if strings.HasSuffix(path, "*") {
		return
	}
	if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") ||
		strings.HasSuffix(path, ".zip") || strings.HasSuffix(path, ".ZIP") {
		return
	}
	return
}