package create

import (
	"fmt"
	"log"
	"os"

	"github.com/LIOU2021/gin-layout-cli/helpers"
)

var env = &create{
	name: "env",
	run: func(fileName string) {
		fmt.Println("current Dir : ", helpers.Pwd())
		createEnv(fileName + ".go")
		fmt.Println("you choose create env file !!")
	},
}

func init() {
	createCli = append(createCli, *env)
}

func createEnv(fileName string) {
	if !helpers.OsExists("./env") {
		log.Fatal("./env not exists !")
	}

	fullFileName := "./env/" + fileName

	if helpers.OsExists(fullFileName) {
		log.Fatal(fullFileName + " already exists !")
	}

	f, err := os.Create(fullFileName)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	_, err2 := f.WriteString("old falcon\n")

	if err2 != nil {
		log.Fatal(err2)
	}

	fmt.Println("done")
}
