package cmd

import (
	"fmt"
	"github.com/qiuzhanghua/autotag/tools"

	"github.com/spf13/cobra"
)

// HashCmd represents the hash command
var HashCmd = &cobra.Command{
	Use:   "hash",
	Short: "hash of tag",
	Long:  `Calculate hash of tag.`,
	Run: func(cmd *cobra.Command, args []string) {
		var tag string
		if len(args) < 1 {
			tag = tools.GitLatestTag()
		} else {
			tag = args[0]
		}
		hash := tools.GitRevOfTag(tag)
		fmt.Print(hash)
	},
}
