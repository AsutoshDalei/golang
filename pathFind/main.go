package main

import (
	"encoding/json"
	"fmt"
	"geofind/internal"
	"log"
	"net/http"
)

type Coordinates struct {
	Lat1 float64 `json:"lat1"`
	Lon1 float64 `json:"lon1"`
	Lat2 float64 `json:"lat2"`
	Lon2 float64 `json:"lon2"`
}

type Response struct {
	Path []int   `json:"Path"`
	Cost float64 `json:"Cost"`
}

func RouteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var coords Coordinates
	err := json.NewDecoder(r.Body).Decode(&coords)
	if err != nil {
		http.Error(w, "Invalid JSON body", http.StatusBadRequest)
		return
	}

	startID := internal.FindClosestNode(graph.Nodes, coords.Lat1, coords.Lon1)
	endID := internal.FindClosestNode(graph.Nodes, coords.Lat2, coords.Lon2)

	path, cost := internal.Dijkstra(graph, startID, endID)

	resp := Response{Path: path, Cost: cost}
	json.NewEncoder(w).Encode(resp)
}

var graph *internal.Graph

func main() {
	nodes, err := internal.LoadNodes("./data/nodes.json")
	if err != nil {
		panic(err)
	}
	edges, err := internal.LoadEdges("./data/edges.json")
	if err != nil {
		panic(err)
	}

	graph = internal.BuildGraph(nodes, edges)

	fmt.Printf("Graph loaded %d nodes and %d edges\n", len(graph.Nodes), len(edges))

	// startLat, startLon := 40.7992437, -73.9628734
	// endLat, endLon := 40.858744, -73.930122

	// startID := internal.FindClosestNode(graph.Nodes, startLat, startLon)
	// endID := internal.FindClosestNode(graph.Nodes, endLat, endLon)

	// fmt.Println("Starting Scan")
	// _, cost := internal.Dijkstra(graph, startID, endID)
	// fmt.Printf("\nShortest path: %.2f km\n", cost/1000)

	http.HandleFunc("/route", RouteHandler)
	log.Println("Server listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
