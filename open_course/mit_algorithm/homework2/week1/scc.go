package week1

import (
	"bufio"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"

	"golang.org/x/exp/maps"
)

type OutNodes []int
type DirectGraph struct {
	nodes []OutNodes // 保存出边
}

func NewGraph(n int) *DirectGraph {
	nodes := make([]OutNodes, n)
	return &DirectGraph{nodes: nodes}
}

func (g *DirectGraph) EdgeNum() int {
	var total int
	for _, outNodes := range g.nodes {
		total += len(outNodes)
	}
	return total
}
func (g *DirectGraph) Add(a, b int) {
	g.nodes[a-1] = append(g.nodes[a-1], b)
}

func (g *DirectGraph) Reverse() *DirectGraph {
	reversedNodes := make([]OutNodes, len(g.nodes))
	for idx, outNodes := range g.nodes {
		for _, outNode := range outNodes {
			reversedNodes[outNode-1] = append(reversedNodes[outNode-1], idx+1)
		}
	}
	return &DirectGraph{nodes: reversedNodes}
}

func buildGraph(filePath string) *DirectGraph {
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
		node1, _ := strconv.Atoi(splits[0])
		node2, _ := strconv.Atoi(splits[1])
		if node1 > maxNode {
			maxNode = node1
		}
		if node2 > maxNode {
			maxNode = node2
		}
	}
	if err = sc.Err(); err != nil {
		log.Printf("scan file error: %v", err)
		return nil
	}
	graph := NewGraph(maxNode)

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
		node1, _ := strconv.Atoi(splits[0])
		node2, _ := strconv.Atoi(splits[1])
		graph.Add(node1, node2)
	}
	if err = sc.Err(); err != nil {
		log.Printf("scan file error: %v", err)
		return nil
	}
	return graph
}

// FinishTime 是保存结束时间的结构
type FinishTime map[int]int // 保存finish(i) = t

// LastToFirst 从晚到早返回节点
func (finish FinishTime) LastToFirst() []int {
	var timeToNode = make(map[int]int, len(finish))
	var pureTimes = make([]int, 0, len(finish))
	var lastToFirstNodes = make([]int, 0, len(finish))
	for node, pureTime := range finish {
		timeToNode[pureTime] = node
		pureTimes = append(pureTimes, pureTime)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(pureTimes)))
	for _, pureTime := range pureTimes {
		lastToFirstNodes = append(lastToFirstNodes, timeToNode[pureTime])
	}
	return lastToFirstNodes
}

// Leaders 是保存节点所属联通leader的结构
type Leaders map[int]int // 保存leader(i) = s

// kosaraju 强联通Component的算法, 输入一个DirectGraph, 返回强联通Leaders
// 线性时间的算法来找到一个有向图的强连通分量。
func kosaraju(g *DirectGraph) Leaders {
	// 构造g.Reverse()
	grev := g.Reverse()
	// firstPass, 对i逆序遍历Grev, 得到节点对应的结束变量时间t: finish(i) = t
	finish := firstPass(grev)
	// secondPass, 对t逆序遍历G, 得到每个节点的领导s: leader(i) = s
	leader := secondPass(g, finish)
	return leader
}

// firstPass 从n到1做DFS, 获得每个节点的结束时间
func firstPass(g *DirectGraph) FinishTime {
	var t = 0
	var visited = make(map[int]bool, len(g.nodes))
	var finish = make(map[int]int, len(g.nodes))
	var dfs func(i int)
	dfs = func(i int) {
		visited[i] = true
		for _, outNode := range g.nodes[i-1] {
			if !visited[outNode] {
				dfs(outNode)
			}
		}
		t++
		finish[i] = t
	}
	for i := len(g.nodes); i > 0; i-- {
		if !visited[i] {
			dfs(i)
		}
	}
	return finish
}

// secondPass 按finish从晚到早做DFS, 选出每组的leader节点
func secondPass(g *DirectGraph, finish FinishTime) Leaders {
	var s = 0 // 当前的leader节点
	var visited = make(map[int]bool, len(g.nodes))
	var leader = make(map[int]int, len(g.nodes))
	var dfs func(i int)
	dfs = func(i int) {
		visited[i] = true
		leader[i] = s
		for _, outNode := range g.nodes[i-1] {
			if !visited[outNode] {
				dfs(outNode)
			}
		}
	}
	for _, i := range finish.LastToFirst() {
		if !visited[i] {
			s = i
			dfs(i)
		}
	}
	return leader
}

// SCC (Strong Connected Component): 输入graph, 从大到小返回前num组的强联通组的元素个数
func SCC(g *DirectGraph, num int) []int {
	leader := kosaraju(g)
	group := make(map[int][]int) // key是leader节点, value是该scc的组员
	for i, l := range leader {
		group[l] = append(group[l], i)
	}
	leaders := maps.Keys(group)
	sort.Slice(leaders, func(i, j int) bool {
		//fmt.Printf("len(group[%d])>len(group[%d])=%t\n", leaders[i], leaders[j], len(group[leaders[i]]) > len(group[leaders[j]]))
		return len(group[leaders[i]]) > len(group[leaders[j]])
	})
	ret := make([]int, 0, len(group))
	for _, l := range leaders {
		//fmt.Printf("leader is %d, groups is %v\n", l, group[l])
		ret = append(ret, len(group[l]))
	}
	//fmt.Printf("groups is %v", ret)
	for len(ret) < num {
		ret = append(ret, 0)
	}
	return ret[:num]
}
