package internal

import (
	"container/heap"
	"math"
)

func Haversine(lat1, lon1, lat2, lon2 float64) float64 {
	const R = 6371e3
	p1 := lat1 * math.Pi / 180
	p2 := lat2 * math.Pi / 180

	delP := (lat2 - lat1) * math.Pi / 180
	delL := (lon2 - lon2) * math.Pi / 180

	a := math.Sin(delP/2)*math.Sin(delP/2) + math.Cos(p1)*math.Cos(p2)*math.Sin(delL/2)*math.Sin(delL/2)

	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return R * c
}

func FindClosestNode(nodes map[int]*Node, lat, lon float64) int {
	minDist := math.Inf(1)
	closestID := -1
	for id, node := range nodes {
		d := Haversine(lat, lon, node.Lat, node.Lon)
		if d < minDist {
			minDist = d
			closestID = id
		}
	}
	return closestID
}

type Item struct {
	NodeID   int
	Priority float64
	Index    int
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

func Dijkstra(graph *Graph, startID, targetID int) ([]int, float64) {
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
