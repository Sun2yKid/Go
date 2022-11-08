package main


import (
	"log"
	"sync"
	"time"

	"work"
)

var names = []string {
	"steve",
	"bob",
	"mary",
	"therese",
	"jason",
}

type namePriter struct {
	name string
}


func (m *namePriter) Task() {
	log.Println(m.name)
	time.Sleep(time.Second)
}


func main() {
	p := work.New(2)

	var wg sync.WaitGroup
	wg.Add(100 * len(names))

	for i := 0; i < 100; i++ {
		for _, name := range names {
			np := namePriter{
				name: name,
			}
			go func ()  {
				p.run(&np)
				wg.Done()
			}()
		}
	}

	wg.Wait()

	p.Shutdown()
}

