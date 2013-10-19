package main

import "./graph"
import "fmt"

func ProcessEdge(x, y int, state *graph.TraversalState) {

	if state.Discovered(y) && state.Parent(x) != y {

		fmt.Printf("Cycle from %d to %d\n", x, y)
		state.Finished(true)

	}
}

func main() {

	g := graph.CreateGraph(false)
	g.ReadGraph("graph-multi-component.txt")

	var state *graph.TraversalState
	var funcs *graph.TraversalFuncs = &graph.TraversalFuncs{Edge: ProcessEdge}

	g.DFS(1, funcs, state)
}
