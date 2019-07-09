package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

func WindowCmd() *cobra.Command {
	cobra := &cobra.Command{
		Use:   "w",
		Short: "Create tmux window",
		Run: func(cmd *cobra.Command, args []string) {
			log.Print("window option is called")
		},
	}
	return cobra
}
