package server

import (
	"fmt"
)

var start = &server{
	name: "start",
	run: func() {
		fmt.Println("server start")
		execC("nohup gin-layout &>> /var/log/gin.log  &")
	},
}

func init() {
	serverCli = append(serverCli, *start)
}
