// Disjoint-Set Implementation with Union-find algorithm
// to demonstrate priority queue.
//
// Correctness not guaranteed. Use at your own risk.
//
// Eddie Blundell eblundell@gmail.com
// (Mostly translated from The Algorithm Design Manual - Steven S, Skiena)

package main

const MAX_SET_SIZE = 1000

type set_union struct {
	p    [MAX_SET_SIZE + 1]int
	size [MAX_SET_SIZE + 1]int
	n    int
}

func (s *set_union) Init(n int) {

	s.n = n
	for i := 0; i < n; i++ {
		s.p[i] = i
		s.size[i] = 1
	}
}

func (s *set_union) Find(x int) int {

	if s.p[x] == x {
		return x
	} else {

		return s.Find(s.p[x])
	}
}

func (s *set_union) UnionSets(s1, s2 int) {

	r1, r2 := s.Find(s1), s.Find(s2)

	if r1 == r2 {
		return
	}

	ns := s.size[r1] + s.size[r2]

	if s.size[r1] >= s.size[r2] {
		s.size[r1] = ns
		s.p[r2] = r1
	} else {
		s.size[r2] = ns
		s.p[r1] = r2
	}
}
