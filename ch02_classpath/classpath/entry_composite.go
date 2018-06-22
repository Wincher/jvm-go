package classpath

import (
  "errors"
  "strings"
)

//组合Entry,包含多种操作classpath的Entry
type CompositeEntry []Entry

//根据path列表,构建组合Entry
func newCompositeEntry(pathList string) CompositeEntry {
  compositeEntry := []Entry{}
  for _, path := range strings.Split(pathList, pathListSeparator) {
    entry := newEntry(path)
    compositeEntry = append(compositeEntry, entry)
  }

  return compositeEntry
}
//读取className类文件的data
func (self CompositeEntry) readClass(className string) ([]byte, Entry, error) {
  for _, entry := range self {
    data, from, err := entry.readClass(className)
    if err == nil {
      return data, from, nil
    }
  }
  return nil, nil, errors.New("class not found: " + className)
}

func (self CompositeEntry) String() string {
  strs := make([]string, len(self))
  for i, entry := range self {
    strs[i] = entry.String()
  }
  return strings.Join(strs, pathListSeparator)
}
