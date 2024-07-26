package pkg2

import (
	"testing"

	"github.com/miniyk2012/awesomego/good_example/module/app2"
	"github.com/stretchr/testify/assert"
)

func TestSub(t *testing.T) {
	app2.Setup()
	t.Logf("TestSub")
	assert.Equal(t, -1, Sub(1, 2))
	app2.Teardown()
	//data := []byte("{\"message\":\"insertDeleteTceServiceTask error\",\"task_status\":1}")
}
