package tmux

import (
	"fmt"
	"muxt/internal/config"
)

func newWindow(session, name, root string, attach bool) error {
    command := []string{"tmux", "new-window", "-c", root, "-n", name, "-t", session};
	if attach {
		command = append(command, "-d");
	}
	fmt.Printf("NEW: %v\n", command);
    err := config.RunExternalCommand(command...);
    if err != nil {
        return err;
    }
    return nil;
}

func renameWindow(session, name string, idx any) error {
	command := []string{"tmux", "rename-window", "-t", fmt.Sprintf("%v:%v", session, idx), name};
	fmt.Printf("RENAME: %v\n", command);
    err := config.RunExternalCommand(command...);
    if err != nil {
        return err;
    }
    return nil;
}

func sendKeys(session string, window any, paneIndex int, keys string) error {
    command := []string{"tmux", "send-keys", "-t", fmt.Sprintf("%v:%v.%v", session, window, paneIndex), keys, "C-m"};
	fmt.Printf("SEND: %v\n", command);
    err := config.RunExternalCommand(command...);
    if err != nil {
        return err;
    }
	return nil;
}
