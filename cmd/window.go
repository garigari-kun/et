package cmd

import (
	"log"
	"strings"

	"github.com/garigari-kun/et/tmux_handler"
	"github.com/spf13/cobra"
)

func WindowCmd() *cobra.Command {
	cobra := &cobra.Command{
		Use:   "w",
		Short: "Create tmux window",
		Run: func(cmd *cobra.Command, args []string) {
			tmux_handler.ListChoicesForWindow()
			choice := tmux_handler.PromptUserChoice()
			if choice == "0" {
				new_window_name := strings.Replace(tmux_handler.PromptUserToNewWindowName(), " ", "-", -1)
				tmux_handler.CreateAndAttachWindow(new_window_name)
			} else {
				log.Print("No function is available.")
			}
		},
	}
	return cobra
}
