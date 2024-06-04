package pkg2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Sub(a, b int) int {
	return a - b
}

func TSub(t *testing.T) {
	t.Logf("TestSub")
	assert.Equal(t, -1, Sub(1, 2))

}
