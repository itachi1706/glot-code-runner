package clojure

import (
	"github.com/itachi1706/glot-code-runner/cmd"
	"path/filepath"
)

func Run(files []string, stdin string) (string, string, error) {
	workDir := filepath.Dir(files[0])
	return cmd.RunStdin(workDir, stdin, "java", "-cp", "/usr/share/java/clojure.jar", "clojure.main", files[0])
}
