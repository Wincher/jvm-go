package main

import (
  "fmt"
  "jvm_go/ch09/classpath"
  "strings"
  "jvm_go/ch09/rtda/heap"
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
  cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)
  classLoader := heap.NewClassLoader(cp, cmd.verboseClassFlag)

  className := strings.Replace(cmd.class, ".", "/", -1)
  mainClass := classLoader.LoadClass(className)
  mainMethod := mainClass.GetMainMethod()
  if mainMethod != nil {
    interpret(mainMethod, cmd.verboseInstFlag, cmd.args)
  } else {
    fmt.Printf("Main method not found in class %s\n", cmd.class)
  }
}
