package main

import "./graph"
import "fmt"

const (
	UNCOLORED = 0
	BLACK     = 1
	WHITE     = 2
)

func Complement(color int) int {
	if color == WHITE {
		return BLACK
	}
	if color == BLACK {
		return WHITE
	}

	return UNCOLORED
}

func BipartitenessTest(g *graph.Graph) bool {

	var color []int = make([]int, 10, graph.MAXVERT)
	bipartite := true

	processEdge := func(x, y int, state *graph.TraversalState) {

		if color[x] == color[y] {
			bipartite = false
		}

		color[y] = Complement(color[x])
	}

	var state *graph.TraversalState = graph.InitTraversalState()
	var funcs *graph.TraversalFuncs = &graph.TraversalFuncs{Edge: processEdge}

	for i := 1; i < g.VertexCount(); i++ {

		if state.Discovered(i) {
			continue
		}

		color[i] = WHITE
		g.BFS(i, funcs, state)
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
