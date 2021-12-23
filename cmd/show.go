// Package cmd
// Copyright © 2021 邱张华 <qiuzhanghua@icloud.com>
package cmd

import (
	"fmt"
	"github.com/Masterminds/semver/v3"
	"github.com/qiuzhanghua/autotag/tools"

	"github.com/spf13/cobra"
)

// showCmd represents the show command
var showCmd = &cobra.Command{
	Use:   "show",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		latest := tools.GetLatestTag()
		if len(latest) < 1 {
			fmt.Println("Tag not found/Not a git repository")
			return
		}
		v, err := semver.NewVersion(latest)
		if err != nil {
			fmt.Printf("%s is not a Semver\n", latest)
			return
		}
		fmt.Printf("Current tag: %s\n", *v)
		fmt.Printf("Next phase:  %s\n", tools.NextPhase(*v))
		fmt.Printf("Next patch:  %s\n", (*v).IncPatch())
		fmt.Printf("Next minor:  %s\n", (*v).IncMinor())
		fmt.Printf("Next major:  %s\n", (*v).IncMajor())
	},
}

func init() {
	rootCmd.AddCommand(showCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// showCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// showCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
