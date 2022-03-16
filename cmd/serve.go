/*
Copyright © 2022 António Inácio

*/
package cmd

import (
	"github.com/a-inacio/hermit-shell/internal/logger"
	"github.com/a-inacio/hermit-shell/internal/server"
	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Run server",
	Long:  `Run hermit-shell in server mode.`,
	Run: func(cmd *cobra.Command, args []string) {
		logger.Init()

		defer func() {
			_ = logger.GetLogger().Sync()
		}()

		server.ServeGrpc()
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
