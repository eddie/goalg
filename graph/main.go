package main

import "./graph"
import "fmt"

func main() {

	g := graph.CreateGraph(false)

	g.ReadGraph("graph.txt")

	fmt.Println("Shortest Path between vertex 6 and 4\n")

	pathFunc := func(x int) {
		fmt.Printf("%d -> ", x)
	}
	g.FindPath(6, 4, pathFunc)

	fmt.Println("")
}
