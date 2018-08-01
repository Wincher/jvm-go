package classpath

import (
  "os"
  "strings"
)

const pathListSeparator = string(os.PathListSeparator)

//操作classpath获取data的基类
type Entry interface {
  readClass(className string) ([]byte, Entry, error)
  String() string
}
//根据不同的path返回不同的子Entry
func newEntry(path string) Entry {
  if strings.Contains(path, pathListSeparator) {
    return newCompositeEntry(path)
  }
  if strings.HasSuffix(path, "*") {
    return newWildcardEntry(path)
  }
  if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") ||
     strings.HasSuffix(path, ".zip") || strings.HasSuffix(path, ".ZIP") {
       return newZipEntry(path)
  }
  return newDirEntry(path)
}
