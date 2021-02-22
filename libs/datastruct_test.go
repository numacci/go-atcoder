package libs

import (
	"container/heap"
	"fmt"
	"testing"
)

func TestStack(t *testing.T) {
	stack := newStack(100)

	if !stack.isEmpty() {
		t.Fatalf("True is expected, but %v\n", stack.isEmpty())
	}

	stack.push(10)
	stack.push(1)
	stack.push(-5)

	if stack.isEmpty() {
		t.Fatalf("False is expected, but %v\n", stack.isEmpty())
	}

	if stack.top() != -5 {
		t.Fatalf("-5 is expected, but %v\n", stack.top())
	}

	stack.pop()
	if stack.top() != 1 {
		t.Fatalf("1 is expected, but %v\n", stack.top())
	}

	stack.pop()
	if stack.top() != 10 {
		t.Fatalf("10 is expected, but %v\n", stack.top())
	}

	stack.pop()
	if !stack.isEmpty() {
		t.Fatalf("True is expected, but %v\n", stack.isEmpty())
	}
}

func TestQueue(t *testing.T) {
	queue := newQueue(10)

	if !queue.isEmpty() {
		t.Fatalf("True is expected, but %v\n", queue.isEmpty())
	}

	queue.push(10)
	queue.push(1)
	queue.push(-5)

	if queue.isEmpty() {
		t.Fatalf("False is expected, but %v\n", queue.isEmpty())
	}

	if queue.front() != 10 {
		t.Fatalf("10 is expected, but %v\n", queue.front())
	}

	queue.pop()
	if queue.front() != 1 {
		t.Fatalf("1 is expected, but %v\n", queue.front())
	}

	queue.pop()
	if queue.front() != -5 {
		t.Fatalf("-5 is expected, but %v\n", queue.front())
	}

	queue.pop()
	if !queue.isEmpty() {
		t.Fatalf("True is expected, but %v\n", queue.isEmpty())
	}
}

func TestPairQueue(t *testing.T) {
	q := newPairQueue(100)
	fmt.Println(q)

	if !q.isEmpty() {
		t.Fatalf("True is expected, but %v\n", q.isEmpty())
	}

	q.push(&pair{1, 2})
	fmt.Println(q)
	q.push(&pair{-3, 1})
	fmt.Println(q)
	q.push(&pair{-1, -4})
	fmt.Println(q)

	if q.isEmpty() {
		t.Fatalf("False is expected, but %v\n", q.isEmpty())
	}

	p1 := q.front()
	fmt.Println(q)
	if p1.first != 1 || p1.second != 2 {
		t.Fatalf("%v is expected, but %v\n", &pair{1, 2}, p1)
	}
	q.pop()

	p2 := q.front()
	fmt.Println(q)
	if p2.first != -3 || p2.second != 1 {
		t.Fatalf("%v is expected, but %v\n", &pair{-3, 1}, p2)
	}
	q.pop()

	p3 := q.front()
	fmt.Println(q)
	if p3.first != -1 || p3.second != -4 {
		t.Fatalf("%v is expected, but %v\n", &pair{-1, -4}, p3)
	}
	q.pop()

	fmt.Println(q)
	if !q.isEmpty() {
		t.Fatalf("True is expected, but %v\n", q.isEmpty())
	}
}

func TestSet(t *testing.T) {
	s := newSet()
	fmt.Println(s)

	s.insert(10)
	if !s.contains(10) {
		t.Fatalf("True is expected, but %v\n", s.contains(10))
	}
	fmt.Println(s)
	s.insert(-3)
	if !s.contains(-3) {
		t.Fatalf("True is expected, but %v\n", s.contains(10))
	}
	fmt.Println(s)

	// duplicated
	s.insert(-3)
	fmt.Println(s)

	s.erase(-3)
	if s.contains(-3) {
		t.Fatalf("True is expected, but %v\n", s.contains(-3))
	}
	fmt.Println(s)
	s.erase(10)
	if s.contains(10) {
		t.Fatalf("True is expected, but %v\n", s.contains(10))
	}
	fmt.Println(s)
}

func TestIntHeap(t *testing.T) {
	h := &intHeap{2, 1, 5}
	heap.Init(h)
	heap.Push(h, 3)
	fmt.Println(*h)
	if (*h)[0] != 1 {
		t.Fatalf("1 is expected, but %v\n", (*h)[0])
	}

	i := 0
	expected := []int{1, 2, 3, 5}
	for h.Len() > 0 {
		if v := heap.Pop(h); v != expected[i] {
			t.Fatalf("%v is expected, but %v\n", expected[i], v)
		}
		i++
	}
}

func TestPriorityQueue(t *testing.T) {
	pq := make(priorityQueue, 0, 1024)
	heap.Init(&pq)
	heap.Push(&pq, &pair{first: 3, second: 1})
	heap.Push(&pq, &pair{first: 10, second: -3})
	heap.Push(&pq, &pair{first: 1, second: 4})
	fmt.Println(pq)

	if pq[0].p.first != 1 {
		t.Fatalf("1 is expected, but %v\n", pq[0].p.first)
	}

	i := 0
	expected := []*pair{
		{1, 4},
		{3, 1},
		{10, -3},
	}
	for pq.Len() > 0 {
		v := heap.Pop(&pq).(*pair)
		if *v != *expected[i] {
			t.Fatalf("%v is expected, but %v\n", expected[i], v)
		}
		fmt.Println(pq)
		i++
	}
}
