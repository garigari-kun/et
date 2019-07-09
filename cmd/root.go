package cmd

import (
	"log"
	"os"
	"strings"

	"github.com/garigari-kun/et/tmux_handler"
	"github.com/spf13/cobra"
)

func RootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "et",
		Short: "List all sessions. You can choose whether you create a new session or attach existing session.",
		Run: func(cmd *cobra.Command, args []string) {
			sessions := tmux_handler.NewTmuxSessions()
			sessions.ListChoicesToTerminal()
			choice := tmux_handler.PromptUserChoice()

			if choice == "0" {
				is_attached := sessions.IsSessionAttached()
				new_session_name := tmux_handler.PromptUserToNewSessionName()
				new_session_name = strings.Replace(new_session_name, " ", "-", -1)
				if is_attached {
					tmux_handler.CreateNewSession(new_session_name)
					tmux_handler.SwitchSession(new_session_name)
				} else {
					tmux_handler.CreateAndAttachSession(new_session_name)
				}
			} else {
				is_attached := sessions.IsSessionAttached()
				session_name := sessions.FindSessionById(choice)
				if is_attached {
					tmux_handler.SwitchSession(session_name)
				} else {
					tmux_handler.AttachSession(session_name)
				}
			}
		},
	}

	cmd.AddCommand(KillCmd())
	return cmd
}

func Execute() {
	cmd := RootCmd()
	if err := cmd.Execute(); err != nil {
		log.Print(err)
		os.Exit(1)
	}
}
