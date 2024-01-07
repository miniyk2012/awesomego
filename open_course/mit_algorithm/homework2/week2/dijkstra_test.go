package week2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadGraph(t *testing.T) {
	filePath := "./dijkstraData.txt"
	graph := LoadGraph(filePath)
	assert.True(t, graph.IsSymmetry())
	it := graph.Adj(1) // 2的出边
	for it.HasNext() {
		t.Log(it.Next())
	}

	filePath = "./testdata/testcase1.txt"
	graph = LoadGraph(filePath)
	assert.True(t, graph.IsSymmetry())
}
