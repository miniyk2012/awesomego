package pkg2

import (
	"testing"

	"github.com/miniyk2012/awesomego/good_example/module/internal/testutil"
	"github.com/stretchr/testify/assert"
)

var _ = testutil.Register(
	TestSub,
)

func TestSub(t *testing.T) {
	t.Logf("TestSub")
	assert.Equal(t, -1, Sub(1, 2))

}
