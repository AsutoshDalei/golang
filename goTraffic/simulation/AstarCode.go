package simulation

import (
	"container/heap"
	"math"
)

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

func (pq PriorityQ) Swap(i int, j int) {
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

func Haversine(p1 *Node, p2 *Node) float64 {
	const R = 6371e3
	lat1 := p1.Lat * math.Pi / 180
	lat2 := p2.Lat * math.Pi / 180
	dLat := (p2.Lat - p1.Lat) * math.Pi / 180
	dLon := (p2.Lon - p1.Lon) * math.Pi / 180

	sinDlat := math.Sin(dLat / 2)
	sinDlon := math.Sin(dLon / 2)

	aComp := sinDlat*sinDlat + math.Cos(lat1)*math.Cos(lat2)*sinDlon*sinDlon
	c := 2 * math.Atan2(math.Sqrt(aComp), math.Sqrt(1-aComp))

	return R * c
}

func reconstructPath(cameFrom map[int]int, currentID int) []int {
	var path []int
	for {
		path = append([]int{currentID}, path...)
		prev, ok := cameFrom[currentID]
		if !ok {
			break
		} else {
			currentID = prev
		}
	}
	return path
}

func AStar(g *Graph, startID int, targetID int) ([]int, float64, bool) {
	startNode := g.Nodes[startID]
	targetNode := g.Nodes[targetID]

	pq := &PriorityQ{}
	heap.Init(pq)
	heap.Push(pq, &Item{NodeID: startID, Priority: 0})

	cameFrom := make(map[int]int)
	gScore := make(map[int]float64)
	fScore := make(map[int]float64)

	for nodeID := range g.Nodes {
		gScore[nodeID] = math.Inf(1)
		fScore[nodeID] = math.Inf(1)
	}

	gScore[startID] = 0
	fScore[startID] = Haversine(startNode, targetNode)

	for pq.Len() > 0 {
		current := heap.Pop(pq).(*Item)
		currentID := current.NodeID
		if currentID == targetID {
			return reconstructPath(cameFrom, currentID), gScore[targetID], true
		}
		for _, edge := range g.Adjacency[currentID] {
			tentativeG := gScore[currentID] + edge.Length
			if tentativeG < gScore[edge.To] {
				cameFrom[edge.To] = currentID
				gScore[edge.To] = tentativeG
				fScore[edge.To] = tentativeG + Haversine(g.Nodes[edge.To], targetNode)
				heap.Push(pq, &Item{NodeID: edge.To, Priority: fScore[edge.To]})
			}
		}
	}
	return nil, 0, false
}
