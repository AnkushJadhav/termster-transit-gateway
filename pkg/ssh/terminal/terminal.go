package terminal

import "golang.org/x/crypto/ssh"

// These constants represent the different known terminals
const (
	XTERM = "xterm"
)

// Config represents the configuration options for a SSH terminal
type Config struct {
	Name   string
	Height int
	Width  int
	Modes  ssh.TerminalModes
}

// GetTerminalConfig sets the config options for the terminal
func GetTerminalConfig() *Config {
	return &Config{
		Name:   XTERM,
		Height: 800,
		Width:  1600,
		Modes: ssh.TerminalModes{
			ssh.ECHO:          1,     // enable echoing
			ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
			ssh.TTY_OP_OSPEED: 14400, // output speed = 14.4kbaud
		},
	}
}
