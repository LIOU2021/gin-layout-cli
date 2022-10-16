package create

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/LIOU2021/gin-layout-cli/helpers"
)

var helper = &create{
	name: "helper",
	run: func(fileName string) {

		createHelper(fileName)
	},
}

func init() {
	createCli = append(createCli, *helper)
}

func createHelper(fileName string) {
	originFileName := fileName
	fileName = fileName + ".go"
	if !helpers.OsExists("./helpers") {
		log.Fatal("./helpers not exists !")
	}

	fullFileName := "./helpers/" + fileName

	if helpers.OsExists(fullFileName) {
		log.Fatal(fullFileName + " already exists !")
	}

	f, err := os.Create(fullFileName)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	_, err2 := f.WriteString(helperContent(originFileName))

	if err2 != nil {
		log.Fatal(err2)
	}

	fmt.Println("done")
}

func helperContent(fileName string) string {
	fileName = strings.ToUpper(fileName[0:1]) + fileName[1:]

	content := `package helpers

func %[1]s()  {
}`

	return fmt.Sprintf(content, fileName)
}
