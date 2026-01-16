package main

import (
	"fmt"
	"os"

	"apriltag/internal"
)

func main() {
	if len(os.Args) < 2 {
		usage()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "generate":
		if err := internal.RunGenerate(os.Args[2:]); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	case "scan":
		fmt.Fprintln(os.Stderr, "scan is not implemented yet")
		os.Exit(1)
	default:
		usage()
		os.Exit(1)
	}
}

func usage() {
	fmt.Fprintln(os.Stderr, "usage: apriltag <command> [args]")
	fmt.Fprintln(os.Stderr, "")
	fmt.Fprintln(os.Stderr, "commands:")
	fmt.Fprintln(os.Stderr, "  generate   generate an AprilTag PNG")
	fmt.Fprintln(os.Stderr, "  scan       scan an AprilTag (not implemented)")
}
