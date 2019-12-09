package test

import (
	"testing"

	"github.com/jimeh/go-tmux"
	"github.com/stretchr/testify/assert"
)

func TestPowerlineDefaultThemes(t *testing.T) {
	tests := []struct {
		filename string
		color1   string
		color2   string
	}{
		{
			filename: "../powerline/default/blue.tmuxtheme",
			color1:   "colour24",
			color2:   "colour33",
		},
		{
			filename: "../powerline/default/cyan.tmuxtheme",
			color1:   "colour39",
			color2:   "colour81",
		},
		{
			filename: "../powerline/default/gray.tmuxtheme",
			color1:   "colour245",
			color2:   "colour250",
		},
		{
			filename: "../powerline/default/green.tmuxtheme",
			color1:   "colour100",
			color2:   "colour190",
		},
		{
			filename: "../powerline/default/magenta.tmuxtheme",
			color1:   "colour125",
			color2:   "colour127",
		},
		{
			filename: "../powerline/default/orange.tmuxtheme",
			color1:   "colour130",
			color2:   "colour166",
		},
		{
			filename: "../powerline/default/purple.tmuxtheme",
			color1:   "colour90",
			color2:   "colour129",
		},
		{
			filename: "../powerline/default/red.tmuxtheme",
			color1:   "colour88",
			color2:   "colour160",
		},
		{
			filename: "../powerline/default/yellow.tmuxtheme",
			color1:   "colour227",
			color2:   "colour227",
		},
	}

	for _, tt := range tests {
		tmuxSetup()

		_, err := tm.Exec("source-file", tt.filename)
		assert.NoErrorf(t, err, `%s: Failed to load theme`, tt.filename)

		tmuxHasOptions(t, tt.filename, tmux.GlobalWindow, tmux.Options{
			"clock-mode-colour":            tt.color1,
			"clock-mode-style":             "24",
			"mode-style":                   "fg=black,bg=" + tt.color1,
			"pane-active-border-style":     "fg=" + tt.color1,
			"pane-border-style":            "fg=colour238",
			"window-status-activity-style": "fg=colour245,bg=colour233",
			"window-status-current-format": "#[fg=colour233,bg=black]\ue0b0#[fg=" + tt.color2 + ",nobold] #I:#W#F #[fg=colour233,bg=black,nobold]\ue0b2",
			"window-status-current-style":  "fg=colour235,bg=colour100",
			"window-status-format":         "  #I:#W#F  ",
			"window-status-separator":      "",
		})

		tmuxHasOptions(t, tt.filename, tmux.GlobalSession, tmux.Options{
			"display-panes-active-colour": "colour245",
			"display-panes-colour":        "colour233",
			"message-command-style":       "fg=black,bg=" + tt.color1,
			"message-style":               "fg=black,bg=" + tt.color1,
			"status-interval":             "1",
			"status-justify":              "centre",
			"status-left":                 "#[fg=colour233,bg=" + tt.color1 + ",bold] #S #[fg=" + tt.color1 + ",bg=colour240,nobold]\ue0b0#[fg=colour233,bg=colour240] #(whoami) #[fg=colour240,bg=colour235]\ue0b0#[fg=colour240,bg=colour235] #I:#P #[fg=colour235,bg=colour233,nobold]\ue0b0",
			"status-left-length":          "40",
			"status-left-style":           "fg=colour243,bg=colour233",
			"status-right":                "#[fg=colour235,bg=colour233]\ue0b2#[fg=colour240,bg=colour235] %H:%M:%S #[fg=colour240,bg=colour235]\ue0b2#[fg=colour233,bg=colour240] %d-%b-%y #[fg=colour245,bg=colour240]\ue0b2#[fg=colour233,bg=colour245,bold] #H ",
			"status-right-length":         "150",
			"status-right-style":          "fg=colour243,bg=colour233",
			"status-style":                "fg=colour240,bg=colour233",
		})

		tmuxTearDown()
	}
}
