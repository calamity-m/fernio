package cmd

import (
	"fmt"
	"log/slog"
	"os"
	"strings"

	"github.com/calamity-m/fernio/pkg/logging"
	"github.com/calamity-m/fernio/pkg/server"
	"github.com/calamity-m/fernio/recorder/api"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgPath string

	cfgOverwrite bool

	debug bool

	rootCmd = &cobra.Command{
		Use:   "recordersrv",
		Short: "Start the recorder server",
		Long:  "The recorder server provides an API for interacting with persistent storage for recording information.",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Starting recorder server")

			// Initialize configuration directory
			dir, err := configDir()
			if err != nil {
				fmt.Printf("Failed to get config directory: %v\n", err)
				os.Exit(1)
			}

			// Create viper object for configuration
			vip := viper.New()
			vip.SetConfigName("recorder")
			vip.SetConfigType("yaml")
			vip.AddConfigPath(dir)
			vip.SetEnvPrefix("RECORDER")
			vip.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
			vip.AutomaticEnv()

			if cfgPath != "" {
				vip.SetConfigFile(cfgPath)
			}

			// Get server configuration
			serverCfg, err := server.NewViperConfig(vip, cfgOverwrite)
			if err != nil {
				fmt.Printf("Failed to load configuration: %v\n", err)
				os.Exit(1)
			}

			// Get logging configuration
			loggingCfg, err := logging.NewFromViper(vip, cfgOverwrite)
			if err != nil {
				fmt.Printf("Failed to load configuration: %v\n", err)
				os.Exit(1)
			}

			// Override level depending on flag
			if debug {
				loggingCfg.Level = slog.LevelDebug
			}

			// Initialize our logger
			logger, err := logging.New(
				logging.WithConfig(*loggingCfg),
				logging.WithEnvironment(serverCfg.Environment),
				logging.WithSystem(serverCfg.System),
				logging.WithRequestIdHeader(serverCfg.RequestIdHeader))

			if err != nil {
				fmt.Printf("Failed to create logger: %v", err)
				os.Exit(1)
			}
			logger.Info("Initialized logger")

			// Display info on our configurations
			logger.Info(fmt.Sprintf("Using the following logging config: %v", loggingCfg))
			logger.Info(fmt.Sprintf("Using the following server config: %v", serverCfg))

			// Initialize our server
			server := server.New(
				server.WithConfig(*serverCfg),
				server.WithLogger(logger))

			api.Serve(server)

		},
	}
)

func configDir() (string, error) {
	cfgDir, err := os.UserConfigDir()
	if err != nil {
		fmt.Printf("Failed to find user config directory %v", err)
		return "", err
	}

	cfgDir = cfgDir + "/fern"
	if err := os.MkdirAll(cfgDir, 0766); err != nil {
		fmt.Printf("Failed to create config directory")
		return "", err
	}

	return cfgDir, nil
}

func Initialize() {
	rootCmd.PersistentFlags().StringVarP(&cfgPath, "config", "c", "", "Config Path")
	rootCmd.PersistentFlags().BoolVarP(&debug, "debug", "d", false, "Enable debug - overrides any configuration set")
	rootCmd.PersistentFlags().BoolVar(&cfgOverwrite, "config-overwrite", false,
		"Override config file on disk when loading configuration. Will load new options into disk")
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Printf("Failed to execute: %v\n", err)
		os.Exit(1)
	}
}
