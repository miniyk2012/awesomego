package tests

import (
	"testing"

	// Use the blank identifier for "side-effect-only" imports
	"github.com/miniyk2012/awesomego/good_example/module/app/pkg1"
	"github.com/miniyk2012/awesomego/good_example/module/app/pkg2"
	"github.com/miniyk2012/awesomego/good_example/module/internal/testutil"
)

func Test1(t *testing.T) {
	testutil.TestAll(t)
}

func Test2(t *testing.T) {
	t.Run("pk1", pkg1.TAdd)
	t.Run("pk2", pkg2.TSub)
}
