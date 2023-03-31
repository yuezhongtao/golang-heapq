package golang_heapq

import "fmt"

type HeapQ[T Ordered] struct {
	heap []T
}

func NewHeapQ[T Ordered]() *HeapQ[T] {
	return &HeapQ[T]{}
}

func NewHeapQWithList[T Ordered](l []T) *HeapQ[T] {
	heap := &HeapQ[T]{heap: l}

	return heap
}

func (h *HeapQ[T]) Push(item T) {
	h.heap = append(h.heap, item)
	h.siftDown(0, h.len()-1)
}

func (h *HeapQ[T]) Pop() *T {
	if h.isEmpty() {
		return nil
	}
	lastelt := h.heap[h.len()-1]
	if !h.isEmpty() {
		h.heap = h.heap[0 : h.len()-1]
		returnitem := h.heap[0]
		h.heap[0] = lastelt
		h.siftUp(0)
		return &returnitem
	}
	return &lastelt
}

func (h HeapQ[T]) Replace(item T) *T {
	if h.isEmpty() {
		return nil
	}
	returnitem := h.heap[0]
	h.heap[0] = item
	h.siftUp(0)
	return &returnitem
}

func (h HeapQ[T]) PushPop(item T) *T {
	return nil
}

func (h HeapQ[T]) isEmpty() bool {
	return h.heap == nil || len(h.heap) == 0
}

func (h *HeapQ[T]) len() int {
	return len(h.heap)
}

func (h *HeapQ[T]) siftDown(startpos int, pos int) {
	newitem := h.heap[pos]
	for pos > startpos {
		parentpos := (pos - 1) >> 1
		parent := h.heap[parentpos]
		if newitem < parent {
			h.heap[pos] = parent
			pos = parentpos
			continue
		}
		break
	}
	h.heap[pos] = newitem
}

func (h *HeapQ[T]) siftUp(pos int) {
	endpos := h.len()
	startpos := pos
	newitem := h.heap[pos]
	// Bubble up the smaller child until hitting a leaf.
	childpos := 2*pos + 1
	for childpos < endpos {
		rightpos := childpos + 1
		if rightpos < endpos && !(h.heap[childpos] < h.heap[rightpos]) {
			childpos = rightpos
		}
		h.heap[pos] = h.heap[childpos]
		pos = childpos
		childpos = 2*pos + 1
	}
	h.heap[pos] = newitem
	h.siftDown(startpos, pos)
}

func (h HeapQ[T]) String() string {
	return fmt.Sprintf("%v", h.heap)
}

func (h HeapQ[T]) PrettyPrint() {
	PrettyPrint(h.heap)
}
