package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"minigit/internal/app/plumbing"
)

var writeTreeCmd = &cobra.Command{
	Use:  "write-tree",
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(plumbing.WriteTree("."))
	},
}
