// Package cmd
// Copyright © 2021 邱张华 <qiuzhanghua@icloud.com>
package cmd

import (
	"fmt"
	"github.com/blang/semver/v4"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "autotag",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello")
		version, err := semver.Make("1.3.2-alpha.2+92a56b79")
		fmt.Println(version, err)
		fmt.Println(version.Pre)
		fmt.Println(version.Build)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	//rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.autotag.yaml)")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
