// Package cmd
// Copyright © 2021 邱张华 <qiuzhanghua@icloud.com>
package cmd

import (
	"fmt"
	"github.com/Masterminds/semver/v3"
	"github.com/qiuzhanghua/autotag/tools"
	"github.com/spf13/cobra"
	"strings"
)

// nextCmd represents the next command
var nextCmd = &cobra.Command{
	Use:   "next",
	Short: "Add tag for next",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 && (args[0] == "pre" || args[0] == "phase" || args[0] == "patch" || args[0] == "minor" || args[0] == "major") {
			latest := tools.GitLatestTag()
			if len(latest) < 1 {
				latest = "0.0.0-alpha.0"
			}
			v, err := semver.NewVersion(latest)
			if err != nil {
				fmt.Printf("%s is not a Semver\n", latest)
				return
			}
			var next string
			if args[0] == "pre" {
				next = tools.IncPrerelease(*v).String()
			} else if args[0] == "phase" {
				next = tools.NextPhase(*v).String()
			} else if args[0] == "patch" {
				next = (*v).IncPatch().String()
			} else if args[0] == "minor" {
				next = (*v).IncMinor().String()
			} else if args[0] == "major" {
				next = (*v).IncMajor().String()
			}
			if strings.HasPrefix(latest, "v") {
				next = "v" + next
			}
			if next == latest {
				fmt.Printf("Add tag %s is not needed.\n", next)
				return
			}
			ret, err := tools.GitAddTag(next)
			if err != nil {
				fmt.Printf("Add tag %s error: %s\n", next, err)
			} else {
				fmt.Printf("Add tag %s ok. %s\n", next, ret)
			}
		} else {
			fmt.Print(`to do next "git tag <next_tag_name>", please use:
autotag next pre
autotag next phase
autotag next patch
autotag next minor
autotag next major
`)
		}
	},
}

func init() {
	rootCmd.AddCommand(nextCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// nextCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// nextCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
