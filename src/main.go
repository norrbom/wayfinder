package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	graph := newGraph()
	graph.addEdge("S", "B", 4)
	graph.addEdge("S", "C", 2)
	graph.addEdge("B", "C", 1)
	graph.addEdge("B", "D", 5)
	graph.addEdge("C", "D", 8)
	graph.addEdge("C", "E", 10)
	graph.addEdge("D", "E", 2)
	graph.addEdge("D", "T", 6)
	graph.addEdge("E", "T", 2)

	cost, path := graph.getPath("S", "T")

	graphData, err := json.MarshalIndent(graph.nodes, "", "    ")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	jsonStr := string(graphData)
	fmt.Println(
		"{\"cost\":", cost, ",",
		"\"path\":\"", path, "\",",
		"\"graph\":", jsonStr, "}",
	)
}
