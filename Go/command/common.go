package command

import (
	"os/exec"
)

type Screen struct {
	DisplayName string
}

// TurnOff Turn off the screen
func (s *Screen) TurnOff() {
	exec.Command("xset", "-display", s.DisplayName, "dpms", "force", "off").Run()
}

// WakeUp Wake up the screen
func (s *Screen) WakeUp() {
	exec.Command("xset", "-display", s.DisplayName, "dpms", "force", "on").Run()
}
