package cmd

import (
	"github.com/spf13/cobra"
	"minigit/internal/app/plumbing"
)

var readTreeCmd = &cobra.Command{
	Use:  "read-tree",
	Args: cobra.ExactValidArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if err := plumbing.ReadTree(args[0]); err != nil {
			panic(err)
		}
	},
}
