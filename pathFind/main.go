package main

import (
	"fmt"
	"geofind/internal"
)

func main() {
	nodes, err := internal.LoadNodes("./data/nodes.json")
	if err != nil {
		panic(err)
	}
	edges, err := internal.LoadEdges("./data/edges.json")
	if err != nil {
		panic(err)
	}

	graph := internal.BuildGraph(nodes, edges)

	fmt.Printf("Loaded %d nodes and %d edges\n", len(graph.Nodes), len(edges))

	startLat, startLon := 40.7992437, -73.9628734
	endLat, endLon := 40.858744, -73.930122

	startID := internal.FindClosestNode(graph.Nodes, startLat, startLon)
	endID := internal.FindClosestNode(graph.Nodes, endLat, endLon)

	fmt.Println("Starting Scan")
	_, cost := internal.Dijkstra(graph, startID, endID)
	fmt.Printf("\nShortest path: %.2f km\n", cost/1000)

}
