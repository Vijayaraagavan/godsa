package src

import (
	"fmt"
	"runtime"
	"strconv"
	"time"
	"unsafe"
)

func Start() {
	open, title, createAt := getConfig()
	fmt.Println(open, title, createAt)
	verifyPointer()
	// sort([]int{5, 8, 2, 4, 0, 1, 3, 7, 9, 6})
	byteTest()
}
func getConfig() (bool, string, time.Time) {
	return false, "info", time.Now()
}
func verifyPointer() {
	var a *int = new(int)
	b := 5
	a = &b
	fmt.Println(*a)
	addr := fmt.Sprintf("%x", a)
	hexAddr, err := strconv.ParseInt(addr, 16, 64)
	if err != nil {
		fmt.Println("conversion error", err)
		return
	}
	fmt.Println(addr, hexAddr)
	an := uintptr(hexAddr)
	ptr := (*int)(unsafe.Pointer(an))
	*ptr = 16
	fmt.Println(*a)
}

func sort(a []int) {
	var l int = len(a)
	for i := 0; i < l; i++ {
		var target = l - i - 1
		var max = a[target]
		for j := 0; j < (l - i); j++ {
			if a[j] > max {
				a[j], a[target], max = max, a[j], a[j]
				// max = a[j]
			}
		}
	}
	fmt.Println(a)
}

func Memtest() {
	const l = 10000000
	var list = make([]int, 1, l+1)
	// var list []int
	for i := 0; i < l; i++ {
		list = append(list, 100)
	}
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("TotalAlloc (Heap) = %v MiB\n", m.TotalAlloc/1024/1024)
}
func byteTest() {
	l1, l2, l3 := linked()
	fmt.Println("Linked :", l1, l2, l3)
	nl1, nl2 := noLink()
	fmt.Println("No Link :", nl1, nl2)
	cl1, cl2 := capLinked()
	fmt.Println("Cap Link:", cl1, cl2)
	cnl1, cnl2 := capNoLink()
	fmt.Println("Cap No Link :", cnl1, cnl2)
	copy1, copy2, copied := copyNoLink()
	fmt.Print("Copy No Link: ", copy1, copy2)
	fmt.Printf(" (Number of elements copied %v)\n", copied)
	a1, a2 := appendNoLink()
	fmt.Println("Append No Link:", a1, a2)
}

func appendNoLink() (int, int) {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := append([]int{}, s1...)
	s1[3] = 99
	return s1[3], s2[3]
}

func copyNoLink() (int, int, int) {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := make([]int, len(s1))
	copied := copy(s2, s1)
	s1[3] = 99
	return s1[3], s2[3], copied
}

func capNoLink() (int, int) {
	s1 := make([]int, 5, 10)
	s1[0], s1[1], s1[2], s1[3], s1[4] = 1, 2, 3, 4, 5
	s2 := s1
	s1 = append(s1, []int{10: 11}...)
	s1[3] = 99
	return s1[3], s2[3]
}

func capLinked() (int, int) {
	s1 := make([]int, 5, 10)
	s1[0], s1[1], s1[2], s1[3], s1[4] = 1, 2, 3, 4, 5
	s2 := s1
	s1 = append(s1, 6)
	s1[3] = 99
	return s1[3], s2[3]
}

func linked() (int, int, int) {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := s1
	s3 := s1[:]
	s1[3] = 99
	return s1[3], s2[3], s3[3]
}

func noLink() (int, int) {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := s1
	s1 = append(s1, 6)
	s1[3] = 99
	return s1[3], s2[3]
}
