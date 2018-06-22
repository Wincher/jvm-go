package main

import (
  "fmt"
  "strings"
  "jvm_go/ch02_classpath/classpath"
)

func main() {
  //根据命令行参数组装Cmd结构体
  cmd := parseCmd()
  //根据参数做不通操作
  if cmd.versionFlag {
    fmt.Println("Version 0.0.1")
  } else if cmd.helpFlag || "" == cmd.class {
    printUsage()
  } else {
    startJVM(cmd)
  }
}

func startJVM(cmd *Cmd) {
  //根据参数的Xjre和Classpath获取classpath结构体
  cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)
  fmt.Printf("classpath:%v class:%v args:%v\n", cp, cmd.class, cmd.args)
  //classname 格式转换
  className := strings.Replace(cmd.class, ".", "/", -1)
  //读取classname文件data
  classData, _, err := cp.ReadClass(className)
  if err != nil {
    fmt.Print("Could not find or load main class %s\n", cmd.class)
    return
  }
  fmt.Printf("class data:%v\n", classData)
}
