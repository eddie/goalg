package main

import "./graph"
import "fmt"

func ConnectedComponents(g *graph.Graph) {

	vertexEarly := func(v int, s *graph.TraversalState) {
		fmt.Printf(" %d", v)
	}

	var state *graph.TraversalState = graph.InitTraversalState()
	var funcs *graph.TraversalFuncs = &graph.TraversalFuncs{Early: vertexEarly}

	c := 0
	for i := 1; i < g.VertexCount(); i++ {

		if state.Discovered(i) {
			continue
		}

		c = c + 1

		fmt.Printf("Component %d\n", c)

		g.BFS(i, funcs, state)
		fmt.Println("")
	}
}

func main() {

	g := graph.CreateGraph(false)
	g.ReadGraph("graph-multi-component.txt")

	ConnectedComponents(g)

}
