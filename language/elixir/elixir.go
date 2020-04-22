package elixir

import (
	"github.com/itachi1706/glot-code-runner/cmd"
	"github.com/itachi1706/glot-code-runner/util"
	"path/filepath"
)

func Run(files []string, stdin string) (string, string, error) {
	workDir := filepath.Dir(files[0])
	sourceFiles := util.FilterByExtension(files, "ex")
	args := append([]string{"elixirc"}, sourceFiles...)
	return cmd.RunStdin(workDir, stdin, args...)
}
