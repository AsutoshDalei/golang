package simulation

import (
	"math/rand"
	"sync"
)

type Vehicle struct {
	ID       int
	Path     []int
	Position int
	Done     bool
}

var wg sync.WaitGroup
var mu sync.Mutex

// Cars := make([]*Vehicle, 10)

func getRandomNodeID(g *Graph) int {
	ids := make([]int, 0, len(g.Nodes))
	for id := range g.Nodes {
		ids = append(ids, id)
	}
	return ids[rand.Intn(len(ids))]
}

func worker(carID int, g *Graph, cars []*Vehicle) {
	defer wg.Done()
	start := getRandomNodeID(g)
	end := getRandomNodeID(g)

	for end == start {
		end = getRandomNodeID(g)
	}

	path, _, ok := AStar(g, start, end)

	if ok && len(path) > 1 {
		car := &Vehicle{ID: carID, Path: path, Position: 0}
		mu.Lock()
		cars[carID] = car
		mu.Unlock()
	}
}

func SimulateMovement(g *Graph, numCars int) []*Vehicle {
	cars := make([]*Vehicle, numCars)
	for i := range numCars {
		wg.Add(1)
		go worker(i, g, cars)
	}
	wg.Wait()
	return cars
}
