package create

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/LIOU2021/gin-layout-cli/helpers"
)

var env = &create{
	name: "env",
	run: func(fileName string) {
		createEnv(fileName)
	},
}

func init() {
	createCli = append(createCli, *env)
}

func createEnv(fileName string) {
	originFileName := fileName
	fileName = fileName + ".go"
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

	_, err2 := f.WriteString(envContent(originFileName))

	if err2 != nil {
		log.Fatal(err2)
	}

	fmt.Println("done")
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
