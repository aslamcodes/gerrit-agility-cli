package main

import (
	"flag"
	"git-cli/cmd"
	"log"
)

func main() {
	var push = flag.Bool("p", false, "Push changes to develop")
	flag.Parse()

	if err := cmd.Execute(*push); err != nil {
		log.Fatalf("Error: %v", err)
	}
}
