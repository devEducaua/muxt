package tmux

import (
	"fmt"
	"muxt/internal/utils"
)


func splitWindow(session, root, direction string, size int64, window any) error {
	if direction != "-h" && direction != "-v" {
		return fmt.Errorf("invalid direction to split window: `%v`", direction);
	}

	command := []string{"tmux", "split-window", "-c", root, "-t", fmt.Sprintf("%v:%v", session, window), direction, "-p", fmt.Sprintf("%v", size)};
    err := utils.RunExternalCommand(command...);
    if err != nil {
        return err;
    }
    return nil;
}

func newWindow(session, name, root string, attach bool) error {
    command := []string{"tmux", "new-window", "-c", root, "-n", name, "-t", session};
	if attach {
		command = append(command, "-d");
	}
    err := utils.RunExternalCommand(command...);
    if err != nil {
        return err;
    }
    return nil;
}

func renameWindow(session, name string, idx any) error {
	command := []string{"tmux", "rename-window", "-t", fmt.Sprintf("%v:%v", session, idx), name};
    err := utils.RunExternalCommand(command...);
    if err != nil {
        return err;
    }
    return nil;
}

func sendKeys(session string, window any, paneIndex int, keys string) error {
    command := []string{"tmux", "send-keys", "-t", fmt.Sprintf("%v:%v.%v", session, window, paneIndex), keys, "C-m"};
    err := utils.RunExternalCommand(command...);
    if err != nil {
        return err;
    }
	return nil;
}
