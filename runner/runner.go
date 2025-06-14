package runner

import (
	"github.com/ollama/ollama/runner/llamarunner"
	"github.com/ollama/ollama/runner/ollamarunner"
	"os"
)

func Execute(args []string) error {
	if args[0] == "runner" {
		args = args[1:]
	}

	var newRunner bool
	if args[0] == "--ollama-engine" {
		args = args[1:]
		newRunner = true
	}
    //?????????????
	if newRunner {
	    _ = os.Setenv("OLLAMA_ENGINE", "ollamarunner")
		return ollamarunner.Execute(args)
	} else {
	    _ = os.Setenv("OLLAMA_ENGINE", "llamarunner")
		return llamarunner.Execute(args)
	}
}
