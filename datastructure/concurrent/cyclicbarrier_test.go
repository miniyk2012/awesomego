package concurrent

import (
	"testing"
)

func TestCyclicBarrierDemo(t *testing.T) {
	cnt := CyclicBarrierDemo()
	t.Log(cnt)
}
