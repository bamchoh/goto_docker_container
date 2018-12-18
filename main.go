package main

import (
	"log"
	"os"
	"os/exec"

	pipeline "github.com/mattn/go-pipeline"
	"github.com/pkg/errors"
)

func main() {
	var err error
	output, err := pipeline.Output(
		[]string{"docker", "ps", "-a", "--format", "{{.Names}}"},
		[]string{"cho"},
	)
	if err != nil {
		log.Fatal(errors.Wrap(err, "[PIPELINE]"))
	}
	container := string(output[0 : len(output)-1])
	err = exec.Command("docker", "start", container).Run()
	if err != nil {
		log.Fatal(errors.Wrap(err, "[DOCKER START]"))
	}
	cmd := exec.Command("docker", "exec", "-it", "-e", "TERM=xterm-256color", container, "bash")
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	err = cmd.Run()
	if err != nil {
		log.Fatal(errors.Wrap(err, "[DOCKER EXEC]"))
	}
}
