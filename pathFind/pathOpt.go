package main

import (
	"container/heap"
	"encoding/json"
	"fmt"
	"io"
	"math"
	"os"
)

type Node struct {
	ID  int     `json:"id"`
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

type Edge struct {
	From   int     `json:"from"`
	To     int     `json:"to"`
	Length float64 `json:"length"`
}

type Graph struct {
	Nodes     map[int]*Node
	Adjacency map[int][]Edge
}

type Item struct {
	NodeID   int
	Priority float64
	Index    int
}

func loadNodes(filename string) ([]Node, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	var nodes []Node
	err = json.Unmarshal(bytes, &nodes)

	return nodes, err
}

func loadEdges(filename string) ([]Edge, error) {
	file, err := os.Open(filename)

	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	var edges []Edge
	err = json.Unmarshal(bytes, &edges)
	return edges, err
}

func buildGraph(nodes []Node, edges []Edge) *Graph {
	g := &Graph{
		Nodes:     make(map[int]*Node),
		Adjacency: make(map[int][]Edge),
	}

	for i := range nodes {
		node := &nodes[i]
		g.Nodes[node.ID] = node
	}

	for _, edge := range edges {
		g.Adjacency[edge.From] = append(g.Adjacency[edge.From], edge)
	}
	return g
}

func haversine(lat1, lon1, lat2, lon2 float64) float64 {
	const R = 6371e3
	p1 := lat1 * math.Pi / 180
	p2 := lat2 * math.Pi / 180

	delP := (lat2 - lat1) * math.Pi / 180
	delL := (lon2 - lon2) * math.Pi / 180

	a := math.Sin(delP/2)*math.Sin(delP/2) + math.Cos(p1)*math.Cos(p2)*math.Sin(delL/2)*math.Sin(delL/2)

	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return R * c
}

func findClosestNode(nodes map[int]*Node, lat, lon float64) int {
	minDist := math.Inf(1)
	closestID := -1
	for id, node := range nodes {
		d := haversine(lat, lon, node.Lat, node.Lon)
		if d < minDist {
			minDist = d
			closestID = id
		}
	}
	return closestID
}

type PriorityQ []*Item

func (pq PriorityQ) Len() int {
	return len(pq)
}

func (pq PriorityQ) Less(i int, j int) bool {
	return pq[i].Priority < pq[j].Priority
}

func (pq PriorityQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}

func (pq *PriorityQ) Push(x any) {
	n := len(*pq)
	item := x.(*Item)
	item.Index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQ) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func dijkstra(graph *Graph, startID, targetID int) ([]int, float64) {
	dist := make(map[int]float64)
	prev := make(map[int]int)
	for id := range graph.Nodes {
		dist[id] = math.Inf(1)
	}
	dist[startID] = 0

	pq := &PriorityQ{}
	heap.Init(pq)
	heap.Push(pq, &Item{NodeID: startID, Priority: 0})

	for pq.Len() > 0 {
		current := heap.Pop(pq).(*Item)
		u := current.NodeID
		if u == targetID {
			break
		}
		for _, edge := range graph.Adjacency[u] {
			alt := dist[u] + edge.Length
			if alt < dist[edge.To] {
				dist[edge.To] = alt
				prev[edge.To] = u
				heap.Push(pq, &Item{NodeID: edge.To, Priority: alt})
			}
		}
	}

	path := []int{}
	for u := targetID; u != startID; u = prev[u] {
		path = append([]int{u}, path...)
	}
	path = append([]int{startID}, path...)

	return path, dist[targetID]
}

func main() {
	nodes, err := loadNodes("./data/nodes.json")
	if err != nil {
		panic(err)
	}
	edges, err := loadEdges("./data/edges.json")
	if err != nil {
		panic(err)
	}

	graph := buildGraph(nodes, edges)

	fmt.Printf("Loaded %d nodes and %d edges\n", len(graph.Nodes), len(edges))

	startLat, startLon := 40.7992437, -73.9628734
	endLat, endLon := 40.858744, -73.930122

	startID := findClosestNode(graph.Nodes, startLat, startLon)
	endID := findClosestNode(graph.Nodes, endLat, endLon)

	fmt.Println("Starting Scan")
	_, cost := dijkstra(graph, startID, endID)
	fmt.Printf("\nShortest path: %.2f km\n", cost/1000)

}
