package cmd

import (
	"fmt"
	"github.com/Masterminds/semver/v3"
	"github.com/qiuzhanghua/autotag/tools"
	"github.com/spf13/cobra"
	"strings"
)

// CurrentCmd represents the current version
var CurrentCmd = &cobra.Command{
	Use:   "current",
	Short: "current tag",
	Long:  `Show current tag.`,
	Run: func(cmd *cobra.Command, args []string) {
		if !tools.GitInstalled() {
			fmt.Print("0.0.0")
			return
		}
		if !tools.GitDirIsRepo(".") {
			fmt.Print("0.0.0")
			return
		}
		if len(tools.GitHeadHash()) == 0 {
			fmt.Print("0.0.0")
			return
		}

		latest := tools.GitLatestTag()
		if len(latest) < 1 {
			latest = "0.0.0"
		}
		v, err := semver.NewVersion(latest)
		if err != nil {
			fmt.Print("0.0.0")
			return
		}
		if strings.HasPrefix(latest, "v") {
			fmt.Printf("v%s", *v)
			return
		}
		fmt.Printf("%s", *v)
	},
}
