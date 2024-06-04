package pkg1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Add(a, b int) int {
	return a + b
}

func TAdd(t *testing.T) {
	t.Logf("TestAdd")
	assert.Equal(t, 3, Add(1, 2))
}
