package graphCode

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
