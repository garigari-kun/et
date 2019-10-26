package tmux_handler

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

type Session struct {
	Id       int
	Name     string
	Attached string
}

type Window struct {
	Id     int
	Name   string
	Active string
}

type LayoutPane struct {
	Id   int
	Name string
}

type Sessions []Session

type Windows []Window

type LayoutPanes []LayoutPane

const (
	InfoColor    = "\033[1;34m%s\033[0m"
	NoticeColor  = "\033[1;36m%s\033[0m"
	WarningColor = "\033[1;33m%s\033[0m"
	ErrorColor   = "\033[1;31m%s\033[0m"
	DebugColor   = "\033[0;36m%s\033[0m"
)

func NewTmuxSessions() Sessions {
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

func NewTmuxWindows() Windows {
	out, err := exec.Command("sh", "-c", "tmux list-windows -F '#{window_name}:#{window_active}'").Output()
	if err != nil {
		log.Print(err)
	}
	splited_out := strings.Fields(string(out))

	var window_slice Windows
	for index, window := range splited_out {
		splited_windows := strings.Split(window, ":")
		window := Window{Id: index + 1, Name: splited_windows[0], Active: splited_windows[1]}
		window_slice = append(window_slice, window)
	}

	return window_slice
}

func NewTmuxLayoutPanes() LayoutPanes {
	var pane_types LayoutPanes
	pane_types = append(pane_types, LayoutPane{Id: 1, Name: "v2"})
	pane_types = append(pane_types, LayoutPane{Id: 2, Name: "v3"})
	pane_types = append(pane_types, LayoutPane{Id: 3, Name: "h2"})
	pane_types = append(pane_types, LayoutPane{Id: 4, Name: "h3"})
	pane_types = append(pane_types, LayoutPane{Id: 5, Name: "tiled"})
	return pane_types
}

func (s Sessions) IsSessionAttached() bool {
	for _, session := range s {
		if session.Attached == "1" {
			return true
		}
	}
	return false
}

func (s Sessions) FindSessionById(id string) string {
	for _, session := range s {
		if strconv.Itoa(session.Id) == id {
			return session.Name
		}
	}
	log.Print("Can't find session name the id you entered.")
	return "false"
}

func (s Sessions) ListTmuxKillingSessions() {
	fmt.Printf(ErrorColor, "=====KILL SESSION=====\n")
	for _, session := range s {
		var list string
		if session.Attached == "1" {
			list = strconv.Itoa(session.Id) + ": " + session.Name + " (Attached)"
		} else {
			list = strconv.Itoa(session.Id) + ": " + session.Name
		}
		fmt.Println(list)
	}
	fmt.Printf(ErrorColor, "=======================\n")
}

func (s Sessions) ListChoicesToTerminal() {
	fmt.Printf(NoticeColor, "=====Create new session or Attach another session=====\n")
	fmt.Printf(WarningColor, "0: Create New Session\n")
	for _, session := range s {
		var list string
		if session.Attached == "1" {
			list = strconv.Itoa(session.Id) + ": " + session.Name + " (Attached)"
		} else {
			list = strconv.Itoa(session.Id) + ": " + session.Name
		}
		fmt.Println(list)
	}
	fmt.Printf(NoticeColor, "======================================================\n")
}

func (w Windows) ListChoicesToTerminalForWindows() {
	fmt.Printf(NoticeColor, "=====Create new window================================\n")
	fmt.Printf(WarningColor, "0: Create New Window\n")
	for _, window := range w {
		var list string
		if window.Active == "1" {
			list = strconv.Itoa(window.Id) + ": " + window.Name + " (Active)"
		} else {
			list = strconv.Itoa(window.Id) + ": " + window.Name
		}
		fmt.Println(list)
	}
	fmt.Printf(NoticeColor, "======================================================\n")
}

func (p LayoutPanes) ListLayoutPanesToTerminal() {
	fmt.Printf(NoticeColor, "=====Select pane layout================================\n")
	for _, pane := range p {
		var list string
		list = strconv.Itoa(pane.Id) + ": " + pane.Name
		fmt.Println(list)
	}
	fmt.Printf(NoticeColor, "=======================================================\n")
}

func (p LayoutPanes) LayoutByChoice(choice string) {
	if choice == "1" {
		LayoutV2()
	} else if choice == "2" {
		LayoutV3()
	} else if choice == "3" {
		LayoutH2()
	} else if choice == "4" {
		LayoutH3()
	} else if choice == "5" {
		LayoutTiled()
	}
}

func LayoutV2() {
	var attach_cmd *exec.Cmd
	attach_cmd = exec.Command("tmux", "split-window", "-h")
	attach_cmd.Stdin = os.Stdin
	attach_cmd.Stdout = os.Stdout
	attach_cmd.Stderr = os.Stderr
	err := attach_cmd.Run()
	if err != nil {
		log.Print(err)
	}
}

func LayoutV3() {
	var attach_cmd *exec.Cmd
	attach_cmd = exec.Command("tmux", "split-window", "-h")
	attach_cmd.Stdin = os.Stdin
	attach_cmd.Stdout = os.Stdout
	attach_cmd.Stderr = os.Stderr
	err := attach_cmd.Run()
	if err != nil {
		log.Print(err)
	}
	attach_cmd = exec.Command("tmux", "split-window", "-v")
	attach_cmd.Stdin = os.Stdin
	attach_cmd.Stdout = os.Stdout
	attach_cmd.Stderr = os.Stderr
	err = attach_cmd.Run()
	if err != nil {
		log.Print(err)
	}
}

func LayoutH2() {
	var attach_cmd *exec.Cmd
	attach_cmd = exec.Command("tmux", "split-window", "-v")
	attach_cmd.Stdin = os.Stdin
	attach_cmd.Stdout = os.Stdout
	attach_cmd.Stderr = os.Stderr
	err := attach_cmd.Run()
	if err != nil {
		log.Print(err)
	}
}

func LayoutH3() {
	var attach_cmd *exec.Cmd
	attach_cmd = exec.Command("tmux", "split-window", "-v")
	attach_cmd.Stdin = os.Stdin
	attach_cmd.Stdout = os.Stdout
	attach_cmd.Stderr = os.Stderr
	err := attach_cmd.Run()
	if err != nil {
		log.Print(err)
	}
	attach_cmd = exec.Command("tmux", "split-window", "-h")
	attach_cmd.Stdin = os.Stdin
	attach_cmd.Stdout = os.Stdout
	attach_cmd.Stderr = os.Stderr
	err = attach_cmd.Run()
	if err != nil {
		log.Print(err)
	}
}

func LayoutTiled() {
	var attach_cmd *exec.Cmd
	attach_cmd = exec.Command("tmux", "split-window", "-v")
	attach_cmd.Stdin = os.Stdin
	attach_cmd.Stdout = os.Stdout
	attach_cmd.Stderr = os.Stderr
	err := attach_cmd.Run()
	if err != nil {
		log.Print(err)
	}
	attach_cmd = exec.Command("tmux", "split-window", "-h")
	attach_cmd.Stdin = os.Stdin
	attach_cmd.Stdout = os.Stdout
	attach_cmd.Stderr = os.Stderr
	err = attach_cmd.Run()
	if err != nil {
		log.Print(err)
	}
	out, err := exec.Command("sh", "-c", "tmux list-panes -F '#{pane_index}'").Output()
	if err != nil {
		log.Print(err)
	}
	splited_out := strings.Fields(string(out))
	last_pane_index := splited_out[len(splited_out)-1]
	index_num, _ := strconv.Atoi(last_pane_index)
	splited_start_at := index_num - 2

	attach_cmd = exec.Command("tmux", "select-pane", "-t", strconv.Itoa(splited_start_at))
	attach_cmd.Stdin = os.Stdin
	attach_cmd.Stdout = os.Stdout
	attach_cmd.Stderr = os.Stderr
	err = attach_cmd.Run()
	if err != nil {
		log.Print(err)
	}
	attach_cmd = exec.Command("tmux", "split-window", "-h")
	attach_cmd.Stdin = os.Stdin
	attach_cmd.Stdout = os.Stdout
	attach_cmd.Stderr = os.Stderr
	err = attach_cmd.Run()
	if err != nil {
		log.Print(err)
	}
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

func CreateNewSession(new_session string, is_switched bool) {
	var attach_cmd *exec.Cmd
	var window_name = new_session + "-main"
	attach_cmd = exec.Command("tmux", "new", "-s", new_session, "-d", "-n", window_name)
	attach_cmd.Stdin = os.Stdin
	attach_cmd.Stdout = os.Stdout
	attach_cmd.Stderr = os.Stderr
	err := attach_cmd.Run()
	if err != nil {
		log.Print(err)
	}

	if is_switched {
		SwitchSession(new_session)
	}
}

func CreateAndAttachSession(new_session string) {
	var attach_cmd *exec.Cmd
	var window_name = new_session + "-main"
	attach_cmd = exec.Command("tmux", "new", "-s", new_session, "-n", window_name)
	attach_cmd.Stdin = os.Stdin
	attach_cmd.Stdout = os.Stdout
	attach_cmd.Stderr = os.Stderr
	err := attach_cmd.Run()
	if err != nil {
		log.Print(err)
	}
}

func CreateNewWindow(new_window string) {
	if len(new_window) == 0 {
		new_window = "subwindow"
	}
	var attach_cmd *exec.Cmd
	attach_cmd = exec.Command("tmux", "new-window", "-n", new_window)
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

func KillSession(session_name string) {
	var attach_cmd *exec.Cmd
	attach_cmd = exec.Command("tmux", "kill-session", "-t", session_name)
	attach_cmd.Stdin = os.Stdin
	attach_cmd.Stdout = os.Stdout
	attach_cmd.Stderr = os.Stderr
	err := attach_cmd.Run()
	if err != nil {
		log.Print(err)
	}
}

func PromptUserChoice() string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf(InfoColor, "Enter what you want: ")
	scanner.Scan()
	text := scanner.Text()
	return text
}

func PromptUserToNewSessionName() string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter new session name: ")
	scanner.Scan()
	new_session_name := scanner.Text()
	return new_session_name
}

func PromptUserToNewWindowName() string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter new window name: ")
	scanner.Scan()
	new_window_name := scanner.Text()
	return new_window_name
}

func PrintBanner() {
	letter :=
		`
███████╗ █████╗ ███████╗██╗   ██╗    ████████╗███╗   ███╗██╗   ██╗██╗  ██╗
██╔════╝██╔══██╗██╔════╝╚██╗ ██╔╝    ╚══██╔══╝████╗ ████║██║   ██║╚██╗██╔╝
█████╗  ███████║███████╗ ╚████╔╝        ██║   ██╔████╔██║██║   ██║ ╚███╔╝ 
██╔══╝  ██╔══██║╚════██║  ╚██╔╝         ██║   ██║╚██╔╝██║██║   ██║ ██╔██╗ 
███████╗██║  ██║███████║   ██║          ██║   ██║ ╚═╝ ██║╚██████╔╝██╔╝ ██╗
╚══════╝╚═╝  ╚═╝╚══════╝   ╚═╝          ╚═╝   ╚═╝     ╚═╝ ╚═════╝ ╚═╝  ╚═╝
																																				 `
	log.Print(letter)
}
