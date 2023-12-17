package main

import (
	"bufio"
	"bytes"
	"encoding/gob"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

// Edges key是对端节点, key是与对端节点的边条数
type Edges map[int]int

// Graph key是vertex
// graph[3] = {1:2, 2:1} 代表3-1间有2条边, 3-2间有1条边
type Graph map[int]Edges

func buildGraph(filePath string) Graph {
	f, err := os.Open(filePath)
	if err != nil {
		return nil
	}
	defer f.Close()
	sc := bufio.NewScanner(f)
	graph := make(map[int]Edges)
	for sc.Scan() {
		line := sc.Text() // GET the line string
		splits := strings.Fields(line)
		if len(splits) == 0 {
			continue
		}
		vertex, _ := strconv.Atoi(splits[0])
		edges := make(map[int]int)
		for _, x := range splits[1:] {
			v, _ := strconv.Atoi(x)
			edges[v]++
		}
		graph[vertex] = edges
	}
	if err := sc.Err(); err != nil {
		log.Printf("scan file error: %v", err)
		return nil
	}
	return graph
}

// 随机选2个点
// 这是错误的, 因为很可能这样选的的点是不相连的
func errorSelectTwoNodes(graph Graph) (int, int) {
	// 依赖map key的随机性
	var a, b int
	for x := range graph {
		a = x
		break
	}
	for y := range graph {
		if a == y {
			continue
		}
		b = y
		break
	}
	return a, b
}

// 随机选1条边, 返回2个端点
func selectTwoNodes(graph Graph) (int, int) {
	// 依赖map key的随机性
	var a, b int
	for x, edges := range graph {
		a = x
		for y := range edges {
			b = y
			return a, b
		}
	}
	panic("error")
}

func randomContraction(graph Graph) (Graph, int) {
	for len(graph) > 2 {
		v1, v2 := selectTwoNodes(graph)
		for v, weight := range graph[v2] {
			if v == v1 {
				continue
			}
			graph[v1][v] += weight
			graph[v][v1] += weight
			delete(graph[v], v2)
		}
		delete(graph, v2)
		delete(graph[v1], v2)
	}
	//log.Printf("graph is %v", graph)
	for _, edges := range graph {
		for _, weight := range edges {
			return graph, weight
		}
	}
	panic("cant reach here")
}

func KargerMinCut(graph Graph) int {
	N := int(float64(len(graph)*len(graph)) * math.Log(float64(len(graph))))
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	dec := gob.NewDecoder(&buf)
	minCut := math.MaxInt
	minCutGraph := graph
	log.Printf("total iteration is %d", N)
	for i := 0; i < N; i++ {
		enc.Encode(graph)
		var theGraph Graph
		if err := dec.Decode(&theGraph); err != nil {
			panic("decode graph panic")
		}
		resultGraph, currentMinCut := randomContraction(theGraph)
		if minCut > currentMinCut {
			minCutGraph = resultGraph
			minCut = currentMinCut
		}
		if i%1000 == 0 {
			log.Printf("iteration %d, minCut=%d, minCutGraph=%v", i, minCut, minCutGraph)
		}
	}
	return minCut

}
