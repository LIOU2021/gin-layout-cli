package create

import (
	"fmt"
	"log"

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

	if helpers.OsExists("./env/" + fileName) {
		log.Fatal("./env/" + fileName + " already exists !")
	}

	fmt.Println("create file")
}
