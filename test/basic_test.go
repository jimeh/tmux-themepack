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

func TestBasicThemepackOverrides(t *testing.T) {
	name := "basic"
	filename := "../" + name + ".tmuxtheme"

	tmuxSetup()

	out, err := tm.Exec("source-file", "themepack-overrides.conf")
	assert.NoErrorf(t, err, `%s: Failed to load overrides: %s`, name, out)

	out, err = tm.Exec("source-file", filename)
	assert.NoErrorf(t, err, `%s: Failed to load theme: %s`, name, out)

	opts, err := tm.GetOptions(tmux.GlobalSession)
	assert.NoError(t, err)
	assert.Contains(t, opts["status-left"], "LLP:LLF:LLS")
	assert.Contains(t, opts["status-left"], "LMP:LMF:LMS")
	assert.Contains(t, opts["status-left"], "LRP:LRF:LRS")
	assert.Contains(t, opts["status-right"], "RLP:RLF:RLS")
	assert.Contains(t, opts["status-right"], "RMP:RMF:RMS")
	assert.Contains(t, opts["status-right"], "RRP:RRF:RRS")

	opts, err = tm.GetOptions(tmux.GlobalWindow)
	assert.NoError(t, err)
	assert.Contains(t, opts["window-status-current-format"], "WSCP:WSCF:WSCS")
	assert.Contains(t, opts["window-status-format"], "WSP:WSF:WSS")

	tmuxTearDown()
}

func TestBasicThemeOverrides(t *testing.T) {
	name := "basic"
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
