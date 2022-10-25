package server

import (
	"log"
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
