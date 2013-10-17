package main

import "./graph"
import "fmt"

func ConnectedComponents(g *graph.Graph) {

	vertexEarly := func(v int) {
		fmt.Printf(" %d", v)
	}

	c := 0
	for i := 1; i < g.VertexCount(); i++ {

		if g.Discovered(i) {
			continue
		}

		c = c + 1

		fmt.Printf("Component %d\n", c)

		g.BFS(i, vertexEarly, nil, nil)

		fmt.Println("")
	}
}

func main() {

	g := graph.CreateGraph(false)
	g.ReadGraph("graph-multi-component.txt")

	ConnectedComponents(g)

}
