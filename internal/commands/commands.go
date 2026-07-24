package commands

import (
	"errors"
	"fmt"
	"muxt/internal/config"
	"muxt/internal/tmux"
	"muxt/internal/utils"
	"os"
	"path/filepath"
	"strings"
)

func New(name string) error {
	layoutsDir, err := config.GetLayoutsDir();
	if err != nil {
		return err;
	}
	fileName := name+".kdl";
	exists, err := utils.FileExistsInDir(fileName, layoutsDir);
	if err != nil {
		return err;
	}
	if exists {
		return fmt.Errorf("layout: `%v` already exists, if you want to open it, run `muxt edit %v`", name, name);
	}

	layoutPath := filepath.Join(layoutsDir, name+".kdl");

	// TODO: maybe turn this is a path, a separated file.
	// TODO: add comments in the example base file.
	baseNewLayout := `
layout %v {
	root ~/

	window {
		pane "$EDITOR" 
	}
}
	`

	f, err := os.OpenFile(layoutPath, os.O_CREATE|os.O_EXCL|os.O_APPEND|os.O_RDWR, 0666);
	if err != nil {
		return err;
	}
	defer f.Close();
	f.WriteString(baseNewLayout);

	err = utils.OpenEditor(layoutPath);
	if err != nil {
		return fmt.Errorf("could not open $EDITOR: %v: but base was wrote to `%v`\n", err, layoutPath);
	}

    return nil;
}

func Edit(name string) error {
	layoutsDir, err := config.GetLayoutsDir();
	if err != nil {
		return err;
	}
	fileName := name+".kdl";
	exists, err := utils.FileExistsInDir(fileName, layoutsDir);
	if err != nil {
		return err;
	}
	if !exists {
		return fmt.Errorf("layout: `%v` doesn't exists", name);
	}

	path := filepath.Join(layoutsDir, fileName);
	err = utils.OpenEditor(path);
	if err != nil {
		return err;
	}

	return nil;
}

func Start(name string) error {
	layoutsDir, err := config.GetLayoutsDir();
	if err != nil {
		return err;
	}

	fileName := name+".kdl";
	exists, err := utils.FileExistsInDir(fileName, layoutsDir);
	if err != nil {
		return err;
	}
	if !exists {
		return fmt.Errorf("layout: `%v` doesn't exists", name);
	}

	filePath := filepath.Join(layoutsDir, fileName);
	dat, err := os.ReadFile(filePath);
	if err != nil {
		return err;
	}

	l, err := config.ParseLayout(dat)
	if err != nil {
		return err;
	}

	running, err := tmux.SessionIsRunning(l.Name);
	if err != nil {
		return err;
	}

	if !running {
		err = tmux.LayoutToSession(l);
		if err != nil {
			return err;
		}
	}

	if l.Attach {
		err = tmux.GoToSession(l.Name);
		if err != nil {
			return err;
		}
	}
	return nil;
}

func Copy(from, to string) error {
	dir, err := config.GetLayoutsDir();
	if err != nil {
		return err;
	}

	dirs, err := os.ReadDir(dir);
	if err != nil {
		return err;
	}

	fromExists := false;
	toExists := false;
	for _,e := range dirs {
		switch e.Name() {
		case from:
			fromExists = true;
		case to:
			toExists = true;
		}
	}

	if !fromExists {
		return errors.New("cannot copy a layout that doesn't exists");
	}
	if toExists {
		return errors.New("cannot copy to a layout that already exists");
	}

	return nil;
}

func List() error {
	dir, err := config.GetLayoutsDir();
	if err != nil {
		return err;
	}

	dirs, err := os.ReadDir(dir);
	if err != nil {
		return err;
	}

	for _,e := range dirs {
		name := strings.TrimSuffix(e.Name(), ".kdl");
		fmt.Printf("%v\n", name);
	}
	return nil;
}

