package pkg1

import (
	"testing"

	"github.com/miniyk2012/awesomego/good_example/module/app2"
	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	app2.Setup()
	t.Logf("TestAdd")
	assert.Equal(t, 3, Add(1, 2))
	app2.Teardown()
}
