package main

import (
	"log"
	"os"
	"os/exec"
)

func main() {
	log.Println("CLI arguments are:", os.Args)
	run(os.Args[1], os.Args[2:])
}

func run(command string, cmdArgs []string) {
	var cmd *exec.Cmd
	if len(cmdArgs) == 0 {
		cmd = exec.Command(command)
	} else {
		cmd = exec.Command(command, cmdArgs...)
	}

	stdout, err := cmd.Output()
	if err != nil {
		panic(err)
	}
	println(string(stdout))
}
