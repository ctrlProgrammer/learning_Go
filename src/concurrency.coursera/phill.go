package main

import (
	"fmt"
	"sync"
	"time"
)

type ChopS struct{ sync.Mutex }

type Philo struct {
	index      int
	leftStick  *ChopS
	rightStick *ChopS
}

func (p Philo) eat(host *chan bool, wg *sync.WaitGroup) {
	for i := 0; i < 3; i++ {
		*host <- true
		fmt.Println("Starting to eat: ", p.index+1)
		p.leftStick.Lock()
		p.rightStick.Lock()
		time.Sleep(1 * time.Second)
		fmt.Println("finishing eating: ", p.index+1)
		p.leftStick.Unlock()
		p.rightStick.Unlock()
		<-*host
	}

	wg.Done()
}

func initialization() (philos []*Philo, host chan bool) {
	sticks := make([]*ChopS, 5)
	philos = make([]*Philo, 5)
	host = make(chan bool, 2)

	for i := 0; i < 5; i++ {
		sticks[i] = new(ChopS)
	}

	for i := 0; i < 5; i++ {
		philos[i] = &Philo{i, sticks[i], sticks[(i+1)%5]}
	}

	return philos, host
}

func main() {
	var wg sync.WaitGroup

	philos, host := initialization()

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go philos[i].eat(&host, &wg)
	}

	wg.Wait()
}
