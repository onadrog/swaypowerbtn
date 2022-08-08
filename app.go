package main

import (
	"context"
	"encoding/json"
	"log"
	"os/exec"
	"strings"
)

// App struct
type App struct {
	ctx context.Context
}

type ExecCmd struct {
	Cmd, Args string
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) Cmd(cmd string) {
	var data ExecCmd
  err := json.Unmarshal([]byte(cmd), &data)
	if err != nil {
		log.Fatal(err)
	}

	command := exec.Command(data.Cmd, strings.Fields(data.Args)...)
	err = command.Run()
	if err != nil {
		log.Printf("Command finished with error: %v", err)
	}
}
