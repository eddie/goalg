package main

import "./graph"
import "fmt"

const (
	UNCOLORED = 0
	BLACK     = 1
	WHITE     = 2
)

func BipartitenessTest(g *graph.Graph) bool {

	var color []int = make([]int, 10, graph.MAXVERT)
	bipartite := true

	complement := func(color int) int {
		if color == WHITE {
			return BLACK
		}
		if color == BLACK {
			return WHITE
		}

		return UNCOLORED
	}
	processEdge := func(x, y int) {

		if color[x] == color[y] {
			bipartite = false
		}

		color[y] = complement(color[x])
	}

	for i := 1; i < g.VertexCount(); i++ {

		if g.Discovered(i) {
			continue
		}

		color[i] = WHITE
		g.BFS(i, nil, nil, processEdge)
	}

	return bipartite
}

func main() {

	g := graph.CreateGraph(false)
	g.ReadGraph("graph-bipartite.txt")

	bipartite := BipartitenessTest(g)

	if bipartite {
		fmt.Println("Graph is bipartite")
	} else {
		fmt.Println("Graph is NOT bipartite")
	}
}
