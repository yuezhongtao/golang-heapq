package main

import (
	golangheapq "golang-heapq"
)

func main() {
	h := golangheapq.NewHeapQ[int]()
	h.Push(1)
	h.Push(3)
	h.Push(67)
	h.Push(6)
	h.Push(3457)
	h.Push(5734)
	h.Push(4)
	h.Push(89)
	h.Push(34)
	h.PrettyPrint()
}
