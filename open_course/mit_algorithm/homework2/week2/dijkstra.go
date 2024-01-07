package week2

import (
	"bufio"
	"container/list"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

const INIT_DISTANCE = 1_000_000

type EdgeIterator struct {
	l      *list.List
	curEle *list.Element // 下一个要迭代的元素
}

func NewEdgeIterator(l *list.List) *EdgeIterator {
	return &EdgeIterator{l: l, curEle: l.Front()}
}

func (it *EdgeIterator) HasNext() bool {
	return it.curEle != nil
}

func (it *EdgeIterator) Next() WeightedDiEdge {
	oldEle := it.curEle
	it.curEle = oldEle.Next()
	return oldEle.Value.(WeightedDiEdge)
}

func (it *EdgeIterator) Rewind() {
	it.curEle = it.l.Front()
}

// WeightedDiEdge 边
type WeightedDiEdge struct {
	v      int // from
	w      int // to
	weight int // 权重
}

func (e WeightedDiEdge) String() string {
	return fmt.Sprintf("%d->%d: %d", e.v+1, e.w+1, e.weight)
}

// WeightedDiGraph 虽然作业里是无向图, 但为了通用性, 我还是创建一个有向图, 并且提供个方法判断该图任意2点间必然有对称的边, 以证明测试数据是无向图数据.
type WeightedDiGraph struct {
	V   int // 点个数
	E   int // 边条数
	adj []*list.List
}

type Iterator interface {
	HasNext() bool
	Next() WeightedDiEdge
	// Rewind 重新回到迭代器的起点，即第一个数据
	Rewind()
}

func NewWeightedDiGraph(V int) *WeightedDiGraph {
	adj := make([]*list.List, V)
	for i := 0; i < V; i++ {
		adj[i] = list.New()
	}
	return &WeightedDiGraph{V: V, E: 0, adj: adj}
}

func (g *WeightedDiGraph) AddEdge(edge WeightedDiEdge) {
	g.adj[edge.v].PushBack(edge)
	g.E++
}

// Adj v的出边, v是数组下标, 要比实际的点小1
func (g *WeightedDiGraph) Adj(v int) Iterator {
	return NewEdgeIterator(g.adj[v])
}

// Edges 该有向图中的所有边
func (g *WeightedDiGraph) Edges() Iterator {
	l := list.New()
	for i := 0; i < g.V; i++ {
		it := g.Adj(i)
		for it.HasNext() {
			l.PushBack(it.Next())
		}
	}
	return NewEdgeIterator(l)
}

func (g *WeightedDiGraph) IsSymmetry() bool {
	for i := 0; i < g.V; i++ {
		it := g.Adj(i)
		for it.HasNext() {
			edge := it.Next()
			it2 := g.Adj(edge.w)
			var hasSymmetry bool
			for it2.HasNext() {
				edge2 := it2.Next()
				if edge2.w == i && edge.weight == edge2.weight {
					hasSymmetry = true
					break
				}
			}
			if !hasSymmetry {
				return false
			}
		}
	}
	return true
}

func LoadGraph(filePath string) *WeightedDiGraph {
	f, err := os.Open(filePath)
	if err != nil {
		return nil
	}
	defer f.Close()
	var maxNode = 0
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		line := sc.Text() // GET the line string
		splits := strings.Fields(line)
		if len(splits) == 0 {
			continue
		}
		node, _ := strconv.Atoi(splits[0])
		if node > maxNode {
			maxNode = node
		}
	}
	if err = sc.Err(); err != nil {
		log.Printf("scan file error: %v", err)
		return nil
	}

	graph := NewWeightedDiGraph(maxNode)
	// call the Seek method first
	_, err = f.Seek(0, io.SeekStart)
	if err != nil {
		log.Fatal(err)
	}
	sc = bufio.NewScanner(f)
	for sc.Scan() {
		line := sc.Text() // GET the line string
		splits := strings.Fields(line)
		if len(splits) == 0 {
			continue
		}
		v, _ := strconv.Atoi(splits[0])
		for _, t := range splits[1:] {
			tt := strings.Split(t, ",")
			w, _ := strconv.Atoi(tt[0])
			weight, _ := strconv.Atoi(tt[1])
			// 下标要比点小1
			graph.AddEdge(WeightedDiEdge{v: v - 1, w: w - 1, weight: weight})
		}
	}
	return graph
}

type NaiveDijkstra struct {
	graph  *WeightedDiGraph
	source int          // source(下标0~V-1)
	X      map[int]bool // 已算出来的节点(key是0~V-1)
	A      []int        // A[i] = 100 表示source到节点i(下标0~V-1)的最短距离是100
	B      []*list.List //  B[i]是source到i的一条最短路径
}

func InitNaiveDijkstra(graph *WeightedDiGraph, s int) *NaiveDijkstra {
	x := make(map[int]bool)
	x[s-1] = true
	a := make([]int, graph.V)
	a[s-1] = 0
	b := make([]*list.List, graph.V)
	for i := range b {
		b[i] = list.New()
	}
	return &NaiveDijkstra{
		graph:  graph,
		source: s - 1,
		X:      x,
		A:      a,
		B:      b,
	}
}

func (n *NaiveDijkstra) Shortest(w int) int {
	return n.A[w]
}
func (n *NaiveDijkstra) Run() {
	for len(n.X) < n.graph.V {
		n.findNewW()
	}
}

func (n *NaiveDijkstra) findNewW() {
	var shortestDistance = INIT_DISTANCE
	var selectedEdge WeightedDiEdge
	for v := range n.X {
		it := n.graph.Adj(v)
		for it.HasNext() {
			edge := it.Next()
			if _, ok := n.X[edge.w]; !ok {
				if n.A[v]+edge.weight < shortestDistance {
					shortestDistance = n.A[v] + edge.weight
					selectedEdge = edge
				}
			}
		}
	}
	n.A[selectedEdge.w] = shortestDistance
	n.X[selectedEdge.w] = true
	l := list.New()
	l.PushBackList(n.B[selectedEdge.v])
	l.PushBack(selectedEdge)
	n.B[selectedEdge.w] = l
}

type HeapBasedDijkstra struct {
	graph  *WeightedDiGraph
	source int // source(下标0~V-1)
}
