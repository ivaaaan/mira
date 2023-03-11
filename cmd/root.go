package cmd

import (
	"fmt"
	"os"

	"github.com/ivaaaan/mira/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	ConfigPath string
	Config     config.Config

	rootCmd = &cobra.Command{
		Use:   "mira",
		Short: "mira is a tool for converting markdown files to Jira tasks",
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVarP(&ConfigPath, "config", "c", "", "Path to a user config")
}

func initConfig() {
	viper.SetConfigType("toml")
	if ConfigPath != "" {
		viper.SetConfigFile(ConfigPath)
	} else {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		viper.AddConfigPath(home + "/.config/mira")
		viper.SetConfigName("config")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Fprintf(os.Stderr, "cannot use config file: %v", err)
		os.Exit(1)
	}

	if err := viper.Unmarshal(&Config); err != nil {
		fmt.Fprintf(os.Stderr, "cannot use config file: %v", err)
		os.Exit(1)
	}
}
