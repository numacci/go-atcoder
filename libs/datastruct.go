package libs

import (
	"fmt"
	"strconv"
	"strings"
)

// Pair implementation
type Pair struct {
	first  int
	second int
}

func (p *Pair) String() string {
	return fmt.Sprintf("{%v, %v}", p.first, p.second)
}

// Stack implementation
type Stack struct {
	data []int
	size int
}

func newStack(cap int) *Stack {
	return &Stack{data: make([]int, 0, cap), size: 0}
}

func (s *Stack) push(n int) {
	s.data = append(s.data, n)
	s.size++
}

func (s *Stack) pop() bool {
	if s.isEmpty() {
		return false
	}
	s.size--
	s.data = s.data[:s.size]
	return true
}

func (s *Stack) top() int {
	return s.data[s.size-1]
}

func (s *Stack) isEmpty() bool {
	return s.size == 0
}

func (s *Stack) String() string {
	return fmt.Sprint(s.data)
}

// Queue implementation
type Queue struct {
	data []int
	size int
}

func newQueue(cap int) *Queue {
	return &Queue{data: make([]int, 0, cap), size: 0}
}

func (q *Queue) push(n int) {
	q.data = append(q.data, n)
	q.size++
}

func (q *Queue) pop() bool {
	if q.isEmpty() {
		return false
	}
	q.size--
	q.data = q.data[1:]
	return true
}

func (q *Queue) front() int {
	return q.data[0]
}

func (q *Queue) isEmpty() bool {
	return q.size == 0
}

func (q *Queue) String() string {
	return fmt.Sprint(q.data)
}

// PairQueue implementation
type PairQueue struct {
	data []*Pair
	size int
}

func newPairQueue(cap int) *PairQueue {
	return &PairQueue{data: make([]*Pair, 0, cap), size: 0}
}

func (q *PairQueue) push(p *Pair) {
	q.data = append(q.data, p)
	q.size++
}

func (q *PairQueue) pop() bool {
	if q.isEmpty() {
		return false
	}
	q.size--
	q.data[0] = nil
	q.data = q.data[1:]
	return true
}

func (q *PairQueue) front() *Pair {
	return q.data[0]
}

func (q *PairQueue) isEmpty() bool {
	return q.size == 0
}

func (q *PairQueue) String() string {
	return fmt.Sprint(q.data)
}

// Set implementation
type Set struct {
	data map[int]struct{}
}

func newSet() *Set {
	return &Set{data: make(map[int]struct{})}
}

func (s *Set) insert(n int) {
	s.data[n] = struct{}{}
}

func (s *Set) erase(n int) {
	delete(s.data, n)
}

func (s *Set) contains(n int) bool {
	_, ret := s.data[n]
	return ret
}

func (s *Set) String() string {
	ret := make([]string, 0, len(s.data))
	for k := range s.data {
		k := strconv.Itoa(k)
		ret = append(ret, k)
	}
	return "{" + strings.Join(ret, ", ") + "}"
}

// IntHeap implementation
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// PriorityQueue implementation
type PriorityQueue []*item

type item struct {
	p     *Pair
	index int
}

func (i *item) String() string {
	return fmt.Sprint(i.p)
}

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].p.first < pq[j].p.first }
func (pq PriorityQueue) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	p := x.(*Pair)
	i := &item{
		p:     p,
		index: n,
	}
	*pq = append(*pq, i)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	i := old[n-1]
	old[n-1] = nil
	i.index = -1
	*pq = old[0 : n-1]
	return i.p
}

// UnionFind implementation
type UnionFind struct {
	par  []int
	rank []int
}

// NewUnionFind initialize UnionFind
func NewUnionFind(n int) *UnionFind {
	uf := &UnionFind{
		par:  make([]int, n),
		rank: make([]int, n),
	}
	for i := 0; i < n; i++ {
		uf.par[i] = -1
	}
	return uf
}

func (uf *UnionFind) Find(x int) int {
	if uf.par[x] < 0 {
		return x
	} else {
		uf.par[x] = uf.Find(uf.par[x])
		return uf.par[x]
	}
}

func (uf *UnionFind) Unite(x, y int) {
	x = uf.Find(x)
	y = uf.Find(y)
	if x == y {
		return
	}

	if uf.rank[x] < uf.rank[y] {
		x, y = y, x
	}
	if uf.rank[x] == uf.rank[y] {
		uf.rank[x]++
	}

	uf.par[x] += uf.par[y]
	uf.par[y] = x
}

func (uf *UnionFind) Same(x, y int) bool {
	return uf.Find(x) == uf.Find(y)
}

func (uf *UnionFind) Size(x int) int {
	return -uf.par[uf.Find(x)]
}

func (uf *UnionFind) String() string {
	return fmt.Sprintf("par = %v, rank = %v", uf.par, uf.rank)
}
