package gotest_test

import (
	"testing"
	"time"

	"github.com/miniyk2012/awesomego/open_course/cloud_native_go/gotest"
	"github.com/stretchr/testify/assert"
)

func parallelSub1(t *testing.T) {
	t.Parallel()
	t.Logf("test1")
	time.Sleep(3 * time.Second)
}

func parallelSub2(t *testing.T) {
	t.Parallel()
	t.Logf("test2")
	time.Sleep(2 * time.Second)
}

func parallelSub3(t *testing.T) {
	t.Parallel()
	t.Logf("test3")
	time.Sleep(1 * time.Second)
	assert.Equal(t, 3, gotest.Add(1, 2))
}

func TestSubParallel(t *testing.T) {
	t.Logf("set up")
	t.Run("group", func(t *testing.T) {
		t.Run("Test1", parallelSub1)
		t.Run("Test2", parallelSub2)
		t.Run("Test3", parallelSub3)
		t.Logf("inner teardown")
	})
	t.Logf("tear down")
}

func TestSubParallel2(t *testing.T) {
	t.Logf("set up")
	t.Cleanup(func() {
		t.Logf("tear down")
	})
	t.Run("Test1", parallelSub1)
	t.Run("Test2", parallelSub2)
	t.Run("Test3", parallelSub3)
	t.Logf("inner teardown")
}

func TestSubParallel3(t *testing.T) {
	t.Logf("set up")
	t.Run("Test1", parallelSub1)
	t.Run("Test2", parallelSub2)
	t.Run("Test3", parallelSub3)
	t.Logf("tear down")
}
