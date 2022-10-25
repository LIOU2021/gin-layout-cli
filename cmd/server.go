/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

// serverCmd represents the server command
var (
	serverCmd = &cobra.Command{
		Use:   "server [command]",
		Short: "server stop/start/restart",
		Long:  `using for gin-layout server server control`,
		Run: func(cmd *cobra.Command, args []string) {
			serverProcess(cmd, args)
			fmt.Println("server called")
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
}
