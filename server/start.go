package server

import (
	"fmt"
)

var start = &server{
	name: "start",
	run: func() {
		fmt.Println("server start")

		execC("nohup gin-layout &> /dev/null  &")
	},
}

func init() {
	serverCli = append(serverCli, *start)
}
