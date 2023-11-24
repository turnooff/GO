package main

import (
	"fmt"
	"sync"
	"time"
)

type BoxOffice struct {
	m      sync.Mutex
	count  map[int]int
	status map[int]string
}

type BoxOfficeWithoutMutex struct {
	count  map[int]int
	status map[int]string
}

func (boxOffice *BoxOffice) value(index int) int {
	boxOffice.m.Lock()
	defer boxOffice.m.Unlock()
	return boxOffice.count[index]
}

func (boxOffice *BoxOfficeWithoutMutex) value(index int) int {
	return boxOffice.count[index]
}

func (boxOffice *BoxOffice) work(index int) {
	for i := 1; i <= 5; i++ {
		ch := make(chan int)
		go func(index int) {
			ch <- index
		}(i)
		go boxOffice.printWorkStatus(index, ch)
		if i == 1 {
			time.Sleep(2 * time.Second)
		}

		if !(i == 5 && index == 3) {
			fmt.Printf("На кассе №%v покупатель %v\n", index, i)
			go boxOffice.inc(index)
		}
		time.Sleep(time.Second)
	}
	fmt.Printf("На кассе №%v побывало %v\n", index, boxOffice.value(index))

}

func (boxOffice *BoxOffice) inc(key int) {
	boxOffice.m.Lock()
	boxOffice.count[key]++
	boxOffice.m.Unlock()
}

func (boxOffice *BoxOfficeWithoutMutex) inc(key int) {
	boxOffice.count[key]++
}

func (boxOffice *BoxOfficeWithoutMutex) work(index int) {
	for i := 1; i <= 5; i++ {
		ch := make(chan int)
		go func(index int) {
			ch <- index
		}(i)
		go boxOffice.printWorkStatus(index, ch)
		if i == 1 {
			time.Sleep(2 * time.Second)
		}

		if !(i == 5 && index == 3) {
			fmt.Printf("На кассе №%v покупатель %v\n", index, i)
			go boxOffice.inc(index)
		}
		time.Sleep(time.Second)

	}

	fmt.Printf("На кассе №%v побывало %v\n", index, boxOffice.value(index))
}

func (boxOffice *BoxOfficeWithoutMutex) printWorkStatus(index int, ch chan int) {
	var i int = <-ch
	helpStr := boxOffice.workStatus(index, i)
	if helpStr == "start" {
		fmt.Printf("Касса %d начала работу \n", index)
	} else if helpStr == "end" {
		fmt.Printf("Касса %d завершила работу \n", index)
	}
}

func (boxOffice *BoxOfficeWithoutMutex) workStatus(index int, i int) string {
	if i == 1 {
		boxOffice.status[index] = "start"
		return "start"
	} else if i == 5 {
		boxOffice.status[index] = "end"
		return "end"
	}
	return "work"
}

func (boxOffice *BoxOffice) workStatus(index int, i int) string {
	boxOffice.m.Lock()
	defer boxOffice.m.Unlock()
	if i == 1 {
		boxOffice.status[index] = "start"
		return "start"
	} else if i == 5 {
		boxOffice.status[index] = "end"
		return "end"
	}
	return "work"
}

func (boxOffice *BoxOffice) printWorkStatus(index int, ch chan int) {
	var i int = <-ch
	helpStr := boxOffice.workStatus(index, i)
	if helpStr == "start" {
		fmt.Printf("Касса %d начала работу \n", index)
	} else if helpStr == "end" {
		fmt.Printf("Касса %d завершила работу \n", index)
	}
}

func main() {
	mu := BoxOffice{count: make(map[int]int), status: make(map[int]string)}
	for i := 1; i < 11; i++ {
		go mu.work(i)
	}
	time.Sleep(10 * time.Second)

	// fmt.Print("\nБез использования примитивов синхронизации:\n\n")
	// some := BoxOfficeWithoutMutex{count: make(map[int]int), status: make(map[int]string)}
	// for i := 0; i < 10; i++ {
	// 	go some.work(i + 1)
	// }
	// time.Sleep(10 * time.Second)
}
