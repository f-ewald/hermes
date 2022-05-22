package templates

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTemplates(t *testing.T) {
	t.Parallel()

	b, err := Templates.ReadFile("statistics.tpl")
	assert.NoError(t, err)
	assert.NotNil(t, b)
}
