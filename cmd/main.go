package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/martinsrso/service-price/config"
	restCmd "github.com/martinsrso/service-price/rest/cmd"
)

var rootCmd = &cobra.Command{}

var (
	cfgFile         string
	restCommandFunc = restCmd.MainCommand
	// dbCommandFunc   = dbCmd.MainCommand
	// grpcCommandFunc = grpcCmd.MainCommand
)

// RootCommand returns the main command, properly configured
func RootCommand() *cobra.Command {
	var rootCommand = &cobra.Command{
		Use:              "service-price",
		Short:            "service-price command-line interface",
		PersistentPreRun: initPreRun,
		Run:              initRun,
	}

	rootCommand.PersistentFlags().StringVar(
		&cfgFile,
		"config",
		"config/default.yaml",
		"",
	)
	rootCommand.AddCommand(restCommandFunc())
	/* rootCommand.AddCommand(dbCommandFunc())
	rootCommand.AddCommand(grpcCommandFunc()) */
	return rootCommand
}

func initRun(cmd *cobra.Command, args []string) {
	zap.S().Info("teste")
}

func initLogger(logLevel string) {
	atom := zap.NewAtomicLevel()

	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.TimeKey = ""

	logger := zap.New(zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderCfg),
		zapcore.Lock(os.Stdout),
		atom,
	))
	defer logger.Sync()

	if logLevel == "debug" {
		atom.SetLevel(zap.DebugLevel)
	} else {
		atom.SetLevel(zap.ErrorLevel)
	}

	zap.ReplaceGlobals(logger)
}

func initPreRun(cmd *cobra.Command, args []string) {
	config := config.GetConfig(cfgFile)
	logLevel := config.GetString("log.level")

	initLogger(logLevel)

	zap.S().Info("start service-price config")
}
