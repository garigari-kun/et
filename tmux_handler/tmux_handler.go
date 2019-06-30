package tmux_handler

import (
  "log"
  "os"
  "os/exec"
  "strings"
)

type Session struct {
  Id int
  Name string
  Attached string 
} 

type Sessions []Session

func SetTmuxSessions() Sessions  {
      out, err := exec.Command("sh", "-c", "tmux ls  -F '#{session_name}:#{session_attached}'").Output()
      if err != nil {
        log.Print(err)
      }
      splited_out := strings.Fields(string(out))
      
      var session_slice Sessions
      for index, session := range splited_out {
        splited_sessions := strings.Split(session, ":")
        session := Session{Id: index + 1, Name: splited_sessions[0], Attached: splited_sessions[1]}
        session_slice = append(session_slice, session)
      }
      return session_slice
}

func CreateNewSession() {
      var attach_cmd *exec.Cmd
      attach_cmd = exec.Command("tmux", "attach", "-t", "minne-app")
      attach_cmd.Stdin = os.Stdin
      attach_cmd.Stdout = os.Stdout
      attach_cmd.Stderr = os.Stderr
      attach_cmd.Run()
}
