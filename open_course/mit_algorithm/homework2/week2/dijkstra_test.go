package week2

import (
	"fmt"
	"path/filepath"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadGraph(t *testing.T) {
	filePath := "./dijkstraData.txt"
	graph := LoadGraph(filePath)
	assert.NotNil(t, graph)
	assert.True(t, graph.IsSymmetry())
	it := graph.Adj(1) // 2的出边
	for it.HasNext() {
		t.Log(it.Next())
	}
	t.Logf("dijkstraData Graph V=%d, E=%d", graph.V, graph.V)

	filePath = "./testdata/testcase1.txt"
	graph = LoadGraph(filePath)
	assert.NotNil(t, graph)
	assert.True(t, graph.IsSymmetry())
	allEdges := graph.Edges()
	var totalEdges int
	for allEdges.HasNext() {
		totalEdges++
		t.Log(allEdges.Next())
	}
	assert.Equal(t, 16, totalEdges)
	assert.Equal(t, 16, graph.E)
	assert.Equal(t, 8, graph.V)
}

var testcases = []struct {
	testcaseFile string
	dist         []int
	paths        [][]int
}{
	{"testcase1.txt", []int{0, 1, 2, 3, 4, 4, 3, 2}, [][]int{
		nil, {2}, {2, 3}, {2, 3, 4}, {2, 3, 4, 5}, {8, 7, 6}, {8, 7}, {8},
	}},
}

func TestNaiveSP(t *testing.T) {
	for _, testcase := range testcases {
		t.Run(testcase.testcaseFile, func(t *testing.T) {
			graph := LoadGraph(filepath.Join("./testdata", testcase.testcaseFile))
			sp := InitNaiveDijkstra(graph, 1)
			sp.Run()
			for i := 0; i < graph.V; i++ {
				var sb strings.Builder
				var path []int
				it := NewEdgeIterator(sp.B[i])
				for it.HasNext() {
					edge := it.Next()
					sb.WriteString(fmt.Sprintf("%s, ", edge))
					path = append(path, edge.to+1)
				}
				line := sb.String()
				sb.Reset()
				if len(line) >= 2 {
					sb.WriteString(line[:len(line)-2])
					sb.WriteString(": ")
				}
				sb.WriteString(fmt.Sprintf("1 to %d distance=%d", i+1, sp.Shortest(i)))
				t.Log(sb.String())
				assert.Equal(t, testcase.dist[i], sp.Shortest(i))
				assert.Equal(t, testcase.paths[i], path)
			}
		})
	}
}

func TestHeapSP(t *testing.T) {
	for _, testcase := range testcases {
		t.Run(testcase.testcaseFile, func(t *testing.T) {
			graph := LoadGraph(filepath.Join("./testdata", testcase.testcaseFile))
			sp := InitHeapBasedDijkstra(graph, 1)
			sp.Run()
			for i := 0; i < graph.V; i++ {
				var sb strings.Builder
				var path []int
				it := NewEdgeIterator(sp.B[i])
				for it.HasNext() {
					edge := it.Next()
					sb.WriteString(fmt.Sprintf("%s, ", edge))
					path = append(path, edge.to+1)
				}
				line := sb.String()
				sb.Reset()
				if len(line) >= 2 {
					sb.WriteString(line[:len(line)-2])
					sb.WriteString(": ")
				}
				sb.WriteString(fmt.Sprintf("1 to %d distance=%d", i+1, sp.Shortest(i)))
				t.Log(sb.String())
				assert.Equal(t, testcase.dist[i], sp.Shortest(i))
				assert.Equal(t, testcase.paths[i], path)
			}
		})
	}
}

var testCases = []struct {
	testcaseFile string
	dist         string
}{
	{"input_random_15_32.txt", "5194,9990,8494,8548,14509,14421,8601,10812,9890,6589"},
	{"input_random_28_256.txt", "561210,512598,559247,660768,485338,534807,364902,307456,511454,453935"},
}

func TestHeapSPUnitCase(t *testing.T) {
	var nodes = []int{7, 37, 59, 82, 99, 115, 133, 165, 188, 197}
	for _, testcase := range testCases {
		t.Run(testcase.testcaseFile, func(t *testing.T) {
			graph := LoadGraph(filepath.Join("./testdata", testcase.testcaseFile))
			sp := InitHeapBasedDijkstra(graph, 1)
			sp.Run()
			var sb strings.Builder
			for i, id := range nodes {
				v := sp.Shortest(id - 1)
				sb.WriteString(strconv.Itoa(v))
				if i < len(nodes)-1 {
					sb.WriteString(",")
				}
			}
			assert.Equal(t, testcase.dist, sb.String())
		})
	}
}

func TestAssignment(t *testing.T) {
	var nodes = []int{7, 37, 59, 82, 99, 115, 133, 165, 188, 197}
	graph := LoadGraph("./dijkstraData.txt")
	assert.NotNil(t, graph)
	sp := InitNaiveDijkstra(graph, 1)
	sp.Run()
	var sb strings.Builder
	for i, id := range nodes {
		v := sp.Shortest(id - 1)
		sb.WriteString(strconv.Itoa(v))
		if i < len(nodes)-1 {
			sb.WriteString(",")
		}
	}
	assert.Equal(t, "2599,2610,2947,2052,2367,2399,2029,2442,2505,3068", sb.String())

	hsp := InitHeapBasedDijkstra(graph, 1)
	sb.Reset()
	hsp.Run()
	for i, id := range nodes {
		v := hsp.Shortest(id - 1)
		sb.WriteString(strconv.Itoa(v))
		if i < len(nodes)-1 {
			sb.WriteString(",")
		}
	}
	assert.Equal(t, "2599,2610,2947,2052,2367,2399,2029,2442,2505,3068", sb.String())
}

func BenchmarkNaiveDijkstra_Run(b *testing.B) {
	graph := LoadGraph("./dijkstraData.txt")
	for i := 0; i < b.N; i++ {
		sp := InitNaiveDijkstra(graph, 1)
		sp.Run()
	}
}

func BenchmarkHeapBasedDijkstra_Run(b *testing.B) {
	graph := LoadGraph("./dijkstraData.txt")
	for i := 0; i < b.N; i++ {
		sp := InitHeapBasedDijkstra(graph, 1)
		sp.Run()
	}
}
