package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"minigit/internal/app"
	"minigit/internal/app/plumbing"
)

var objectType string

var hashObjectCmd = &cobra.Command{
	Use:  "hash-object <file>",
	Args: cobra.ExactValidArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		data, err := ioutil.ReadFile(args[0])
		if err != nil {
			panic(err)
		}
		fmt.Println(plumbing.NewHashObject(data, app.ObjectTypeBlob))
	},
}

func init() {
	hashObjectCmd.PersistentFlags().StringVarP(&objectType, "", "t", string(app.ObjectTypeBlob), "type of object")
}
