package create

import (
	"fmt"
	"strings"
)

var env = &create{
	name: "env",
	run: func(fileName string) {
		createFile(fileName, "./env", envContent(fileName))
	},
}

func init() {
	createCli = append(createCli, *env)
}

func envContent(fileName string) string {
	fileName = strings.ToUpper(fileName[0:1]) + fileName[1:]

	content := `package env

type %[1]s struct {
}

var %[1]sSetting = &%[1]s{}

func init() {
	EnvSlice = append(EnvSlice, %[1]sSetting)
}`

	return fmt.Sprintf(content, fileName)
}
