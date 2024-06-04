package pkg1

import (
	"testing"

	"github.com/miniyk2012/awesomego/good_example/module/internal/testutil"
	"github.com/stretchr/testify/assert"
)

func init() {
	var _ = testutil.Register(
		TestAdd,
	)
}
func TestAdd(t *testing.T) {
	t.Logf("TestAdd")
	assert.Equal(t, 3, Add(1, 2))
}
