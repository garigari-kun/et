package cmd

import (
	"github.com/garigari-kun/et/tmux_handler"
	"github.com/spf13/cobra"
)

func KillCmd() *cobra.Command {
	cobra := &cobra.Command{
		Use:   "k",
		Short: "Kill tmux session",
		Run: func(cmd *cobra.Command, args []string) {
			sessions := tmux_handler.SetTmuxSessions()
			tmux_handler.ListTmuxSessionsForKilling(sessions)
			choice := tmux_handler.PromptUserChoice()
			session_name := tmux_handler.FindSessionById(sessions, choice)
			tmux_handler.KillSession(session_name)
		},
	}
	return cobra
}
