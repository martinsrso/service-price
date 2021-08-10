package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var rootCmd = &cobra.Command{}

var (
	cfgFile string
	// restCommandFunc = restCmd.MainCommand
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
		"config file (default is service-price/config/default.yaml)",
	)
	/* rootCommand.AddCommand(restCommandFunc())
	rootCommand.AddCommand(dbCommandFunc())
	rootCommand.AddCommand(grpcCommandFunc()) */
	return rootCommand
}

func initRun(cmdd *cobra.Command, args []string) {
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

func initPreRun(cmdd *cobra.Command, args []string) {
	config := getConfig()
	logLevel := config.GetString("log.level")

	initLogger(logLevel)

	zap.S().Info("start service-price config")
}

func getConfig() *viper.Viper {
	config := viper.New()

	if cfgFile != "" {
		config.SetConfigFile(cfgFile)
	}
	config.SetConfigType("yaml")
	config.AddConfigPath(".")
	config.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	config.AutomaticEnv()

	// If a config file is found, read it in.
	if err := config.ReadInConfig(); err != nil {
		zap.S().Panicf("config file %s failed to load: %s.\n", cfgFile, err.Error())
	}

	return config
}
