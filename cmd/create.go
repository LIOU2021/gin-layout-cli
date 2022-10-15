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

// createCmd represents the create command
var (
	// Used for flags.
	fileName string

	createCmd = &cobra.Command{
		Use:   "create [file type]",
		Short: "create some file",
		Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
		Run: func(cmd *cobra.Command, args []string) {

			if len(args) >= 2 {
				log.Fatal("not support second or more args", args)
			}

			create.Match(args[0], fileName)
			fmt.Println("create called")
		},
	}
)

func init() {
	rootCmd.AddCommand(createCmd)

	createCmd.Flags().StringVarP(&fileName, "name", "n", "", "file name")
	createCmd.MarkFlagRequired("name")

}
