package create

import (
	"fmt"
	"strings"
)

var controller = &create{
	name: "controller",
	run: func(fileName string) {
		createFile(fileName, "./controller", controllerContent(fileName))
	},
}

func init() {
	createCli = append(createCli, *controller)
}

func controllerContent(fileName string) string {
	var originFileName = fileName
	fileName = strings.ToUpper(fileName[0:1]) + fileName[1:]

	content := `package controller

import (
	"net/http"
	
	"github.com/gin-gonic/gin"
)
	
type %[2]s struct{}
	
var %[1]s = &%[2]s{}
	
func (controller *%[2]s) Index(c *gin.Context) {
	c.JSON(http.StatusOK, "%[2]s @Index !")
}
	`

	return fmt.Sprintf(content, fileName, originFileName)
}
