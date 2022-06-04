package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"minigit/internal/app/plumbing"
)

var hashObjectCmd = &cobra.Command{
	Use:  "hash-object <file>",
	Args: cobra.ExactValidArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		data, err := ioutil.ReadFile(args[0])
		if err != nil {
			panic(err)
		}
		fmt.Println(plumbing.NewHashObject(data))
	},
}
