package cli

import (
	"io"
	"os"
)

type App struct {
	Out io.Writer
	Err io.Writer
}

func DefaultApp() *App {
	return &App{
		Out: os.Stdout,
		Err: os.Stderr,
	}
}
