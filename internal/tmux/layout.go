package tmux

import (
	"muxt/internal/config"
	"os"
	"strings"
)

func LayoutToTmux(layout config.Layout) error {
	conf, err := config.GetConfig();
	if err != nil {
		return err;
	}

	home, err := os.UserHomeDir();
	if err != nil {
		return err;
	}
	root := strings.Replace(layout.Root, "~", home, 1)

	err = config.RunExternalCommand("tmux", "new-session", "-d", "-c", root, "-s", layout.Name);
	if err != nil {
		return err;
	}

	for i,w := range layout.Windows {
		if i != 0 {
			err = newWindow(layout.Name, w.Name, root, layout.Attach); 
			if err != nil {
				return err;
			}
		}

		if i == 0 {
			err = renameWindow(layout.Name, w.Name, conf.BaseIndex);
			if err != nil {
				return err;
			}
		}

		err = sendKeys(layout.Name, w.Name, conf.BaseIndex, w.Cmd);
		if err != nil {
			return err;
		}
	}

	return nil;
}
