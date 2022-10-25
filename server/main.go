package server

import (
	"log"
	"os"
	"os/exec"
)

type runInterface func()

type server struct {
	name string
	run  runInterface
}

var serverCli []server

func Match(fileType string) {
	exists := false
	for _, element := range serverCli {
		if element.name == fileType {
			element.run()
			exists = true
			break
		}
	}

	if !exists {
		log.Fatalf("not find command : %s", fileType)
	}

}

func GetAllCommand() []string {
	var list []string
	for _, element := range serverCli {
		list = append(list, element.name)

	}
	return list
}

func execC(cli string) {
	cmd := exec.Command("/bin/bash", "-c", cli)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		log.Fatalf("failed to call cmd.Run(): %v", err)
	}
}
