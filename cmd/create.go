/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

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
			fmt.Println("args : ", args)
			fmt.Println("file name is : ", fileName)
			fmt.Println("create called")
		},
	}
)

func init() {
	rootCmd.AddCommand(createCmd)

	createCmd.Flags().StringVarP(&fileName, "name", "n", "", "file name")
	createCmd.MarkFlagRequired("name")

}
