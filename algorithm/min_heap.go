package algorithm

import "fmt"

func Min_heap(a []int) {
	fmt.Println(a)
	heapify(a, len(a)-1)
	insert(&a, 3)
	fmt.Println("after insert", a)
	delete(a)
}

func parent(i int) int {
	return (i - 1) / 2
}
func left(i int) int {
	return (2 * i) + 1
}
func right(i int) int {
	return (2 * i) + 2
}

/*
You can ignore return value from toptoBottom func. because slice modification will reflect in original array
when slice is passed to func, its descriptor is copied, including underlying array pointer.
but if you change length of slice like appending, reference to array is cut in called func in turn
not reflecting in caller func
*/
func heapify(a []int, index int) {
	for i := (len(a) / 2) - 1; i >= 0; i-- {
		// fmt.Printf("loop index: %d\n", i)
		topToBottom(a, i)
	}
}

func insert(a *[]int, el int) {
	// fmt.Println("before insert", *a)
	*a = append(*a, el)
	bottomToTop(*a, len(*a)-1)
	fmt.Println(*a)
}

func delete(a []int) {
	a[0] = a[len(a)-1]
	a = a[:len(a)-1]
	topToBottom(a, 0)
	// fmt.Println("after delete", a)
}

func bottomToTop(a []int, index int) []int {
	if index <= 0 {
		return a
	}
	parent := parent(index)
	if a[parent] > a[index] {
		a[parent], a[index] = a[index], a[parent]
	}
	// fmt.Printf("for index: %d, arr: %v\n", index, a)
	bottomToTop(a, parent)
	return a
}
func topToBottom(a []int, index int) []int {
	left := left(index)
	right := right(index)
	smallest := index
	if left < len(a) && a[smallest] > a[left] {
		// a[left], a[index] = a[index], a[left]
		smallest = left
	}
	if right < len(a) && a[smallest] > a[right] {
		// a[right], a[index] = a[index], a[right]
		smallest = right
	}
	// fmt.Printf("for index: %d, arr: %v\n", index, a)
	if smallest != index {
		a[smallest], a[index] = a[index], a[smallest]
		topToBottom(a, smallest)
		return a
	}
	return a
}
