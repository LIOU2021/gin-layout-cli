package server

import (
	"fmt"
)

var status = &server{
	name: "status",
	run: func() {
		fmt.Println("server status")

		execC("ps -aux | grep gin-layout")
	},
}

func init() {
	serverCli = append(serverCli, *status)
}
