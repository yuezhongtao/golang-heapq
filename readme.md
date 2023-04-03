## Heap queue algorithm (a.k.a. priority queue).

Heaps are arrays for which a[k] <= a[2*k+1] and a[k] <= a[2*k+2] for
all k, counting elements from 0.  For the sake of comparison,
non-existing elements are considered to be infinite.  The interesting
property of a heap is that a[0] is always its smallest element.

### Usage Example:
```
import (
	golangheapq "golang-heapq"
)

h := golang_heapq.NewHeapQ[int]()            // creates an empty heap
h.Push(1)                                    // pushes a new item on the heap
h.Push(3)
h.Push(67)
h.Push(6)
h.Push(3457)
h.Push(5734)
h.Push(4)
h.Push(89)
h.Push(34)
h.PrettyPrint()
```

**output**:
```
          _____1_____     
         /           \    
     ___3_          __4   
    /     \        /   \  
_6      3457   5734   67
/  \                     
89   34                   
```   


