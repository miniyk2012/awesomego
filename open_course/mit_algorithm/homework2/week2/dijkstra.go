package week2

import (
	"bufio"
	"container/heap"
	"container/list"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

const InitDistance = 1_000_000

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
	from   int
	to     int
	weight int // 权重
}

func (e WeightedDiEdge) String() string {
	return fmt.Sprintf("%d->%d: %d", e.from+1, e.to+1, e.weight)
}

// WeightedDiGraph 虽然作业里是无向图, 但为了通用性, 我还是创建一个有向图, 并且提供个方法判断该图任意2点间必然有对称的边, 以证明测试数据是无向图数据.
// 没有平行边
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
	g.adj[edge.from].PushBack(edge)
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
			it2 := g.Adj(edge.to)
			var hasSymmetry bool
			for it2.HasNext() {
				edge2 := it2.Next()
				if edge2.to == i && edge.weight == edge2.weight {
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
		log.Printf("scan file error: %to", err)
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
			graph.AddEdge(WeightedDiEdge{from: v - 1, to: w - 1, weight: weight})
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
	var shortestDistance = InitDistance
	var selectedEdge WeightedDiEdge
	for v := range n.X {
		it := n.graph.Adj(v)
		for it.HasNext() {
			edge := it.Next()
			if _, ok := n.X[edge.to]; !ok {
				if n.A[v]+edge.weight < shortestDistance {
					shortestDistance = n.A[v] + edge.weight
					selectedEdge = edge
				}
			}
		}
	}
	n.A[selectedEdge.to] = shortestDistance
	n.X[selectedEdge.to] = true
	l := list.New()
	l.PushBackList(n.B[selectedEdge.from])
	l.PushBack(selectedEdge)
	n.B[selectedEdge.to] = l
}

type Elem struct {
	w     int             // 节点下标
	score int             // 最短贪心距离, 若不存在则为无穷大
	edge  *WeightedDiEdge // 贪心边, 如果不存在则为nil
}

// Heap 实现参考: https://www.cnblogs.com/huxianglin/p/6925119.html
type Heap struct {
	s   []Elem
	idx map[int]int // key是节点, value是在heap中的下标
}

func (h Heap) Len() int { return len(h.s) } // 绑定len方法,返回长度
func (h Heap) Less(i, j int) bool { // 绑定less方法
	return h.s[i].score < h.s[j].score
}
func (h Heap) Swap(i, j int) { // 绑定swap方法，交换两个元素位置
	h.s[i], h.s[j] = h.s[j], h.s[i]
	h.idx[h.s[i].w] = i
	h.idx[h.s[j].w] = j
}

func (h *Heap) Pop() any { // 绑定pop方法，从最后拿出一个元素并返回
	old := h.s
	n := len(old)
	x := old[n-1]
	h.s = old[0 : n-1]
	delete(h.idx, x.w)
	return x
}

func (h *Heap) Push(x any) { // 绑定push方法，插入新元素
	h.s = append(h.s, x.(Elem))
	h.idx[x.(Elem).w] = len(h.s) - 1
}

// Index 找到节点v在Heap中的位置
func (h Heap) Index(v int) int {
	return h.idx[v]
}

type HeapBasedDijkstra struct {
	graph  *WeightedDiGraph
	source int          // source(下标0~V-1)
	h      *Heap        // 未处理节点的贪心距离最小堆
	A      []int        // A[i] = 100 表示source到节点i(下标0~V-1)的最短距离是100
	X      map[int]bool // 已算出来的节点(key是0~V-1)
	B      []*list.List //  B[i]是source到i的一条最短路径
}

func InitHeapBasedDijkstra(graph *WeightedDiGraph, s int) *HeapBasedDijkstra {
	x := make(map[int]bool)
	x[s-1] = true
	a := make([]int, graph.V)
	b := make([]*list.List, graph.V)
	for i := range b {
		b[i] = list.New()
	}
	// 初始化Heap
	// 法1
	//hs := make([]Elem, 0, graph.V)
	//h := &Heap{s: hs, idx: make(map[int]int)}
	//for i := 0; i < graph.V; i++ {
	//	if i == s-1 {
	//		heap.Push(h, Elem{w: i, score: 0})
	//	} else {
	//		heap.Push(h, Elem{w: i, score: InitDistance})
	//	}
	//}

	// 法2
	hs := make([]Elem, graph.V)
	h := &Heap{s: hs, idx: make(map[int]int)}
	for i := 0; i < graph.V; i++ {
		if i == s-1 {
			h.s[i] = Elem{w: i, score: 0}
		} else {
			h.s[i] = Elem{w: i, score: InitDistance}
		}
		h.idx[i] = i // 用heap.Init(h)的话, 记得要把idx也做初始化, 否则就不对
	}
	heap.Init(h)
	return &HeapBasedDijkstra{
		graph:  graph,
		source: s - 1,
		h:      h,
		A:      a,
		X:      x,
		B:      b,
	}
}

func (hsp *HeapBasedDijkstra) Run() {
	for hsp.h.Len() > 0 {
		hsp.findNewW()
	}
}

func (hsp *HeapBasedDijkstra) findNewW() {
	eleW := heap.Pop(hsp.h).(Elem)
	hsp.A[eleW.w] = eleW.score
	hsp.X[eleW.w] = true
	l := list.New()
	if eleW.edge != nil {
		l.PushBackList(hsp.B[eleW.edge.from])
		l.PushBack(*eleW.edge)
	}
	hsp.B[eleW.w] = l
	// 更新贪心距离最小堆
	it := hsp.graph.Adj(eleW.w)
	for it.HasNext() {
		edge := it.Next()
		if _, ok := hsp.X[edge.to]; !ok {
			eleV := hsp.h.s[hsp.h.Index(edge.to)]
			if eleV.score > hsp.A[eleW.w]+edge.weight {
				// 当距离变小时才修改堆, 用heap.Fix速度会略微快一些
				eleV.score = hsp.A[eleW.w] + edge.weight
				eleV.edge = &edge
				hsp.h.s[hsp.h.Index(edge.to)] = eleV
				heap.Fix(hsp.h, hsp.h.Index(edge.to))
			}
			//eleV := heap.Remove(hsp.h, hsp.h.Index(edge.to)).(Elem)
			//if eleV.score > hsp.A[eleW.w]+edge.weight {
			//	eleV.score = hsp.A[eleW.w] + edge.weight
			//	eleV.edge = &edge
			//}
			//heap.Push(hsp.h, eleV)
		}
	}
}

func (hsp *HeapBasedDijkstra) Shortest(w int) int {
	return hsp.A[w]
}
