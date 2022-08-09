package main

import (
	"fmt"
	"sync"
	"time"
)

const hunger = 3

type Fork struct {
	idx int
	sync.Mutex
}

type philosopher struct {
	p_idx int
	left  *Fork
	right *Fork
}

var wg sync.WaitGroup
var sleepTime = 1 * time.Second
var eatTime = 2 * time.Second
var thinkTime = 1 * time.Second
var num int = 5
var orderedList []int
var is_valid int

func diningProblme(p *philosopher, ch chan int) {
	defer wg.Done()
	fmt.Println(p.p_idx, "is seated")
	//time.Sleep(sleepTime)
	for i := hunger; i > 0; i-- {
		fmt.Println(p.p_idx, "is hungry")
		//time.Sleep(sleepTime)

		p.left.Lock()
		fmt.Println(p.p_idx, " is using left fork")
		p.right.Lock()
		fmt.Println(p.p_idx, " is using right fork")
		fmt.Println(p.p_idx, "is using both fork")
		//		time.Sleep(eatTime)

		// thinking time
		fmt.Println(p.p_idx, " is thinking.")
		//		time.Sleep(thinkTime)

		p.left.Unlock()
		fmt.Println(p.p_idx, " stop using left fork")
		p.right.Unlock()
		fmt.Println(p.p_idx, " stop using right fork")

		//		time.Sleep(sleepTime)
	}

	fmt.Println(p.p_idx, "is satisfied")
	// time.Sleep(sleepTime)
	fmt.Println(p.p_idx, "has left the table")
	ch <- p.p_idx
}

func main() {
	num := 5
	//initilize forks for all table
	forks := make([]Fork, num)
	for i := 0; i < num; i++ {
		forks[i] = Fork{
			idx: i,
		}
	}

	//assign left and right fork for each philosopher
	philosophers := make([]philosopher, num)
	for i := 0; i < num; i++ {
		philosophers[i] = philosopher{
			p_idx: i + 1,
			left:  &forks[i],
			right: &forks[(i+1)%num],
		}
	}

	wg.Add(len(philosophers))
	orders := make(chan int, len(philosophers))
	for i := 0; i < num; i++ {
		go diningProblme(&philosophers[i], orders)
	}
	wg.Wait()

	is_valid = 0
	for i := 0; i < num; i++ {
		index := <-orders
		is_valid += index
		fmt.Print(index)
	}
	close(orders)
	fmt.Println(is_valid)
	fmt.Println("The table is empty")
}
