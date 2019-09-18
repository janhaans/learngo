package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

//Player has a name
type Player struct {
	name string
}

//Play tennis.
func (p *Player) Play(c chan int) {
	defer wg.Done()
	for {
		hitball, more := <-c
		if more {
			if rand.Intn(100)%7 == 0 {
				fmt.Printf("Player %s missed the ball\n", p.name)
				close(c)
				return
			}
			fmt.Printf("Player %s hits ball %d\n", p.name, hitball)
			hitball++
			time.Sleep(time.Second)
			c <- hitball
		} else {
			fmt.Printf("Player %s has won the game\n", p.name)
			return
		}
	}
}

var wg sync.WaitGroup

//StartGame start a new tennis game
func StartGame(c chan int, p1 *Player, p2 *Player) {
	if rand.Intn(10)%2 == 0 {
		go p1.Play(c)
		go p2.Play(c)
	} else {
		go p2.Play(c)
		go p1.Play(c)
	}
	c <- 1
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	wg.Add(2)
	c := make(chan int)

	player1 := Player{
		name: "Jan",
	}
	player2 := Player{
		name: "Anna-Maria",
	}

	fmt.Println("The game has started")
	StartGame(c, &player1, &player2)

	wg.Wait()
	fmt.Println("The game has finished")

}
