package config

import (
	"errors"
	"os"
	"os/exec"
	"path/filepath"
)

func GetBaseDir() (string, error) {
	xdg := os.Getenv("XDG_CONFIG_HOME");

	if xdg == "" {
		home, err := os.UserHomeDir();
		if err != nil {
			return "", err;
		}
		xdg = filepath.Join(home, ".config");
	}
	path := filepath.Join(xdg, "muxt")
	if err := os.MkdirAll(path, 0755); err != nil {
		return "", err;
	}

	return path, nil;
}

func GetLayoutsDir() (string, error) {
	base, err := GetBaseDir();
	if err != nil {
		return "", err;
	}
	p := filepath.Join(base, "layouts/");
	return p, nil;
}

func OpenEditor(path string) error {
	editor := os.Getenv("EDITOR");
	if editor == "" {
		return errors.New("failed to open $EDITOR, environment variable is not defined.");
	}
	err := RunExternalCommand(editor, path);
	if err != nil {
		return err;
	}
	return nil;
}

func RunExternalCommand(args ...string) error {
	cmd := exec.Command(args[0], args[1:]...);
	cmd.Stdout = os.Stdout;
	cmd.Stderr = os.Stderr;
	cmd.Stdin = os.Stdin;
	if err := cmd.Run(); err != nil {
		return err;
	}
	return nil;
}

