package classpath

import (
  "os"
  "path/filepath"
  "fmt"
)

//描述classspath结构体
type Classpath struct {
  bootClasspath   Entry
  extClasspath    Entry
  userClasspath   Entry
}

//生成classpath实体
func Parse(jreOption, cpOption string) *Classpath {
  cp := &Classpath{}
  cp.parseBootAndExtClasspath(jreOption)
  cp.parseUserClasspath(cpOption)
  fmt.Println(cp)
  return cp
}

//classpath的ReadClass方法,获取claspath下className的data数据,依次在classpath下寻找
func (self *Classpath) ReadClass(className string) ([]byte, Entry, error) {
  className = className + ".class"
  if data, entry, err := self.bootClasspath.readClass(className); err == nil {
    return data, entry, err
  }
  if data, entry, err := self.extClasspath.readClass(className); err == nil {
    return data, entry, err
  }
  return self.userClasspath.readClass(className)
}

func (self *Classpath) String() string {
  return self.userClasspath.String()
}

//根据参数初始化classpath的BootClasspath和extClasspath
func (self *Classpath) parseBootAndExtClasspath(jreOption string) {
  jreDir := getJreDir(jreOption)
  // jre/lib/*
  jreLibPath := filepath.Join(jreDir, "lib", "*")
  self.bootClasspath = newWildcardEntry(jreLibPath)

  // jre/lib/ext/*
  jreExtPath := filepath.Join(jreDir, "lib", "ext", "*")
  self.extClasspath = newWildcardEntry(jreExtPath)
}
//根据参数初始化classpath的UserClasspath
func (self *Classpath) parseUserClasspath(cpOption string) {
  if cpOption == "" {
    cpOption = "."
  }
  self.userClasspath = newEntry(cpOption)
}

//根据菜蔬获取jre目录地址
func getJreDir(jreOption string) string {

  if jreOption != "" && exists(jreOption) {
    return jreOption
  }
  if exists("./jre") {
    return "./jre"
  }
  if jh := os.Getenv("JAVA_HOME"); jh != "" {
    println("$JAVA_HOME:===>", jh)
    return filepath.Join(jh, "jre")
  }
  panic("Can not find jre folder!")
}
//判断目录是否在
func exists(path string) bool {
  if _, err := os.Stat(path); err != nil {
    if os.IsNotExist(err) {
      fmt.Println(err)
      return false
    }
  }
  return true
}
