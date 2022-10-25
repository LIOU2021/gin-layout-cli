package server

import (
	"fmt"
)

var stop = &server{
	name: "stop",
	run: func() {
		fmt.Println("run server stop")

		execC("pkill gin-layout")
		execC("ps -aux | grep gin-layout")
	},
}

func init() {
	serverCli = append(serverCli, *stop)
}
