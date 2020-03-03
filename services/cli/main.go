package main

import (
	"os"
	"superdecimal/gmicro/services/cli/cmd"
	"superdecimal/gmicro/services/cli/config"

	"gopkg.in/abiosoft/ishell.v2"
)

func main() {
	shell := ishell.New()
	shell.Println("GMicro CLI Tool")

	// Keeps the history between runs
	shell.SetHomeHistoryPath(".gmicro-cli")

	conf, err := config.Read()
	if err != nil {
		shell.Println(err)
	}

	shell.AddCmd(cmd.CalcCommands(conf))

	// when started with "exit" as first argument, assume non-interactive execution
	if len(os.Args) > 1 && os.Args[1] == "exit" {
		if err := shell.Process(os.Args[2:]...); err != nil {
			shell.Println(err)
		}
	} else {
		// start shell
		shell.Run()
		// teardown
		shell.Close()
	}
}
