package main

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	cwd, _ := os.Getwd()
	path := filepath.Join(cwd, "two_agent", "store", "docker-compose.yml")

	if _, err := os.Stat(path); err == nil {
		dir := filepath.Dir(path)

		cmd := exec.CommandContext(context.TODO(), "docker", "compose", "up", "-d")
		cmd.Dir = dir

		output, err := cmd.CombinedOutput()

		fmt.Println(string(output))

		if err != nil {
			fmt.Println(err)
			return
		}

		return
	}

}
