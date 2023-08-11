package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/rafamaciel/git-hooks/hooks"
	"github.com/spf13/cobra"
)

var disableCmd = &cobra.Command{
	Use:   "disable",
	Short: "Disable a Git Hook",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		path := GetHooksDir()
		hook, err := hooks.NewHook(filepath.Join(path, args[0]))
		if err != nil {
			return
		}

		hook.Disable()
		fmt.Println("hook disbled")
	},
}

func init() {
	rootCmd.AddCommand(disableCmd)
}
