package src

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
	"unicode"
)

func Valid(id string) (resp bool, pure string) {
	resp = true
	// var pure string
	for _, v := range id {
		if !unicode.IsDigit(v) && !unicode.IsSpace(v) {
			fmt.Println("setting sfalse")
			resp = false
			break
		} else if unicode.IsSpace(v) {
			continue
		}
		pure += string(v)
	}
	fmt.Println(pure, resp)

	return resp, doubleSecond(pure)
}

func doubleSecond(id string) string {
	var l int = len(id) - 1
	var resp []string = make([]string, 0)
	for i := l; i >= 0; i-- {
		fmt.Printf("index: %d, len: %d, value: %s end", i, l, string(id[i]))
		fmt.Println(resp)
		v, err := strconv.Atoi(string(id[i]))
		if err != nil {
			fmt.Println("s to i failed")
			return ""
		}
		if (l-i)%2 != 0 {
			if (v * 2) > 9 {
				v = (v * 2) - 9
			} else {
				v = v * 2
			}
			resp = prepend(resp, fmt.Sprint(v))
		} else {
			resp = prepend(resp, string(id[i]))
		}
	}
	fmt.Println(resp)
	return strings.Join(resp, "")
}

func prepend(src []string, v string) []string {
	return append([]string{v}, src...)
}

type Anime struct {
	Name     string
	volumes  int
	chapters int
}

func (a Anime) chaptersPerVolume() (r int) {
	return a.chapters / a.volumes
}

func Manipulate(a *int) {
	// defer fmt.Println("ending")
	// defer func(a int) {
	// 	fmt.Println(a * a)
	// }(5)
	defer func() {
		msg := recover()
		if msg != nil {
			fmt.Println(msg)
			fmt.Println("return to")
		}
	}()
	*a += 10
	fmt.Println(&a)
	var n *int = new(int)
	*n = 50
	fmt.Println(*n)

	c := "vijay ra"
	if strings.Contains(c, "ra") {
		fmt.Println("present")
	} else {

		fmt.Println("not present")
	}
	var arr [3]int = [3]int{2, 10}
	fmt.Println(arr)
	var multi = [3][2]int{{1, 2}, {7, 6}, {16, 250}}
	for i, v := range multi {
		fmt.Printf("Index: %d; value: %v\n", i, v)
		for i1, v1 := range v {
			fmt.Printf("inner => Index: %d; value: %v\n", i1, v1)
		}
	}
	fmt.Println(multi)
	ar := arr
	var slice []int = ar[:]
	slice[2] = 20
	fmt.Printf("slice: %v; len: %d; cap: %d\n", slice, len(slice), cap(slice))
	fmt.Println(ar, arr)
	// rand.Seed(time.Now().UnixNano())
	boruto := Anime{Name: "boruto", chapters: 80, volumes: 18}
	fmt.Println("total", boruto.chaptersPerVolume())
	jjk := new(Anime)
	jjk.Name = "Jujutsu kaisen"
	jjk.chapters = 255
	jjk.volumes = 24

	fmt.Println("total", jjk.chaptersPerVolume())
	mainChannel := make(chan int)

	go producer(mainChannel)
	select {
	case <-mainChannel:
		fmt.Println("finished job")
	case <-time.After(time.Second * 7):
		fmt.Println("timeout")
	}
	// case <-
	var wg sync.WaitGroup
	runtime.GOMAXPROCS(3)

	fmt.Printf("default cpus: %d", runtime.NumCPU())
	wg.Add(limit + 1)

	testit(&wg)
	wg.Wait()
}

var limit int = 5

func producer(main chan int) {
	defer func() {
		msg := recover()
		fmt.Println("panic ocscured", msg)
		main <- 1
	}()
	ch := make(chan int)
	go consumer(ch)
	for i := 0; i < 1; i++ {
		ch <- i * 2
	}
	main <- 1

}

func consumer(c chan int) {
	for i := 0; i < 10; i++ {
		if i == 5 {
			time.Sleep(time.Second * 3)
			close(c)
		}
		v, ok := <-c
		fmt.Printf("receiving %d", v)
		if !ok {
			return
		}
	}
}

func testit(wg *sync.WaitGroup) {
	for i := 0; i < limit; i++ {
		go stress(wg, i)
	}
}

func stress(wg *sync.WaitGroup, parent int) {
	i := 0
	for j := 0; j < 10000000; j++ {
		i = i + 10
		// fmt.Printf("i = %d; j: %d; parent: %d\n", i, j, parent)
	}
	wg.Done()
}
