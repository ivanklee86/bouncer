package bouncer

import (
	"io"
	"os"
)

type Config struct {
	NoExitCode bool
}

// Bouncer is the logic/orchestrator.
type Bouncer struct {
	*Config

	// Allow swapping out stdout/stderr for testing.
	Out io.Writer
	Err io.Writer
}

// New returns a new instance of Bouncer.
func New() *Bouncer {
	config := Config{}

	return &Bouncer{
		Config: &config,
		Out:    os.Stdout,
		Err:    os.Stdin,
	}
}
