package create

import (
	"fmt"
	"strings"
)

var middleware = &create{
	name: "middleware",
	run: func(fileName string) {
		createFile(fileName, "./middleware", middlewareContent(fileName))
	},
}

func init() {
	createCli = append(createCli, *middleware)
}

func middlewareContent(fileName string) string {
	var originFileName = fileName
	fileName = strings.ToUpper(fileName[0:1]) + fileName[1:]

	content := `package middleware

import (
	"github.com/gin-gonic/gin"
)

func %[1]s(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"msg": "%[1]s used"})
}`

	return fmt.Sprintf(content, fileName, originFileName)
}
