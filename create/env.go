package create

import "fmt"

var env = &create{
	name: "env",
	run: func(fileName string) {
		fmt.Println("you choose create env file !!")
		fmt.Println("fileName : ", fileName)
	},
}

func init() {
	createCli = append(createCli, *env)
}
