package libs

import (
	"fmt"
	"strconv"
	"strings"
)

// pair implementation
type pair struct {
	first  int
	second int
}

func (p *pair) String() string {
	return fmt.Sprintf("{%v, %v}", p.first, p.second)
}

// stack implementation
type stack struct {
	data []int
	size int
}

func newStack(cap int) *stack {
	return &stack{data: make([]int, 0, cap), size: 0}
}

func (s *stack) push(n int) {
	s.data = append(s.data, n)
	s.size++
}

func (s *stack) pop() bool {
	if s.isEmpty() {
		return false
	}
	s.size--
	s.data = s.data[:s.size]
	return true
}

func (s *stack) top() int {
	return s.data[s.size-1]
}

func (s *stack) isEmpty() bool {
	return s.size == 0
}

func (s *stack) String() string {
	return fmt.Sprint(s.data)
}

// queue implementation
type queue struct {
	data []int
	size int
}

func newQueue(cap int) *queue {
	return &queue{data: make([]int, 0, cap), size: 0}
}

func (q *queue) push(n int) {
	q.data = append(q.data, n)
	q.size++
}

func (q *queue) pop() bool {
	if q.isEmpty() {
		return false
	}
	q.size--
	q.data = q.data[1:]
	return true
}

func (q *queue) front() int {
	return q.data[0]
}

func (q *queue) isEmpty() bool {
	return q.size == 0
}

func (q *queue) String() string {
	return fmt.Sprint(q.data)
}

// pairQueue implementation
type pairQueue struct {
	data []*pair
	size int
}

func newPairQueue(cap int) *pairQueue {
	return &pairQueue{data: make([]*pair, 0, cap), size: 0}
}

func (q *pairQueue) push(p *pair) {
	q.data = append(q.data, p)
	q.size++
}

func (q *pairQueue) pop() bool {
	if q.isEmpty() {
		return false
	}
	q.size--
	q.data[0] = nil
	q.data = q.data[1:]
	return true
}

func (q *pairQueue) front() *pair {
	return q.data[0]
}

func (q *pairQueue) isEmpty() bool {
	return q.size == 0
}

func (q *pairQueue) String() string {
	return fmt.Sprint(q.data)
}

// set implementation
type set struct {
	data map[int]struct{}
}

func newSet() *set {
	return &set{data: make(map[int]struct{})}
}

func (s *set) insert(n int) {
	s.data[n] = struct{}{}
}

func (s *set) erase(n int) {
	delete(s.data, n)
}

func (s *set) contains(n int) bool {
	_, ret := s.data[n]
	return ret
}

func (s *set) String() string {
	ret := make([]string, 0, len(s.data))
	for k := range s.data {
		k := strconv.Itoa(k)
		ret = append(ret, k)
	}
	return "{" + strings.Join(ret, ", ") + "}"
}

// intHeap implementation
type intHeap []int

func (h intHeap) Len() int           { return len(h) }
func (h intHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h intHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *intHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *intHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// priorityQueue implementation
type priorityQueue []*item

type item struct {
	p     *pair
	index int
}

func (i *item) String() string {
	return fmt.Sprint(i.p)
}

func (pq priorityQueue) Len() int           { return len(pq) }
func (pq priorityQueue) Less(i, j int) bool { return pq[i].p.first < pq[j].p.first }
func (pq priorityQueue) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }

func (pq *priorityQueue) Push(x interface{}) {
	n := len(*pq)
	p := x.(*pair)
	i := &item{
		p:     p,
		index: n,
	}
	*pq = append(*pq, i)
}

func (pq *priorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	i := old[n-1]
	old[n-1] = nil
	i.index = -1
	*pq = old[0 : n-1]
	return i.p
}
