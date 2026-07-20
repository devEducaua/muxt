package commands

import (
	"fmt"
	"muxt/internal/config"
	"muxt/internal/tmux"
	"os"
	"path/filepath"
)

func New(name string) error {
	base, err := config.GetBaseDir();    
	if err != nil {
		return err;
	}

	// TODO: serialize this name
	layoutsDirPath := filepath.Join(base, "layouts/");
	if err := os.MkdirAll(layoutsDirPath, 0755); err != nil {
		return err;
	}

	layoutPath := filepath.Join(layoutsDirPath, name+".yml");

	// TODO: maybe turn this is a path, a separated file.
	baseNewLayout := fmt.Sprintf(`
# the name that the session will have
name: %v
# the working directory of your session
root: ~/
`, name)

	f, err := os.OpenFile(layoutPath, os.O_CREATE|os.O_EXCL|os.O_APPEND|os.O_RDWR, 0666);
	if err != nil {
		if os.IsExist(err) {
			return fmt.Errorf("layout: `%v` already exists", name);
		}
		return err;
	}
	defer f.Close();
	f.WriteString(baseNewLayout);

	err = config.OpenEditor(layoutPath);
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

	fileName := name+".yml";
	// TODO: maybe permit the edit opens layouts that doesn't exists.
	dirs, err := os.ReadDir(layoutsDir);
	if err != nil {
		return err;
	}
	exists := false;
	for _,e := range dirs {
		if e.Name() == fileName {
			exists = true;
		}
	}
	if !exists {
		return fmt.Errorf("layout: `%v` doesn't exists", name);
	}

	path := filepath.Join(layoutsDir, fileName);
	err = config.OpenEditor(path);
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
	dirs, err := os.ReadDir(layoutsDir);
	if err != nil {
		return err;
	}
	exists := false;
	for _,e := range dirs {
		if e.Name() == fileName {
			exists = true;
		}
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

	err = tmux.LayoutToTmux(l);
	if err != nil {
		return err;
	}

	return nil;
}

