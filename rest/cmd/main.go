package cmd

import (
	"github.com/martinsrso/service-price/config"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

// MainCommand starts the web server supporting the REST API
func MainCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "web",
		Short: "Starts the Web server",
		Run:   webServer,
	}
}

func webServer(cmd *cobra.Command, args []string) {
	// Here is where we configure our web server and start it
	cfgFile, err := cmd.Flags().GetString("config")
	if err != nil {
		zap.S().Error(err)
	}

	config := config.GetConfig(cfgFile)

	zap.S().Infof("start web server for env: %s", config.GetString("environment"))
}
