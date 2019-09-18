//Same as tennis but without WaitGroup, channel only
package main

import (
	"fmt"
	"math/rand"
	"time"
)

//Player has a name
type Player struct {
	name string
}

//Play tennis.
func (p *Player) Play(c chan int, done chan bool) {
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
			done <- true
			return
		}
	}
}

//PlayTennis plays a tennis game
func PlayTennis(c chan int, done chan bool, p1 *Player, p2 *Player) {
	//Toss what player hits the first ball
	if rand.Intn(10)%2 == 0 {
		go p1.Play(c, done)
		go p2.Play(c, done)
	} else {
		go p2.Play(c, done)
		go p1.Play(c, done)
	}
	//Start the tennis game
	c <- 1
	//Tennis game has finished
	<-done
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	c := make(chan int)
	done := make(chan bool)

	player1 := Player{
		name: "Jan",
	}
	player2 := Player{
		name: "Anna-Maria",
	}

	fmt.Println("The tennis game has started")
	PlayTennis(c, done, &player1, &player2)
	fmt.Println("The tennisgame has finished")

}
