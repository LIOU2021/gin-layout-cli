/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/LIOU2021/gin-layout-cli/server"
	"github.com/spf13/cobra"
)

// serverCmd represents the server command
var (
	serverCmd = &cobra.Command{
		Use:   "server [command]",
		Short: "server stop/start/restart",
		Long:  `using for gin-layout server server control`,
		Run: func(cmd *cobra.Command, args []string) {
			switch supportType {
			case true:
				supportServerProcess()
			case false:
				serverProcess(cmd, args)
			}
		},
	}
)

func init() {
	rootCmd.AddCommand(serverCmd)
}

func serverProcess(cmd *cobra.Command, args []string) {
	ln := len(args)

	if ln < 1 {
		log.Fatalln("lose command, like start/stop/restart")
	}

	if ln >= 2 {
		log.Fatalln("too many command")
	}

	server.Match(args[0])
}

func supportServerProcess() {
	fmt.Printf("support : %v\n", server.GetAllCommand())
}
