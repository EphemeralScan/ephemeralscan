package main

import (
	"fmt"
	"os"
)

const (
	Version = "0.1.0-dev"
)

func main() {
	if len(os.Args) < 2 {
		usage()
		return
	}

	switch os.Args[1] {

	case "version":
		version()

	case "doctor":
		doctor()

	default:
		fmt.Printf("Unknown command: %s\n\n", os.Args[1])
		usage()
		os.Exit(1)
	}
}

func version() {
	fmt.Printf("EphemeralScan %s\n", Version)
}

func doctor() {
	fmt.Println("Doctor checks are not implemented yet.")
}

func usage() {
	fmt.Println("EphemeralScan")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  ephemeralscan <command>")
	fmt.Println()
	fmt.Println("Available commands:")
	fmt.Println("  version")
	fmt.Println("  doctor")
}
