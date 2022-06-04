package cmd

import (
	"github.com/spf13/cobra"
	"minigit/internal/app/plumbing"
)

var catFileCmd = &cobra.Command{
	Use:  "cat-file <object>",
	Args: cobra.ExactValidArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		plumbing.CatFile(args[0])
	},
}
