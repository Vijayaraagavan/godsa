package algorithm

import "fmt"

var arr [][]int = [][]int{}

func inits() {
	arr = append(arr, []int{0, 1, 1, 0})
	arr = append(arr, []int{0, 0, 0, 0})
	arr = append(arr, []int{0, 1, 0, 0})
	arr = append(arr, []int{1, 1, 0, 0})

}

func directed_adjacent_list() {
	fmt.Println(arr)
}

func Form_graph(mode int) {
	inits()
	switch mode {
	case 1:
		directed_adjacent_list()
	}
}
