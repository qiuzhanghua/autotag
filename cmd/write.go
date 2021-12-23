// Package cmd
// Copyright © 2021 邱张华 <qiuzhanghua@icloud.com>
package cmd

import (
	"fmt"
	"github.com/Masterminds/semver/v3"
	"github.com/qiuzhanghua/autotag/tools"
	"github.com/spf13/cobra"
)

// writeCmd represents the write command
var writeCmd = &cobra.Command{
	Use:   "write",
	Short: "Write git info to files",
	Long: `Write git info to files of project:

for Go Project, write to autotag.go
for NodeJS Project, write to autotag.js
`,
	Run: func(cmd *cobra.Command, args []string) {
		latest := tools.GitLatestTag()
		v, _ := semver.NewVersion(latest)
		if len(latest) > 0 {
			fmt.Printf("next phase %s\n", tools.NextPhase(*v))
			fmt.Printf("next patch %s\n", (*v).IncPatch())
			fmt.Printf("next minor %s\n", (*v).IncMinor())
			fmt.Printf("next major %s\n", (*v).IncMajor())
		}
	},
}

func init() {
	rootCmd.AddCommand(writeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// writeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// writeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
