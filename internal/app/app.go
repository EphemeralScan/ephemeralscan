package app

import (
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/EphemeralScan/ephemeralscan/configs"
)

const Version = "0.1.0-dev"

const configFileName = "ephemeralscan.yaml"

func Run() {
	if len(os.Args) < 2 {
		usage()
		return
	}

	switch os.Args[1] {

	case "version":
		version()

	case "doctor":
		doctor()

	case "config":
		if err := configCommand(os.Args[2:], os.Stdout); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}

	default:
		fmt.Printf("Unknown command: %s\n\n", os.Args[1])
		usage()
		os.Exit(1)
	}
}

func configCommand(args []string, stdout io.Writer) error {
	if len(args) != 1 || args[0] != "init" {
		return fmt.Errorf("usage: ephemeralscan config init")
	}

	return initConfig(stdout)
}

func initConfig(stdout io.Writer) error {
	file, err := os.OpenFile(configFileName, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0o644)
	if err != nil {
		if errors.Is(err, os.ErrExist) {
			return fmt.Errorf("configuration file %q already exists", configFileName)
		}
		return fmt.Errorf("create configuration file %q: %w", configFileName, err)
	}

	if _, err := file.Write(configs.Template()); err != nil {
		return cleanupIncompleteConfig(file, fmt.Errorf("write configuration file %q: %w", configFileName, err))
	}

	if err := file.Close(); err != nil {
		return errors.Join(
			fmt.Errorf("close configuration file %q: %w", configFileName, err),
			os.Remove(configFileName),
		)
	}

	if _, err := fmt.Fprintf(stdout, "Created %s\n", configFileName); err != nil {
		return fmt.Errorf("write success message: %w", err)
	}

	return nil
}

func cleanupIncompleteConfig(file *os.File, cause error) error {
	return errors.Join(cause, file.Close(), os.Remove(configFileName))
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
	fmt.Println("  config init")
	fmt.Println("  version")
	fmt.Println("  doctor")
}
