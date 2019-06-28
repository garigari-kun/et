package cmd

import (
  "github.com/spf13/cobra"
  "fmt"
  "log"
  "os"
  "bufio"
  "os/exec"
  "strings"
)

func RootCmd() *cobra.Command {
  cmd := &cobra.Command{
    Use:   "et",
    Short: "easy-tmux",
    Run: func(cmd *cobra.Command, args []string) {
      out, err := exec.Command("sh", "-c", "tmux list-sessions | awk '{print $1}'").Output()
      s_out := strings.Split(string(out), ":")
      if err != nil {
        log.Print(err)
      }

      fmt.Println(s_out)

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
