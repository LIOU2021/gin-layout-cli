package create

import (
	"fmt"

	"github.com/LIOU2021/gin-layout-cli/helpers"
)

var env = &create{
	name: "env",
	run: func(fileName string) {
		fmt.Println("current Dir : ", helpers.Pwd())
		fmt.Println("you choose create env file !!")
		fmt.Println("fileName : ", fileName)
	},
}

func init() {
	createCli = append(createCli, *env)
}
