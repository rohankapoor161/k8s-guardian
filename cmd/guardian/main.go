package main

import (
	"fmt"
	"os"

	"github.com/rohankapoor/k8s-guardian/pkg/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}