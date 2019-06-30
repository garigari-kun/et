package cmd

import (
  "github.com/spf13/cobra"
  "log"
  "os"
  "github.com/garigari-kun/easy-tmux/tmux_handler"
)

func RootCmd() *cobra.Command {
  cmd := &cobra.Command{
    Use:   "et",
    Short: "easy-tmux",
    Run: func(cmd *cobra.Command, args []string) {
      env := os.Getenv("TMUX")
      if env == "" {
        log.Print("env load error")
      }

      sessions := tmux_handler.SetTmuxSessions()
      tmux_handler.ListChoicesToTerminal(sessions)
      choice := tmux_handler.PromptUserChoice()

      if choice == "0" {
        new_session_name := tmux_handler.PromptUserChoice()
        tmux_handler.CreateNewSession(new_session_name)  
        tmux_handler.SwitchSession(new_session_name)
      } else {
        session_name := tmux_handler.FindSessionById(sessions, choice)
        tmux_handler.SwitchSession(session_name)
      }
    },
  }
  return cmd
}

func Execute() {
  cmd := RootCmd()
  if err := cmd.Execute(); err != nil {
    log.Print(err)
    os.Exit(1)
  }
}
