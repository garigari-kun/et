package cmd

import (
  "github.com/spf13/cobra"
  "log"
  "os"
  "bufio"
  // "strings"
  "fmt"
  "strconv"
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
      fmt.Println("Create new session or Attach another session.")
      fmt.Println("0: Create New Session")
      for _, session := range sessions {
        fmt.Println(strconv.Itoa(session.Id) + ": " + session.Name)
      }
      
      scanner := bufio.NewScanner(os.Stdin)
      fmt.Println("Enter what you want: ")
      scanner.Scan()
      text := scanner.Text()
      fmt.Println(text)
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
