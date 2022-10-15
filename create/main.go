package create

import "log"

type runInterface func(fileName string)
type create struct {
	name string
	run  runInterface
}

var createCli []create

func Match(fileType string, fileName string) {
	exists := false
	for _, element := range createCli {
		if element.name == fileType {
			element.run(fileName)
			exists = true
			break
		}
	}

	if !exists {
		log.Fatalf("not find command file type : %s", fileType)
	}

}
