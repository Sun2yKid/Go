package main

import "fmt"
import "sync"
import "time"


func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum    // send sum to c
}

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func fibonacci2(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <- quit:
			fmt.Println("quit")
			return
		}
	}
}

// SafeCounter is safe to use concurrently
type SafeCounter struct {
	mu sync.Mutex
	v map[string]int
}

// Inc increments the counter for the given key.
func (c *SafeCounter) Inc(key string) {
	c.mu.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	c.v[key]++
	c.mu.Unlock()
}

// Value returns the current value of the counter for the given key.
func (c *SafeCounter) Value(key string) int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.v[key]
}


func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <- c, <-c

	fmt.Println(x, y, x+y)

	c2 := make(chan int, 10)
	go fibonacci(cap(c2), c2)
	for i := range c2 {
		fmt.Println(i)
	}

	c3 := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i<10; i++ {
			fmt.Println(<-c3)
		}
		quit <- 0
	}()
	fibonacci2(c3, quit)

	// sync.Mutex
	c4 := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c4.Inc("somekey")
	}

	time.Sleep(time.Second)
	fmt.Println(c4.Value("somekey"))
}
