/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"
	"strconv"
	"wg/wg"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "automatically modify wg port",
	Long:  `automatically modify wg port`,
	Run: func(cmd *cobra.Command, args []string) {

		mode := gin.DebugMode
		if os.Getenv("GIN_MODE") == "release" {
			mode = gin.ReleaseMode
		}
		gin.SetMode(mode)

		// create gin http server listen on port :8080
		router := gin.Default()
		router.GET("/", func(c *gin.Context) {
			port, err := wg.GetPortNumber()
			if err != nil {
				c.String(500, err.Error())
				return
			}

			// convert port to string
			portStr := strconv.Itoa(int(port))
			c.String(200, portStr)
		})

		router.POST("/", func(c *gin.Context) {
			port, err := wg.IncrPortNumber()
			if err != nil {
				c.String(500, err.Error())
				return
			}

			if err := wg.RestartService(); err != nil {
				c.String(500, err.Error())
				return
			}

			// convert port to string
			portStr := strconv.Itoa(int(port))
			c.String(200, portStr)
		})

		// listen and serve
		router.Run(":8080")
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serverCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serverCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
