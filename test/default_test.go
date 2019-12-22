package test

import (
	"testing"

	"github.com/jimeh/go-tmux"
	"github.com/stretchr/testify/assert"
)

func TestDefaultTheme(t *testing.T) {
	tmuxSetup()
	defer tmuxTearDown()
	theme := "../default.tmuxtheme"

	_, err := tm.Exec("source-file", theme)
	assert.NoErrorf(t, err, `%s: Failed to load theme`, theme)

	tmuxHasOptions(t, theme, tmux.GlobalWindow, tmux.Options{
		"clock-mode-colour":            "blue",
		"clock-mode-style":             "24",
		"mode-style":                   "fg=black,bg=yellow",
		"pane-active-border-style":     "fg=green",
		"pane-border-style":            "fg=white",
		"window-status-activity-style": "fg=green,bg=black",
		"window-status-current-format": "#I:#W#F",
		"window-status-current-style":  "fg=black,bg=green",
		"window-status-format":         "#I:#W#F",
		"window-status-separator":      " ",
	})

	tmuxHasOptions(t, theme, tmux.GlobalSession, tmux.Options{
		"display-panes-active-colour": "red",
		"display-panes-colour":        "blue",
		"message-command-style":       "fg=black,bg=green",
		"message-style":               "fg=black,bg=yellow",
		"status-interval":             "15",
		"status-justify":              "left",
		"status-left":                 "[#S] ",
		"status-left-length":          "40",
		"status-left-style":           "fg=black,bg=green",
		"status-right":                " \"#H\" %H:%M %d-%b-%y",
		"status-right-length":         "40",
		"status-right-style":          "fg=black,bg=green",
		"status-style":                "fg=black,bg=green",
	})
}

func TestDefaultThemeOverrides(t *testing.T) {
	name := "default"
	filename := "../" + name + ".tmuxtheme"

	tmuxSetup()

	out, err := tm.Exec("source-file", "theme-overrides.conf")
	assert.NoErrorf(t, err, `%s: Failed to load overrides: %s`, name, out)

	out, err = tm.Exec("source-file", filename)
	assert.NoErrorf(t, err, `%s: Failed to load theme: %s`, name, out)

	opts, err := tm.GetOptions(tmux.GlobalSession)
	assert.NoError(t, err)

	assertHasPrefix(t, opts["status-left"], "SLP=")
	assertHasSuffix(t, opts["status-left"], "=SLS")
	assertHasPrefix(t, opts["status-right"], "SRP=")
	assertHasSuffix(t, opts["status-right"], "=SRS")

	opts, err = tm.GetOptions(tmux.GlobalWindow)
	assert.NoError(t, err)

	assertHasPrefix(t, opts["window-status-current-format"], "WSCP=")
	assertHasSuffix(t, opts["window-status-current-format"], "=WSCS")
	assertHasPrefix(t, opts["window-status-format"], "WSP=")
	assertHasSuffix(t, opts["window-status-format"], "=WSS")

	tmuxTearDown()
}
