package cmd

import (
	"log"
	"os"

	"github.com/garigari-kun/et/tmux_handler"
	"github.com/spf13/cobra"
)

func PaneLayoutCmd() *cobra.Command {
	cobra := &cobra.Command{
		Use:   "l",
		Short: "Layout current window",
		Run: func(cmd *cobra.Command, args []string) {
			sessions := tmux_handler.NewTmuxSessions()
			if sessions.IsSessionAttached() {
				layout_panes := tmux_handler.NewTmuxLayoutPanes()
				layout_panes.ListLayoutPanesToTerminal()
				choice := tmux_handler.PromptUserChoice()
				layout_panes.LayoutByChoice(choice)
			} else {
				log.Print("Session is not attached. Can't layout pane")
				os.Exit(1)
			}
		},
	}

	return cobra
}
