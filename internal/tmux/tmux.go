package tmux

import (
	"fmt"
	"muxt/internal/utils"
	"os"
	"os/exec"
	"strings"
)

func GoToSession(session string) error {
	env := os.Getenv("TMUX");
	var err error;
	if env != "" {
		err = utils.TmuxRun("switch-client", "-t", session);
	} else {
		err = utils.TmuxRun("attach-session", "-t", session);
	}
	if err != nil {
		return err;
	}

	return nil;
}

func SessionIsRunning(session string) (bool, error) {
	cmd := exec.Command("tmux", "list-sessions");
	output, err := cmd.CombinedOutput();
	if err != nil {
		return false, err;
	}

	for l := range strings.SplitSeq(string(output), "\n") {
		line := strings.TrimSpace(l);
		if strings.HasPrefix(line, session+":") {
			return true, nil;
		}
	}

	return false, nil;
}

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

func newWindow(session, name, root string) error {
    command := []string{"tmux", "new-window", "-d", "-c", root, "-n", name, "-t", session};
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
