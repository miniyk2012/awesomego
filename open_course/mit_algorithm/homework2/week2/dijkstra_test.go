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
	t.Log()
	filePath = "./testdata/testcase1.txt"
	graph = LoadGraph(filePath)
	assert.True(t, graph.IsSymmetry())
	allEdges := graph.Edges()
	var totalEdges int
	for allEdges.HasNext() {
		totalEdges++
		t.Log(allEdges.Next())
	}
	assert.Equal(t, 16, totalEdges)
}
