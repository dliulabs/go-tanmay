package graph

import (
	"bytes"
	"strconv"
	"text/tabwriter"
)

const (
	INFINITY = uint(99)
	NIL_NODE = ""
)

type Node struct {
	Name  string
	Links []Edge
}

type Edge struct {
	From *Node
	To   *Node
	Cost uint
}
type Graph struct {
	Nodes map[string]*Node
}

func NewGraph() *Graph {
	return &Graph{
		// Nodes: map[string]*Node{},
		Nodes: make(map[string]*Node),
	}
}
func (g *Graph) AddNode(names ...string) {
	for _, name := range names {
		if _, ok := g.Nodes[name]; !ok {
			node := Node{
				Name:  name,
				Links: []Edge{},
			}
			g.Nodes[name] = &node
		}
	}
}

func (g *Graph) AddLink(a, b string, cost uint) {
	nodeA, _ := g.Nodes[a]
	nodeB, _ := g.Nodes[b]
	nodeA.Links = append(nodeA.Links, Edge{
		From: nodeA,
		To:   nodeB,
		Cost: cost,
	})
}

type Distance map[string]uint
type Parent map[string]string
type Visited map[string]bool

func ClosestUnvisitedNode(dist Distance, visited Visited) string {
	mincost := INFINITY
	minnode := NIL_NODE

	for node, cost := range dist {
		if cost >= mincost {
			continue
		}
		if _, ok := visited[node]; !ok {
			mincost = cost
			minnode = node
		}
	}
	return minnode
}

// two dictionary: Costs & Parents
func (g *Graph) Dijkstra(source string) (dist map[string]uint, prev map[string]string) {
	dist, prev = make(Distance, 0), make(Parent, 0)
	for _, node := range g.Nodes {
		dist[node.Name] = INFINITY
		prev[node.Name] = NIL_NODE
	}
	dist[source] = 0
	var visited = make(Visited, 0)

	// build visited & update prev
	for u := source; u != NIL_NODE; u = ClosestUnvisitedNode(dist, visited) {
		uDist := dist[u]
		// for each link from the source find the min costs routes to the dest node
		for _, link := range g.Nodes[u].Links {
			if _, ok := visited[link.To.Name]; ok {
				continue
			}
			v := link.To.Name
			newCost := uDist + link.Cost
			if newCost < dist[v] {
				dist[v] = newCost
				prev[v] = u
			}

		}
		visited[u] = true
	}
	return
}

func DijkstraString(dist Distance, prev Parent) string {
	buf := &bytes.Buffer{}
	writer := tabwriter.NewWriter(buf, 1, 5, 2, ' ', 0)
	writer.Write([]byte("Node\tDistance\tPrev Node\t\n"))
	for key, value := range dist {
		writer.Write([]byte(key + "\t"))
		writer.Write([]byte(strconv.FormatUint(uint64(value), 10) + "\t"))
		writer.Write([]byte(prev[key] + "\t\n"))
	}
	writer.Flush()
	return buf.String()
}
