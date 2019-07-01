package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func KillCmd() *cobra.Command {
	cobra := &cobra.Command{
		Use:   "k",
		Short: "Kill tmux session",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Kill command called")
		},
	}
	return cobra
}
