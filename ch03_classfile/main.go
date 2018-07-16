package main

import (
  "fmt"
  "strings"
  "jvm_go/ch03_classfile/classpath"
  "jvm_go/ch03_classfile/classfile"
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

  cf := loadClass(className, cp)
  fmt.Println(cmd.class)
  printClassInfo(cf)
}
func printClassInfo(cf *classfile.ClassFile) {
  fmt.Printf("version: %v.%v\n", cf.MajorVersion(), cf.MinorVersion())
  fmt.Printf("constants count: %v\n", len(cf.ConstantPool()))
  fmt.Printf("access flags: 0x%x\n", cf.AccessFlags())
  fmt.Printf("this class: %v\n", cf.ClassName())
  fmt.Printf("super class: %v\n", cf.SuperClassName())
  fmt.Printf("interfaces: %v\n", cf.InterfaceNames())
  fmt.Printf("fields count: %v\n", len(cf.Fields()))
  for _, f := range cf.Fields() {
    fmt.Printf("  %s\n", f.Name())
  }
  fmt.Printf("methods count: %v\n", len(cf.Methods()))
  for _, m := range cf.Methods() {
    fmt.Printf("  %s\n", m.Name())
  }
}
func loadClass(className string, cp *classpath.Classpath) *classfile.ClassFile {
  //读取classname文件data
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
