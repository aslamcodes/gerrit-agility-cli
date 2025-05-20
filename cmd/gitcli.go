package cmd

import (
	"fmt"
	"git-cli/pkg"
	"strings"
)

func Execute(push bool) error {
	var ticket string
	var message string

	if err := git.AddChanges(); err != nil {
		return err
	}

	fmt.Println("Enter ticket number (Eg. B-123123)")
	fmt.Scan(&ticket)

	// Validations
	if len(ticket) != 8 {
		return fmt.Errorf("invalid ticket, should be equal to 8 characters")
	}

	if !strings.HasPrefix(ticket, "B-") {
		return fmt.Errorf("invalid ticket, should be starting with B-")
	}

	fmt.Println("Enter commit message")
	fmt.Scan(&message)

	if err := git.CommitChanges(ticket, message); err != nil {
		return err
	}

	if !push {
		fmt.Print("Changes committed")
		return nil
	}

	fmt.Println("Pushing the changes")
	if err := git.PushChanges(ticket); err != nil {
		return err
	}

	fmt.Println("Changes committed and pushed")
	return nil
}
