package main

import (
  "fmt"
  "jvm_go/ch05_instructions/classpath"
  "strings"
  "jvm_go/ch05_instructions/classfile"
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
  className := strings.Replace(cmd.class, ".", "/", -1)
  cf := loadClass(className, cp)
  mainMethod := getMainMethod(cf)
  if mainMethod != nil {
    interpret(mainMethod)
  } else {
    fmt.Println("Main method not found in class %s\n", cmd.class)
  }
}

func loadClass(className string, cp *classpath.Classpath) *classfile.ClassFile {
  classData, _, err := cp.ReadClass(className)
  if err != nil {
    panic(err)
  }

  cf, err := classfile.Parse(classData)
  if err != nil {
    panic(err)
  }

  return cf
}

func getMainMethod(cf *classfile.ClassFile) *classfile.MemberInfo {
  for _, m := range cf.Methods() {
    if m.Name() == "main" && m.Descriptor() == "([Ljava/lang/String;)V" {
      return m
    }
  }
  return nil
}