package server

import (
	"fmt"
)

var start = &server{
	name: "start",
	run: func() {
		fmt.Println("server start")

		execC("gin-layout")
	},
}

func init() {
	serverCli = append(serverCli, *start)
}
