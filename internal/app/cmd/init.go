package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"minigit/internal/app"
	"os"
	"path"
)

var initCmd = &cobra.Command{
	Use:  "init [dir]",
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, _ []string) {
		if err := os.Mkdir(app.GIT_DIR, os.ModePerm); err != nil {
			panic(err)
		}
		if err := os.Mkdir(fmt.Sprintf("%s/objects", app.GIT_DIR), os.ModePerm); err != nil {
			panic(err)
		}

		wd, err := os.Getwd()
		if err != nil {
			panic(err)
		}
		fmt.Printf("Initialized empty ugit repository in %s/%s\n", path.Base(wd), app.GIT_DIR)
	},
}
