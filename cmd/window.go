package cmd

import (
	"log"
	"os"

	"github.com/garigari-kun/et/tmux_handler"
	"github.com/spf13/cobra"
)

func WindowCmd() *cobra.Command {
	cobra := &cobra.Command{
		Use:   "w",
		Short: "Tmux window",
		Run: func(cmd *cobra.Command, args []string) {
			log.Print("Window command is called")
			sessions := tmux_handler.NewTmuxSessions()
			if sessions.IsSessionAttached() {
				log.Print("session is attached")
			} else {
				log.Print("Session is not attached. Can't create window")
				os.Exit(1)
			}
		},
	}

	return cobra
}
