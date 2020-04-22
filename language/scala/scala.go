package scala

import (
	"github.com/itachi1706/glot-code-runner/cmd"
	"path/filepath"
)

func Run(files []string, stdin string) (string, string, error) {
	workDir := filepath.Dir(files[0])

	args := append([]string{"scalac"}, files...)
	stdout, stderr, err := cmd.Run(workDir, args...)
	if err != nil {
		return stdout, stderr, err
	}

	return cmd.RunStdin(workDir, stdin, "scala", "Main")
}
