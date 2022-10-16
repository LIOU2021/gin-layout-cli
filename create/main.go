package create

import (
	"fmt"
	"log"
	"os"

	"github.com/LIOU2021/gin-layout-cli/helpers"
)

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

func GetAllType() []string {
	var list []string
	for _, element := range createCli {
		list = append(list, element.name)

	}
	return list
}

func createFile(fileName string, dir string, content string) {
	fileName = fileName + ".go"
	if !helpers.OsExists(dir) {
		log.Fatalf("%s not exists !", dir)
	}

	fullFileName := dir + "/" + fileName

	if helpers.OsExists(fullFileName) {
		log.Fatal(fullFileName + " already exists !")
	}

	f, err := os.Create(fullFileName)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	_, err2 := f.WriteString(content)

	if err2 != nil {
		log.Fatal(err2)
	}

	fmt.Println("done")
}
