// Package cmd
// Copyright © 2021 邱张华 <qiuzhanghua@icloud.com>
package cmd

import (
	"fmt"
	"github.com/Masterminds/semver/v3"
	"github.com/qiuzhanghua/autotag/tools"
	"strings"

	"github.com/spf13/cobra"
)

// ShowCmd represents the show command
var ShowCmd = &cobra.Command{
	Use:   "show",
	Short: "show current tag and next tags",
	Long:  `show current tag and next (pre/phase/patch/minor/major) tags.`,
	Run: func(cmd *cobra.Command, args []string) {
		latest := tools.GitLatestTag()
		if len(latest) < 1 {
			fmt.Println("Tag not found/Not a git repository")
			return
		}
		v, err := semver.NewVersion(latest)
		if err != nil {
			fmt.Printf("%s is not a Semver\n", latest)
			return
		}
		if strings.HasPrefix(latest, "v") {
			fmt.Printf("Current tag:  v%s\n", *v)
			fmt.Printf("next pre   :  v%s\n", tools.IncPrerelease(*v))
			fmt.Printf("next phase :  v%s\n", tools.NextPhase(*v))
			fmt.Printf("next patch :  v%s\n", (*v).IncPatch())
			fmt.Printf("next minor :  v%s\n", (*v).IncMinor())
			fmt.Printf("next major :  v%s\n", (*v).IncMajor())
			return
		}
		fmt.Printf("Current tag:  %s\n", *v)
		fmt.Printf("next pre   :  %s\n", tools.IncPrerelease(*v))
		fmt.Printf("next phase :  %s\n", tools.NextPhase(*v))
		fmt.Printf("next patch :  %s\n", (*v).IncPatch())
		fmt.Printf("next minor :  %s\n", (*v).IncMinor())
		fmt.Printf("next major :  %s\n", (*v).IncMajor())
	},
}
