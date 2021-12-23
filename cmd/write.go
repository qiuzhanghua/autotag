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
		//currentHash, err := tools.GitCurrentHash()
		//if err != nil {
		//	fmt.Println("can't get current hash")
		//	os.Exit(1)
		//}
		//fmt.Println("current hash", currentHash)
		//fmt.Println(tools.GitCurrenTag())
		//fmt.Println(tools.GitAllTags())
		latest := tools.GetLatestTag()
		fmt.Println(latest)
		v, _ := semver.NewVersion(latest)
		fmt.Println(v)
		if len(latest) > 0 {
			fmt.Printf("next phase %s\n", tools.NextPhase(*v))
			fmt.Printf("next patch %s\n", (*v).IncPatch())
			fmt.Printf("next minor %s\n", (*v).IncMinor())
			fmt.Printf("next major %s\n", (*v).IncMajor())
		}
		//latest = "v1.2.3-beta.1+build34"
		//ver, err := semver.NewVersion(latest)
		//fmt.Println(ver, err)
		//v := ver.IncPatch()
		//fmt.Println(v)
		//v = v.IncMinor()
		//fmt.Println(v)
		//v = v.IncMajor()
		//fmt.Println(v)
		//v, err = v.SetPrerelease("alpha.1")
		//v, err = v.SetMetadata("build345")
		//fmt.Println(v, err)
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
