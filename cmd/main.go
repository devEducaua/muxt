package main

import (
	"errors"
	"fmt"
	"os"
	"muxt/internal/commands"
)

func main() {
	argv := os.Args[1:];
	if err := handleCommandLineArgs(argv); err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %v\n", err);
		os.Exit(1);
	}
}

func handleCommandLineArgs(argv []string) error {

	if len(argv) < 1 {
		return nil;
	}
	
	command := argv[0];

	switch command {
	case "new":
        if len(argv) < 2 {
            return errors.New("command: `new` expects argument for layout name");
        }
        name := argv[1];
        err := commands.New(name);
        if err != nil {
            return err;
        }
	case "edit":
        if len(argv) < 2 {
            return errors.New("command: `new` expects argument for layout name");
        }
        name := argv[1];
        err := commands.Edit(name);
        if err != nil {
            return err;
        }
	case "start":
        if len(argv) < 2 {
            return errors.New("command: `start` expects argument for layout name");
        }
        name := argv[1];
        err := commands.Start(name);
        if err != nil {
            return err;
        }

	default:
		return fmt.Errorf("unknown command: `%v`", command);
	}
	
	return nil;
}

