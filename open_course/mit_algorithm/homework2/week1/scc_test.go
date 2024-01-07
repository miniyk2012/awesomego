package week1

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildGraph(t *testing.T) {
	path := "./testdata/testcase1.txt"
	graph := buildGraph(path)
	t.Logf("%+v", graph)
	grev := graph.Reverse()
	t.Logf("%+v", grev)
	graph = grev.Reverse()
	t.Logf("%+v", graph)

	path = "scc.txt"
	graph = buildGraph(path)
	assert.Equal(t, 875714, len(graph.nodes))
	t.Logf("edge num is %d", graph.EdgeNum())
	grev = graph.Reverse()
	assert.Equal(t, 875714, len(grev.nodes))
	assert.Equal(t, graph.EdgeNum(), grev.EdgeNum())
}

var testcases = []struct {
	testcaseFile string
	group        []int
}{
	{"testcase1.txt", []int{3, 3, 3, 0, 0}},
	{"testcase2.txt", []int{3, 3, 2, 0, 0}},
	{"testcase3.txt", []int{3, 3, 1, 1, 0}},
	{"testcase4.txt", []int{7, 1, 0, 0, 0}},
	{"testcase5.txt", []int{6, 3, 2, 1, 0}},
}

func TestSCC(t *testing.T) {
	for _, testcase := range testcases {
		t.Run(testcase.testcaseFile, func(t *testing.T) {
			graph := buildGraph(filepath.Join("./testdata", testcase.testcaseFile))
			result := SCC(graph, 5)
			t.Logf("expected=%v, result=%v", testcase.group, result)
			assert.Equal(t, testcase.group, result)
		})
	}
}

func TestAssignment(t *testing.T) {
	path := "scc.txt"
	graph := buildGraph(path)
	result := SCC(graph, 5)
	t.Logf("result is %v", result)
	assert.Equal(t, []int{434821, 968, 459, 313, 211}, result)
}
