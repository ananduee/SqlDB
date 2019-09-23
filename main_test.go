package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
	"testing"
)

func executeShellCommands(input []string, t *testing.T) []string {
	cmd := exec.Command("./SqlDB")
	stdin, err := cmd.StdinPipe()
	if err != nil {
		t.Errorf("Failed to obtain stdin for command - %v.", err)
	}
	go func() {
		defer stdin.Close()
		for _, inputCommand := range input {
			io.WriteString(stdin, inputCommand+"\n")
		}
		io.WriteString(stdin, "exit\n")
	}()
	out, err := cmd.CombinedOutput()

	if err != nil {
		t.Errorf("Failed to get output of command execution - %v.", err)
	}
	return strings.Split(string(out), "\n")
}

func TestInsertHappyCase(t *testing.T) {
	inputCommands := []string{
		"insert 1 anasinha anandms91@gmail.com",
		"insert 2 anasinha anandms91@gmail.com",
		"select",
	}
	output := executeShellCommands(inputCommands, t)
	for _, outputValue := range output {
		fmt.Println(outputValue)
	}
}

// TestMain compiles the binary so that tests can be executed against latest binary
func TestMain(m *testing.M) {
	fmt.Println("Starting test setup.")
	make := exec.Command("make")
	err := make.Run()
	if err != nil {
		fmt.Printf("could not make binary for %v.", err)
		os.Exit(1)
	}
	fmt.Println("Compiled binary.")
	os.Exit(m.Run())
}
