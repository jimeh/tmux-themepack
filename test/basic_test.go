package test

import (
	"testing"

	"github.com/jimeh/go-tmux"
	"github.com/stretchr/testify/assert"
)

func TestBasicTheme(t *testing.T) {
	tmuxSetup()
	defer tmuxTearDown()

	theme := "../basic.tmuxtheme"

	_, err := tm.Exec("source-file", theme)
	assert.NoErrorf(t, err, `%s: Failed to load theme`, theme)

	tmuxHasOptions(t, theme, tmux.GlobalWindow, tmux.Options{
		"clock-mode-colour":            "red",
		"clock-mode-style":             "24",
		"mode-style":                   "bg=red",
		"pane-active-border-style":     "fg=green",
		"pane-border-style":            "default",
		"window-status-activity-style": "fg=yellow,bg=black",
		"window-status-current-format": " #I:#W#F ",
		"window-status-current-style":  "fg=black,bg=red",
		"window-status-format":         " #I:#W#F ",
		"window-status-separator":      "",
	})

	tmuxHasOptions(t, theme, tmux.GlobalSession, tmux.Options{
		"display-panes-active-colour": "default",
		"display-panes-colour":        "default",
		"message-command-style":       "default",
		"message-style":               "default",
		"status-interval":             "1",
		"status-justify":              "centre",
		"status-left":                 "#S #[fg=white]» #[fg=yellow]#I #[fg=cyan]#P",
		"status-left-length":          "40",
		"status-left-style":           "fg=green,bg=black",
		"status-right":                "#H #[fg=white]« #[fg=yellow]%H:%M:%S #[fg=green]%d-%b-%y",
		"status-right-length":         "40",
		"status-right-style":          "fg=cyan,bg=black",
		"status-style":                "fg=cyan,bg=black",
	})
}
