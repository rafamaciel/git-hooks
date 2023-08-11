package cmd

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/rafamaciel/git-hooks/hooks"
	"github.com/spf13/cobra"
)

var enableCmd = &cobra.Command{
	Use:   "enable",
	Short: "Enable a Git Hook",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		name := strings.Join([]string{args[0], "sample"}, ".")
		path := GetHooksDir()
		hook, err := hooks.NewHook(filepath.Join(path, name))
		if err != nil {
			return
		}

		hook.Enable()
		fmt.Println("hook enabled")
	},
}

func init() {
	rootCmd.AddCommand(enableCmd)
}
