package cmd

import (
	"fmt"
	"github.com/qiuzhanghua/autotag/tools"

	"github.com/spf13/cobra"
)

// DateCmd represents the hash command
var DateCmd = &cobra.Command{
	Use:   "date",
	Short: "date of hash",
	Long:  `Calculate date of hash.`,
	Run: func(cmd *cobra.Command, args []string) {
		var hash string
		if len(args) < 1 {
			tag := tools.GitLatestTag()
			hash = tools.GitRevOfTag(tag)
		} else {
			hash = args[0]
		}
		date := tools.GitDateOfHash(hash)
		fmt.Print(date)
	},
}
