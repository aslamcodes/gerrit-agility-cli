package main

import (
	"flag"
	"fmt"
	"gcommit/cmd"
	"log"
)

func main() {
	var push = flag.Bool("p", false, "Push changes to develop")
	flag.Usage = func() {
		fmt.Println("Never Again!")
		flag.PrintDefaults()
	}
	flag.Parse()

	if err := cmd.Execute(*push); err != nil {
		log.Fatalf("Error: %v", err)
	}
}
