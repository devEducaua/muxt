package utils

import (
	"errors"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

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

func ExpandTilde(p string) (string, error) {
	if !strings.HasPrefix(p, "~") {
		return p, nil;
	}

	home, err := os.UserHomeDir();
	if err != nil {
		return "", err;
	}

	trimmed := strings.TrimPrefix(p, "~");
	path := filepath.Join(home, trimmed);

	return path, nil;
}

func TmuxRun(args ...string) error {
	cmd := exec.Command("tmux", args...);
	cmd.Stdout = os.Stdout;
	cmd.Stderr = os.Stderr;
	cmd.Stdin = os.Stdin;
	if err := cmd.Run(); err != nil {
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

func FileExistsInDir(filename, dirPath string) (bool, error) {
	dirs, err := os.ReadDir(dirPath);
	if err != nil {
		return false,err;
	}

	exists := false;
	for _,e := range dirs {
		if e.Name() == filename {
			exists = true;
		}
	}

	return exists, nil;
}

