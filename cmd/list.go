/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/rafamaciel/git-hooks/hooks"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List git hooks",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		src := GetHooksDir()

		list, err := hooks.LoadHooksFrom(src)

		if err != nil {
			fmt.Println(err)
			return
		}
		w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
		fmt.Fprintln(w, "Hook Kind\tStatus")
		for _, hook := range list {
			status := "disabled"
			if hook.Status == hooks.ENABLED {
				status = "enabled"
			}
			fmt.Fprintf(w, "%s\t%s\n", hook.Kind, status)
		}
		w.Flush()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
