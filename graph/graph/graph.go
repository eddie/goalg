// Basic Graph Adjacency List implementation and various operations
// Correctness or efficiency not guaranteed. Use at your own risk.
//
// Eddie Blundell - eblundell@gmail.com
// (Mostly translated from The Algorithm Design Manual - Steven S. Skiena) 

package graph

import (
	"../queue"
	"bufio"
	"fmt"
	"os"
)

const MAXVERT = 100

type Edgenode struct {
	y      int
	weight int
	next   *Edgenode
}

// TODO: merge processed,discovered to enum type
type Graph struct {
	edges      []*Edgenode
	degree     []int
	processed  []bool
	discovered []bool
	parents    []int
	nvertices  int
	nedges     int
	directed   bool
}

func CreateGraph(directed bool) (g *Graph) {

	g = &Graph{
		make([]*Edgenode, 10, MAXVERT),
		make([]int, 10, MAXVERT),
		make([]bool, 10, MAXVERT),
		make([]bool, 10, MAXVERT),
		make([]int, 10, MAXVERT),
		0,
		0,
		directed,
	}

	return g
}

func (g *Graph) InsertEdge(x, y int, directed bool) {

	edgenode := &Edgenode{y, 0, nil}
	edgenode.next = g.edges[x]

	g.edges[x] = edgenode
	g.degree[x]++

	if !directed {
		g.InsertEdge(y, x, true)
	} else {
		g.nedges++
	}
}

func (g *Graph) ReadGraph(fname string) {

	file, err := os.Open(fname)
	if err != nil {
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)

	var x, y int

	for {
		n, err := fmt.Fscanf(reader, "%d %d\n", &x, &y)

		if err != nil {
			break
		}

		if n <= 0 {
			continue
		}

		g.InsertEdge(x, y, g.directed)
		g.nvertices++
	}
}

func (g *Graph) PrintGraph() {
	for i := 1; i <= g.nvertices; i++ {

		var p *Edgenode
		p = g.edges[i]

		if p != nil {
			fmt.Printf("Vertex: %d Adjency Vertices: ", i)
		}
		for p != nil {

			fmt.Printf("%d ", p.y)
			p = p.next
		}

		fmt.Printf("\n")
	}
}

func (g *Graph) BFS(start int, pve func(int), pvl func(int), pe func(int, int)) {

	q := queue.NewQueue(20)
	q.Push(queue.Node(start))

	g.processed[start] = true
	g.parents[start] = -1

	for q.Empty() == false {

		v := int(q.Pop())

		g.discovered[v] = true
		p := g.edges[v]

		if pve != nil {
			pve(v)
		}

		for p != nil {

			y := p.y

			if g.processed[y] == false || g.directed {
				if pe != nil {
					pe(v, y)
				}
			}

			if g.discovered[y] == false {
				q.Push(queue.Node(y))
				g.discovered[y] = true
				g.parents[y] = v
			}

			p = p.next
		}

		if pvl != nil {
			pvl(v)
		}
	}
}

func (g *Graph) FindPathExt(start, end int, cb func(int), parents []int) {

	if (start == end) || (end == -1) {
		cb(start)
	} else {
		g.FindPathExt(start, parents[end], cb, parents)
		cb(end)
	}
}

func (g *Graph) FindPath(start, end int, cb func(int)) {

	g.BFS(start, nil, nil, nil)
	g.FindPathExt(start, end, cb, g.parents)
}

func (g *Graph) Discovered(v int) bool {
	return g.discovered[v]
}

func (g *Graph) VertexCount() int {
	return g.nvertices
}
