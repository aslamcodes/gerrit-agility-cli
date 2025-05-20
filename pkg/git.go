package git

import (
	"fmt"
	"os/exec"
)

// AddChanges stages all changes
func AddChanges() error {
	cmd := exec.Command("git", "add", ".")
	if err := cmd.Run(); err != nil {
		output, _ := cmd.CombinedOutput()
		return fmt.Errorf("failed to add changes: %s", output)
	}
	return nil
}

// CommitChanges commits the changes with the provided ticket and message
func CommitChanges(ticket, message string) error {
	cmd := exec.Command("git", "commit", "-m", fmt.Sprintf("%s %s", ticket, message))
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to commit changes: %v", err)
	}
	return nil
}

// PushChanges pushes the changes to the develop branch
func PushChanges(ticket string) error {
	cmd := exec.Command("git", "push", "origin", fmt.Sprintf("HEAD:refs/for/develop%%topic=%s", ticket))
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to push changes: %v", err)
	}
	return nil
}
