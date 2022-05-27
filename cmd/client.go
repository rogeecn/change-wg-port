/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"wg/wg"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// clientCmd represents the client command
var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "change local wg conf",
	Long:  `change local wg conf`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := wg.GetEndpointPort(); err != nil {
			logrus.Error(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(clientCmd)
}
