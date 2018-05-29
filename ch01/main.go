package main

import "fmt"
import "jvmgo/ch01/data"

func main() {
	cmd := data.ParseCmd()

	if cmd.VersionFlag {
		fmt.Println("version 0.0.1")
	} else if cmd.HelpFlag || cmd.Class == "" {
		data.PrintUsage()
	} else {
		startJVM(cmd)
	}
}

func startJVM(cmd *data.Cmd) {
	fmt.Printf("classpath:%s class:%s args:%v\n",
		cmd.CpOption, cmd.Class, cmd.Args)
}

