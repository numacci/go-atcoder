package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

// solve implements code to solve the problem.
func solve() {
	a, b := nextInt(), nextInt()
	s := next()

	out("a =", a)
	out("b =", b)
	out("s =", s)
}

// do not edit main as it is just an entrypoint.
func main() {
	defer w.Flush()
	solve()
}

// IO Utils
var (
	sc = bufio.NewScanner(os.Stdin)
	w  = bufio.NewWriter(os.Stdout)
)

const (
	initBufSize = 1024 * 1024
	maxBufSize  = 1024 * 1024 * 1024
)

func init() {
	buf := make([]byte, initBufSize)
	sc.Buffer(buf, maxBufSize)
	sc.Split(bufio.ScanWords)
}

func next() string {
	sc.Scan()
	return sc.Text()
}

func nextInt() int {
	x, _ := strconv.Atoi(next())
	return x
}

func nextFloat() float64 {
	x, _ := strconv.ParseFloat(next(), 64)
	return x
}

func nextStrs(n int) []string {
	ret := make([]string, n)
	for i := 0; i < n; i++ {
		ret[i] = next()
	}
	return ret
}

func nextInts(n int) []int {
	ret := make([]int, n)
	for i := 0; i < n; i++ {
		ret[i] = nextInt()
	}
	return ret
}

func nextFloats(n int) []float64 {
	ret := make([]float64, n)
	for i := 0; i < n; i++ {
		ret[i] = nextFloat()
	}
	return ret
}

func split(s string) []string {
	ret := make([]string, len([]rune(s)))
	for i, v := range []rune(s) {
		ret[i] = string(v)
	}
	return ret
}

func out(x ...interface{}) {
	fmt.Fprintln(w, x)
}

// Math Utils
func abs(a int) int {
	if a < 0 {
		a *= -1
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func nextPermutation(x []int) bool {
	i := len(x) - 2
	if i < 0 {
		return false
	}

	for ; x[i] >= x[i+1]; i-- {
		// completed
		if i == 0 {
			return false
		}
	}

	j := len(x) - 1
	for x[i] >= x[j] {
		j--
	}
	// swap
	x[i], x[j] = x[j], x[i]

	// reverse
	l, r := i+1, len(x)-1
	for l < r {
		x[l], x[r] = x[r], x[l]
		l++
		r--
	}
	return true
}

// Iterator Utils
func reverse(x interface{}) interface{} {
	switch x.(type) {
	case string:
		ret := []rune(x.(string))
		for i := 0; i < len(ret)/2; i++ {
			ret[i], ret[len(ret)-i-1] = ret[len(ret)-i-1], ret[i]
		}
		return string(ret)
	case []string:
		ret := x.([]string)
		for i := 0; i < len(ret)/2; i++ {
			ret[i], ret[len(ret)-i-1] = ret[len(ret)-i-1], ret[i]
		}
		return ret
	case []int:
		ret := x.([]int)
		for i := 0; i < len(ret)/2; i++ {
			ret[i], ret[len(ret)-i-1] = ret[len(ret)-i-1], ret[i]
		}
		return ret
	case []float64:
		ret := x.([]float64)
		for i := 0; i < len(ret)/2; i++ {
			ret[i], ret[len(ret)-i-1] = ret[len(ret)-i-1], ret[i]
		}
		return ret
	default:
		log.Fatalf("Invalid input type for reverse: %T\n", x)
	}
	return nil
}

func lowerBound(x interface{}, key interface{}) int {
	switch x.(type) {
	case []int:
		x, key := x.([]int), key.(int)
		return sort.Search(len(x), func(i int) bool { return x[i] >= key })
	case []float64:
		x, key := x.([]float64), key.(float64)
		return sort.Search(len(x), func(i int) bool { return x[i] >= key })
	default:
		log.Fatalf("Invalid input type for reverse: %T\n", x)
	}
	return -1
}

func upperBound(x interface{}, key interface{}) int {
	switch x.(type) {
	case []int:
		x, key := x.([]int), key.(int)
		return sort.Search(len(x), func(i int) bool { return x[i] > key })
	case []float64:
		x, key := x.([]float64), key.(float64)
		return sort.Search(len(x), func(i int) bool { return x[i] > key })
	default:
		log.Fatalf("Invalid input type for reverse: %T\n", x)
	}
	return -1
}

func count(x interface{}, key interface{}) int {
	ret := 0
	switch x.(type) {
	case string:
		x, key := strings.Split(x.(string), ""), key.(string)
		for _, v := range x {
			if v == key {
				ret++
			}
		}
	case []string:
		x, key := x.([]string), key.(string)
		for _, v := range x {
			if v == key {
				ret++
			}
		}
	case []int:
		x, key := x.([]int), key.(int)
		for _, v := range x {
			if v == key {
				ret++
			}
		}
	case []float64:
		x, key := x.([]float64), key.(float64)
		for _, v := range x {
			if v == key {
				ret++
			}
		}
	default:
		log.Fatalf("Invalid input type for reverse: %T\n", x)
	}
	return ret
}

func find(x interface{}, key interface{}) int {
	ret := -1
	switch x.(type) {
	case string:
		x, key := strings.Split(x.(string), ""), key.(string)
		for i, v := range x {
			if v == key {
				ret = i
				break
			}
		}
		if ret < 0 {
			ret = len(x)
		}
	case []string:
		x, key := x.([]string), key.(string)
		for i, v := range x {
			if v == key {
				ret = i
				break
			}
		}
		if ret < 0 {
			ret = len(x)
		}
	case []int:
		x, key := x.([]int), key.(int)
		for i, v := range x {
			if v == key {
				ret = i
				break
			}
		}
		if ret < 0 {
			ret = len(x)
		}
	case []float64:
		x, key := x.([]float64), key.(float64)
		for i, v := range x {
			if v == key {
				ret = i
				break
			}
		}
		if ret < 0 {
			ret = len(x)
		}
	default:
		log.Fatalf("Invalid input type for reverse: %T\n", x)
	}
	return ret
}

// Data Structure Utils
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
