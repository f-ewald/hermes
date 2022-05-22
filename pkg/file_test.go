package hermes

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestMessageDBFilename(t *testing.T) {
	t.Parallel()

	s, err := MessageDBFilename()
	assert.NoError(t, err)
	assert.True(t, strings.HasSuffix(s, "/Library/Messages/chat.db"))
}
