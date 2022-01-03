package main

import (
	"fmt"
	"sync"
	"time"
)

var diners sync.WaitGroup
var host = make(chan bool, 2)
var servings = 3
var eat = time.Second / 100
var dinersCount = 5
var names = []string{"J", "Li", "L", "C", "K"}

type Chopstick struct{ sync.Mutex }

type Philospher struct {
	name       string
	chopStick1 *Chopstick
	chopStick2 *Chopstick
}

func (p Philospher) dine() {
	for i := 0; i < servings; i++ {
		host <- true

		// pick a random chopstick first
		// rc := rand.Intn(2)
		// if rc == 0 {
		p.chopStick1.Lock()
		// defer p.chopStick1.Unlock()

		p.chopStick2.Lock()
		// defer p.chopStick2.Unlock()
		// } else {
		// 	p.chopStick2.Lock()
		// 	defer p.chopStick2.Unlock()

		// 	p.chopStick1.Lock()
		// 	defer p.chopStick1.Unlock()
		// }

		fmt.Println("Eating   : ", p.name)
		time.Sleep(eat)
		fmt.Println("Finished : ", p.name)

		p.chopStick1.Unlock()
		p.chopStick2.Unlock()

		<-host
	}

	diners.Done()
}

func main() {
	chopsticks := make([]*Chopstick, dinersCount)
	for i := 0; i < dinersCount; i++ {
		chopsticks[i] = new(Chopstick)
	}

	philosphers := make([]*Philospher, dinersCount)
	for i := 0; i < dinersCount; i++ {
		philosphers[i] = &Philospher{names[i], chopsticks[i], chopsticks[(i+1)%dinersCount]}
	}

	for i := 0; i < dinersCount; i++ {
		diners.Add(1)
		go philosphers[i].dine()
	}

	diners.Wait()
}
