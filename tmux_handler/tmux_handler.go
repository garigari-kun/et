package tmux_handler

import (
  "log"
  "os"
  "os/exec"
  "strings"
  "strconv"
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

func IsSessionAttached(sessions Sessions) bool {
  for _, session := range sessions {
    if session.Attached == "1" {
      return true
    }
  }
  return false 
}

func FindSessionById(sessions Sessions, id string) string {
  for _, session := range sessions {
    if strconv.Itoa(session.Id) == id {
      return session.Name
    }
  }
  log.Print("Can't find session name the id you entered.")
  return "false"
}

func AttachSession(session_name string) {
  var attach_cmd *exec.Cmd
  attach_cmd = exec.Command("tmux", "attach", "-t", session_name)
  attach_cmd.Stdin = os.Stdin
  attach_cmd.Stdout = os.Stdout
  attach_cmd.Stderr = os.Stderr
  err := attach_cmd.Run()
  if err != nil {
    log.Print(err)
  }
}

func CreateNewSession(new_session string) {
  var attach_cmd *exec.Cmd
  attach_cmd = exec.Command("tmux", "new", "-s", new_session, "-d")
  attach_cmd.Stdin = os.Stdin
  attach_cmd.Stdout = os.Stdout
  attach_cmd.Stderr = os.Stderr
  err := attach_cmd.Run()
  if err != nil {
    log.Print(err)
  }
}

func SwitchSession(new_session string) {
  var attach_cmd *exec.Cmd
  attach_cmd = exec.Command("tmux switch-client -t " + new_session)
  attach_cmd = exec.Command("tmux", "switch-client", "-t", new_session)
  attach_cmd.Stdin = os.Stdin
  attach_cmd.Stdout = os.Stdout
  attach_cmd.Stderr = os.Stderr
  err := attach_cmd.Run()
  if err != nil {
    log.Print(err)
  }
}

func DetachSession() {
  _, err := exec.Command("sh", "-c", "tmux detach-client").Output()
  if err != nil {
    log.Print("Failed: Detach session.")
  }
}

