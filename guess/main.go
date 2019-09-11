//Generate a random number from 1 to 100. This number is the target to be guessed
//Prompt the player to guess what the target number is
//If player's guess is less than target number say, "The guess was LOW"
//If player's guess is greater than target number say, "The gueass was HIGH"
//If player's gues is equal to target number say, "You have guessed the target number" and stop guessing
//The player is allowed to guess up to 10 times
//Let player know how many guesses they have left
//If after 10 times the player has not guessed the target number, the play ends, say "You did not guess the number"package main

package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	target := GetTarget(100)
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Guess number from 1 to 100")
	for i := 0; i < 10; i++ {
		fmt.Printf("Guess %d 0f 10 :", i+1)
		g, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		g = strings.TrimSpace(g)
		guess, err := strconv.ParseInt(g, 10, 64)
		if err != nil {
			fmt.Println("Dat ging niet goed !!!")
			log.Fatal(err)
		}

		if guess == target {
			fmt.Println("You have guessed the target number = ", target)
			break
		}
		if guess > target {
			fmt.Println("Your guess is too HIGH")
		} else {
			fmt.Println("Your guess is too LOW")
		}
	}
}

// GetTarget returns a random number from 0 to n
func GetTarget(n int) int64 {
	rand.Seed(time.Now().UnixNano())
	target := rand.Intn(n) + 1
	return int64(target)
}
