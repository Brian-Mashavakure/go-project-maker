package commandfunctions

import (
	"fmt"
	"os/exec"
)

func RunModTidyCommand(projectPath string) error {
	cmd := exec.Command("go", "mod", "tidy")
	cmd.Dir = projectPath

	//executing the command
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("Mod Tidy Error : %v", err)
	}

	return nil
}

func RunInitCommand(projectPath string, modulePath string) error {
	//creating the command
	cmd := exec.Command("go", "mod", "init", modulePath)
	cmd.Dir = projectPath

	//executing the command
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("Error running go mod init: %w", err)
	}

	return nil
}
