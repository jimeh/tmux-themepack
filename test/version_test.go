package test

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTmuxVersion(t *testing.T) {
	tmuxSetup()
	defer tmuxTearDown()

	if tmuxVersion != "" {
		out, err := tm.Exec("-V")
		require.NoError(t, err)

		assert.Equal(
			t,
			"tmux "+tmuxVersion,
			string(bytes.TrimSpace(out)),
		)
	}
}
