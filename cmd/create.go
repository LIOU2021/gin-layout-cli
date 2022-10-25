/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/LIOU2021/gin-layout-cli/create"
	"github.com/spf13/cobra"
)

var (
	fileName  string
	createCmd = &cobra.Command{
		Use:   "create [file type]",
		Short: "create some file",
		Long:  `support file type : env, helper ... `,
		Run: func(cmd *cobra.Command, args []string) {
			switch supportType {
			case true:
				supportCreateProcess()
			case false:
				createProcess(cmd, args)

			}
		},
	}
)

func init() {
	rootCmd.AddCommand(createCmd)

	createCmd.Flags().StringVarP(&fileName, "name", "n", "", "file name")
	// createCmd.MarkFlagRequired("name")

}

func createProcess(cmd *cobra.Command, args []string) {
	if len(args) >= 2 {
		log.Fatal("not support second or more args", args)
	}

	if len(args) == 0 {
		log.Fatal("lose command [file type]")
	}

	if fileName == "" {
		log.Fatal("lose '-n {fileName}'")
	}

	create.Match(args[0], fileName)
}

func supportCreateProcess() {
	fmt.Printf("support : %v\n", create.GetAllType())
}
