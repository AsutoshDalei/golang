package main

import (
	"fmt"
	"traffic/simulation"
)

func main() {
	nodes, err := simulation.LoadNodes("data/nodes.json")
	if err != nil {
		panic(err)
	}

	edges, err := simulation.LoadEdges("data/edges.json")
	if err != nil {
		panic(err)
	}

	graph := simulation.BuildGraph(nodes, edges)
	fmt.Println("Graph loaded with", len(graph.Nodes), "nodes and", len(edges), "edges")

	cars := simulation.SimulateMovement(graph, 100)

}
