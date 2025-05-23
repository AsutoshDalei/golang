package simulation

// I have a map of NYC and there are some 100 car agents with their own destinations to go to and they start off from random points in the map and we get to visualize/see how they interact with each other in the city, in terms of traffic?

import (
	"encoding/json"
	"fmt"
	"io"
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

func LoadNodes(filename string) ([]Node, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("Error in loading node file: %w", err)
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("Error in loading node file: %w", err)
	}

	var nodes []Node
	err = json.Unmarshal(bytes, &nodes)

	return nodes, err
}

func LoadEdges(filename string) ([]Edge, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("Error in loading edge file: %w", err)
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("Error in loading edge file: %w", err)
	}

	var edges []Edge
	err = json.Unmarshal(bytes, &edges)

	return edges, err
}

func BuildGraph(nodes []Node, edges []Edge) *Graph {
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
