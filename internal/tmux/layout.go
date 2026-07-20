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
	// TODO: verify if session is running before this, and go to he.
	err = config.RunExternalCommand("tmux", "new-session", "-d", "-c", root, "-s", layout.Name);
	if err != nil {
		return err;
	}

	for wIdx, w := range layout.Windows {
		if wIdx == 0 {
			err = renameWindow(layout.Name, w.Name, conf.BaseIndex);
		} else {
			err = newWindow(layout.Name, w.Name, root, layout.Attach); 
		}

		if err != nil {
			return err;
		}
		for pIdx, p := range w.Panes {
			if pIdx != 0 {
				var size int64 = 50;
				split := "v";

				propSplit := p.Props["split"];
				propSize := p.Props["size"];

				if propSize != nil {
					size = propSize.(int64);
				}
				if propSplit != nil {
					split = propSplit.(string);
				}

				err = splitWindow(layout.Name, root, "-"+split, size, w.Name);
				if err != nil {
					return err;
				}
			}	

			err = sendKeys(layout.Name, w.Name, pIdx+conf.BaseIndex, p.Cmd);
			if err != nil {
				return err;
			}
		}
	}

	return nil;
}
