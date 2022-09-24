/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"test_go_project/pkg/logger"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// Used for flags.
	cfgFile string
	env     string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "test_go_project",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Run...")
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is ./configs/.config.yaml)")
	rootCmd.PersistentFlags().StringVar(&env, "env", "", "environment (dev|prod)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		// home, err := os.UserHomeDir()
		// cobra.CheckErr(err)
		// Search config in home directory with name ".cobra" (without extension).
		viper.AddConfigPath("./configs")
		viper.SetConfigType("env")
		viper.SetConfigName("config")
	}

	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err == nil {
		logger.Info.Println("Using config file:", viper.ConfigFileUsed())
	} else {
		logger.Error.Fatalln("Load config failed:", viper.ConfigFileUsed(), "\n", err)
	}
	if env != "" {
		viper.SetConfigName("config." + env)
		if err := viper.MergeInConfig(); err == nil {
			logger.Info.Println("Override config file with:", viper.ConfigFileUsed())
		} else {
			logger.Error.Fatalln("Override config failed:", viper.ConfigFileUsed(), "\n", err)
		}
	}
}
func ConfigFile() string {
	return cfgFile
}
func ENV() string {
	return env
}
