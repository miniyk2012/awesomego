package main

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func readAssignmentGraph() Graph {
	filePath := "./graph.txt"
	return buildGraph(filePath)
}

func TestBuildGraph(t *testing.T) {
	graph := readAssignmentGraph()
	assert.Equalf(t, 200, len(graph), "vetext size!=200")
	t.Logf("first node edges is %v", graph[1])
	t.Logf("second node edges is %v", graph[2])
	t.Logf("23th node edges is %v", graph[23])
	assert.Equal(t, 20, len(graph[23]))

	graph = buildGraph("./testdata/testcase1.txt")
	t.Logf("testcase1 graph is %v", graph)

	graph = buildGraph("./testdata/testcase5.txt")
	assert.Equalf(t, 40, len(graph), "num is not 80")
}

func TestRandomContraction(t *testing.T) {
	graph := readAssignmentGraph()
	minCutGraph, minCut := randomContraction(graph)
	t.Logf("random mincut is %d, minCutGraph=%v", minCut, minCutGraph)
}

func TestKargerMinCutAssignment(t *testing.T) {
	t.Skipf("耗时比较长, 先不跑")
	graph := readAssignmentGraph()
	minCut := KargerMinCut(graph)
	t.Logf("final mincut is %d", minCut)
	assert.Equal(t, 17, minCut)
}

var testcases = []struct {
	testcaseFile string
	minCut       int
}{
	{"testcase1.txt", 2},
	{"testcase2.txt", 2},
	{"testcase3.txt", 1},
	{"testcase4.txt", 1},
	{"testcase5.txt", 3},
	{"testcase6.txt", 2},
}

func TestKargerMinCut(t *testing.T) {
	for _, testcase := range testcases {
		t.Run(testcase.testcaseFile, func(t *testing.T) {
			graph := buildGraph(filepath.Join("./testdata", testcase.testcaseFile))
			result := KargerMinCut(graph)
			assert.Equal(t, testcase.minCut, result)
		})
	}
}
