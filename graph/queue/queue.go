package queue

// FIFO Queue implementation for graph operations
// Modified https://gist.github.com/moraes/2141121, credit to original author

type Node int

const QUEUE_CAPACITY = 100

// Queue is a basic FIFO queue based on a circular list that resizes as needed.
type Queue struct {
	nodes []Node
	size  int
	head  int
	tail  int
	count int
}

func NewQueue(size int) *Queue {
	return &Queue{
		nodes: make([]Node, size, QUEUE_CAPACITY),
		size:  size,
	}
}

// Push adds a node to the queue.
func (q *Queue) Push(n Node) {
	if q.head == q.tail && q.count > 0 {
		nodes := make([]Node, len(q.nodes)+q.size)
		copy(nodes, q.nodes[q.head:])
		copy(nodes[len(q.nodes)-q.head:], q.nodes[:q.head])
		q.head = 0
		q.tail = len(q.nodes)
		q.nodes = nodes
	}
	q.nodes[q.tail] = n
	q.tail = (q.tail + 1) % len(q.nodes)
	q.count++
}

// Pop removes and returns a node from the queue in first to last order.
func (q *Queue) Pop() Node {
	if q.count == 0 {
		return -1
	}
	node := q.nodes[q.head]
	q.head = (q.head + 1) % len(q.nodes)
	q.count--
	return node
}

func (q *Queue) Empty() bool {
	if q.count == 0 {
		return true
	}
	return false
}
