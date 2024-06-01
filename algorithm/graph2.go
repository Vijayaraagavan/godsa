package algorithm

import (
	"fmt"
)

type Node struct {
	Value int
	Next  *Node
}

func StartGraph() {
	// fd, err := os.Create("main.prof")
	// if err != nil {
	// 	fmt.Println("not able to create file")
	// 	return
	// }
	// pprof.StartCPUProfile(fd)
	// pprof.StopCPUProfile()
	// log.Println(http.ListenAndServe("localhost:6060", nil))
	// input := [][]int{{11, 10}, {11, 12}, {12, 10}}
	input := [][]int{{0, 1}, {0, 2}, {1, 3}, {2, 4}, {3, 4}}
	n := 5
	fmt.Println(input)
	gRep := graphRep(input, n)
	fmt.Println(gRep)
	bfsGraph(gRep, n)
}

func graphRep(in [][]int, totalVertices int) []Node {
	var out []Node
	for i := 0; i < totalVertices; i++ {
		edge := in[i]
		v1, v2 := edge[0], edge[1]
		insertNode(&out, v1, v2)
	}
	return out
}

func insertNode(out *[]Node, v1 int, v2 int) {
	idx := getNode(*out, v1)
	if idx != -1 {
		// if !vertexExist((*out)[idx], v2) {
		// }
		appendVertex(&(*out)[idx], v2)
	} else {
		head := Node{Value: v1}
		head.Next = &Node{Value: v2}
		*out = append(*out, head)
	}
	traverse(*out)
	// if (*out)[0].Next == nil {
	// 	// head := Node{Value: v1}
	// 	// *out = append(*out, head)
	// 	fmt.Println(*out)
	// }
}

func traverse(out []Node) {
	fmt.Println(out)
	for i := 0; i < len(out); i++ {
		var p Node = out[i]
		fmt.Printf("value: %d; ", p.Value)
		for p.Next != nil {
			p = *p.Next
			fmt.Printf("value: %d; ", p.Value)
		}
		fmt.Println()
	}
}

func getNode(out []Node, target int) int {
	for i := 0; i < len(out); i++ {
		if out[i].Value == target {
			return i
		}
	}
	return -1
}
func vertexExist(v Node, target int) bool {
	for v.Next != nil {
		if v.Value == target {
			return true
		}
		v = *v.Next
	}
	return false
}
func appendVertex(v *Node, target int) {
	var p *Node = v
	for (*p).Next != nil {
		p = (*p).Next
	}
	// fmt.Printf("adding %d to vertex %d\n", target, v.Value)
	p.Next = &Node{Value: target}
}

func bfsGraph(in []Node, n int) {
	var q queue = make([]Node, 0)
	var h hashMap = make(map[int]bool)
	q.enqueue(in[0])
	fmt.Printf("traversal values: ")
	for len(q) != 0 {
		var node Node = q.dequeue()
		fmt.Printf("%d ", node.Value)
		idx := getNode(in, node.Value)
		if idx != -1 {
			var topNode Node = in[idx]
			for topNode.Next != nil {
				if !h.exist((*topNode.Next).Value) {
					h[(*topNode.Next).Value] = true
					q.enqueue(*topNode.Next)
				}
				topNode = *topNode.Next
			}
		}
	}
	fmt.Println(h)
}

type queue []Node

func (q *queue) enqueue(i Node) {
	*q = append(*q, i)
}
func (q *queue) dequeue() Node {
	if len(*q) <= 0 {
		panic("too small in queue")
	}
	var a Node = (*q)[0]
	*q = (*q)[1:]
	return a
}

// type hashMap map[Node]bool

//	func (h hashMap) exist(i Node) bool {
//		if h[i] == true {
//			return true
//		} else {
//			return false
//		}
//	}
type hashMap map[int]bool

func (h hashMap) exist(i int) bool {
	if h[i] {
		return true
	} else {
		return false
	}
}
