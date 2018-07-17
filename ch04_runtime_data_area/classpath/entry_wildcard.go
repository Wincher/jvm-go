package classpath

import (
  "os"
  "path/filepath"
  "strings"
)
//使用了通配符的,路径我们选择返回组合Entry
func newWildcardEntry(path string) CompositeEntry {
  baseDir := path[:len(path)-1]
  compositeEntry := []Entry{}

  walkFn := func(path string, info os.FileInfo, err error) error {
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
  }

  filepath.Walk(baseDir, walkFn)

  return compositeEntry
}
