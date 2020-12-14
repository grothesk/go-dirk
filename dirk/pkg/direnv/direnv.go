package direnv

import (
	"log"
	"os/exec"
)

var direnvCmd string = "direnv"

// Exists checks if direnv is on the path
func Exists() bool {
	_, err := exec.LookPath(direnvCmd)
	return err == nil
}

// AllowPath grants direnv to load the .envrc in the given path/directory
func AllowPath(p string) {
	cmd := exec.Command("direnv", "allow", p)
	err := cmd.Run()
	if err != nil {
		log.Fatalf("failed executing \"direnv allow\": %s\n", err)
	}
}
