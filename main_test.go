package main

import (
	"bytes"
	"os/exec"
	"strings"
	"testing"
)

func TestMonitor(t *testing.T) {
	tests := []struct {
		name          string
		scriptContent string
		desiredOutput string
	}{
		{
			name:          "PythonScript",
			scriptContent: "scripts/test.py",
			desiredOutput: "test3",
		},
		{
			name:          "BashScript",
			scriptContent: "scripts/test.sh",
			desiredOutput: "test3",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var output bytes.Buffer
			cmd := exec.Command("go", "run", "main.go", tt.scriptContent, tt.desiredOutput)
			cmd.Stdout = &output
			err := cmd.Run()
			if err != nil {
				t.Fatalf("error executing command: %v", err)
			}

			if !strings.Contains(output.String(), tt.desiredOutput) {
				t.Errorf("expected output to contain %q, got %q", tt.desiredOutput, output.String())
			}
		})
	}
}
