package main

import (
	"fmt"
)

func main() {
	nodes, err := graph.loadNodes("./data/nodes.json")
	if err != nil {
		panic(err)
	}
	edges, err := graph.loadEdges("./data/edges.json")
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
