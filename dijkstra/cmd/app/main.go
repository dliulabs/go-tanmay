package main

import (
	"dijkstra/graph"
	"fmt"
)

func main() {
	g := graph.NewGraph()
	g.AddNode("a", "b", "c", "d", "e")
	g.AddLink("a", "b", 6)
	g.AddLink("d", "a", 1)
	g.AddLink("b", "e", 2)
	g.AddLink("b", "d", 1)
	g.AddLink("c", "e", 5)
	g.AddLink("c", "b", 5)
	g.AddLink("e", "d", 1)
	g.AddLink("e", "c", 4)

	// dist, prev := make(map[string]uint, 0), make(map[string]string, 0)
	// g.AddNode("a", "b", "c", "d", "e")
	// for _, node := range g.Nodes {
	// dist[node.Name] = uint(99)
	// prev[node.Name] = "x"
	// }
	dist, prev := g.Dijkstra("a")
	fmt.Println(graph.DijkstraString(dist, prev))

}
