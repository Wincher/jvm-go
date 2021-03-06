package main

import (
  "fmt"
  "flag"
  "os"
)

//描述Cmd命令的结构体
type Cmd struct {
  helpFlag          bool
  versionFlag       bool
  cpOption          string
  XjreOption        string
  class             string
  args              []string
}
//初始化根据命令参数生成Cmd对象
func parseCmd() *Cmd {
  cmd := &Cmd{}

  //flag类帮助绑定命令行参数到Cmd结构体的字段
  flag.Usage = printUsage
  flag.BoolVar(&cmd.helpFlag, "help", false, "print help message")
  flag.BoolVar(&cmd.helpFlag, "?", false, "print help message")
  flag.BoolVar(&cmd.versionFlag, "version", false, "print version and exit")
  flag.StringVar(&cmd.cpOption, "classpath", "", "classpath")
  flag.StringVar(&cmd.cpOption, "cp", "", "classpath")
  flag.StringVar(&cmd.XjreOption, "Xjre", "", "path to jre")
  flag.Parse()

  args := flag.Args()
  if (len(args) > 0) {
    cmd.class = args[0]
    cmd.args = args[1:]
  }

  return cmd
}
//打印帮助信息,显示参数
func printUsage() {
  fmt.Printf("Usage: %s [-options] class [args...]\n", os.Args[0])
}
