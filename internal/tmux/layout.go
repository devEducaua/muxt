package tmux

import (
	"muxt/internal/config"
)

func LayoutToTmux(layout config.Layout) error {
	conf, err := config.GetConfig();
	if err != nil {
		return err;
	}

	root, err := config.ExpandTilde(layout.Root);
	if err != nil {
		return err;
	}

	// TODO: verify if session is running before this, and go to he.
	err = config.RunExternalCommand("tmux", "new-session", "-d", "-c", root, "-s", layout.Name);
	if err != nil {
		return err;
	}

	for wIdx, w := range layout.Windows {
		if wIdx == 0 {
			err = renameWindow(layout.Name, w.Name, conf.BaseIndex);
		} else {
			paneRoot := root;
			if propRoot, ok := w.Panes[0].Props["root"]; ok {
				paneRoot = propRoot.(string);
			}
			paneRoot, err = config.ExpandTilde(paneRoot);
			if err != nil {
				return err;
			}
			err = newWindow(layout.Name, w.Name, paneRoot, layout.Attach); 
		}

		if err != nil {
			return err;
		}
		for pIdx, p := range w.Panes {
			if pIdx != 0 {
				var size int64 = 50;
				split := "v";

				if propSplit, ok := p.Props["split"]; ok {
					split = propSplit.(string);
				}

				if propSize, ok := p.Props["size"]; ok {
					size = propSize.(int64);
				}

				paneRoot, err := config.ExpandTilde(p.Props["root"].(string));
				if err != nil {
					return err;
				}
				if paneRoot == "" {
					paneRoot = root;	
				}

				err = splitWindow(layout.Name, paneRoot, "-"+split, size, w.Name);
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

