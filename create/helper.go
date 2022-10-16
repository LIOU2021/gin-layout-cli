package create

import (
	"fmt"
	"strings"
)

var helper = &create{
	name: "helper",
	run: func(fileName string) {
		createFile(fileName, "./helpers", helperContent(fileName))
	},
}

func init() {
	createCli = append(createCli, *helper)
}

func helperContent(fileName string) string {
	fileName = strings.ToUpper(fileName[0:1]) + fileName[1:]

	content := `package helpers

func %[1]s()  {
}`

	return fmt.Sprintf(content, fileName)
}
