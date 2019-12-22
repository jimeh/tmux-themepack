package test

import (
	"log"
	"os"
	"testing"

	"github.com/jimeh/go-tmux"
	"github.com/stretchr/testify/assert"
)

var tm = tmux.New()
var tmuxVersion = os.Getenv("TMUX_VERSION")
var tmuxBinPath = os.Getenv("TMUX_BIN_PATH")
var tmuxConfig = os.Getenv("TMUX_CONFIG")

func tmuxSetup() {
	if tmuxBinPath != "" {
		tm.BinPath = tmuxBinPath
	}
	tm.SocketPath = "./tmux-test-socket"

	if tmuxConfig != "" {
		tmuxConfig = "./tmux.conf"
	}

	_, err := tm.Exec("-f", tmuxConfig, "new-session", "-d", "sleep", "30")
	if err != nil {
		log.Fatal(err)
	}
}

func tmuxTearDown() {
	_, err := tm.Exec("kill-server")
	if err != nil {
		log.Fatal(err)
	}

	if tm.SocketPath != "" {
		err = os.Remove(tm.SocketPath)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func tmuxHasOptions(t *testing.T, theme string, s tmux.Scope, m tmux.Options) {
	opts, err := tm.GetOptions(s)
	assert.NoErrorf(t, err, `%s: Failed to get options`, theme)

	if err == nil {
		for k, v := range m {
			_, ok := opts[k]
			assert.Truef(t, ok, `Key "%s" is not available in %s`, k, theme)
			if ok {
				assert.Equalf(t, v, opts[k], `Key "%s" in "%s"`, k, theme)
			}
		}
	}
}
