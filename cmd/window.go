package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

func WindowCmd() *cobra.Command {
	cobra := &cobra.Command{
		Use:   "w",
		Short: "Tmux window",
		Run: func(cmd *cobra.Command, args []string) {
			log.Print("Window command is called")
		},
	}

	return cobra
}
