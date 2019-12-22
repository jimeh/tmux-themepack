package test

import (
	"testing"

	"github.com/jimeh/go-tmux"
	"github.com/stretchr/testify/assert"
)

var powerlineDoubleNames = []string{
	"blue",
	"cyan",
	"green",
	"magenta",
	"orange",
	"purple",
	"red",
	"yellow",
}

func TestPowerlineDoubleThemes(t *testing.T) {
	for _, name := range powerlineDoubleNames {
		filename := "../powerline/double/" + name + ".tmuxtheme"
		c := powerlineColors[name]

		tmuxSetup()

		out, err := tm.Exec("source-file", filename)
		assert.NoErrorf(t, err, `%s: Failed to load theme: %s`, filename, out)

		if err != nil {
			continue
		}

		tmuxHasOptions(t, filename, tmux.GlobalWindow, tmux.Options{
			"clock-mode-colour":            c.color1,
			"clock-mode-style":             "24",
			"mode-style":                   "fg=black,bg=" + c.color1,
			"pane-active-border-style":     "fg=" + c.color1,
			"pane-border-style":            "fg=colour238",
			"window-status-activity-style": "fg=colour245,bg=colour233",
			"window-status-current-format": "#[fg=colour233,bg=black]\ue0b0#[fg=" + c.color2 + ",nobold] #I:#W#F #[fg=colour233,bg=black,nobold]\ue0b2",
			"window-status-current-style":  "fg=" + c.color2 + ",bg=black",
			"window-status-format":         "  #I:#W#F  ",
			"window-status-separator":      "",
		})

		tmuxHasOptions(t, filename, tmux.GlobalSession, tmux.Options{
			"display-panes-active-colour": "colour245",
			"display-panes-colour":        "colour233",
			"message-command-style":       "fg=black,bg=" + c.color1,
			"message-style":               "fg=black,bg=" + c.color1,
			"status-interval":             "1",
			"status-justify":              "centre",
			"status-left":                 "#[fg=colour233,bg=" + c.color1 + ",bold] #S #[fg=" + c.color1 + ",bg=colour240,nobold]\ue0b0#[fg=colour233,bg=colour240] #(whoami) #[fg=colour240,bg=colour235]\ue0b0#[fg=colour240,bg=colour235] #I:#P #[fg=colour235,bg=colour233,nobold]\ue0b0",
			"status-left-length":          "40",
			"status-left-style":           "fg=colour243,bg=colour233",
			"status-right":                "#[fg=colour235,bg=colour233]\ue0b2#[fg=colour240,bg=colour235] %H:%M:%S #[fg=colour240,bg=colour235]\ue0b2#[fg=colour233,bg=colour240] %d-%b-%y #[fg=" + c.color1 + ",bg=colour240]\ue0b2#[fg=colour233,bg=" + c.color1 + ",bold] #H ",
			"status-right-length":         "150",
			"status-right-style":          "fg=colour243,bg=colour233",
			"status-style":                "fg=colour240,bg=colour233",
		})

		tmuxTearDown()
	}
}

func TestPowerlineDoubleThemepackOverrides(t *testing.T) {
	for _, name := range powerlineDoubleNames {
		filename := "../powerline/double/" + name + ".tmuxtheme"

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
		assert.Contains(t, opts["window-status-current-format"],
			"WSCP:WSCF:WSCS")
		assert.Contains(t, opts["window-status-format"], "WSP:WSF:WSS")

		tmuxTearDown()
	}
}

func TestPowerlineDoubleThemeOverrides(t *testing.T) {
	for _, name := range powerlineDoubleNames {
		filename := "../powerline/double/" + name + ".tmuxtheme"

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
}
