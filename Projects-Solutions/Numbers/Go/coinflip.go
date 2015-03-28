package main

import (
	"fmt"
	"math/rand"
	"time"
)

//Coin Flip Simulation:
//Write some code that simulates flipping a single coin however many times
//the user decides. The code should record the outcomes and count the number of tails and heads.

//Returns random number between 1 and 2
func flip() int {
	return 1 + rand.Int()%2
}

func main() {
	var num int
	heads := 0
	tails := 0

	rand.Seed(time.Now().Unix())

	fmt.Print("Enter number of times to flip coin: ")
	_, err := fmt.Scanf("%d", &num)

	if err != nil {
		panic(err)
	}

	for i := 0; i < num; i++ {
		f := flip()
		if f == 1 { /*1 == heads*/
			heads++
		} else if f == 2 { /* 2 == tails*/
			tails++
		}
	}

	fmt.Println("The number of heads is", heads)
	fmt.Println("The number of tails is", tails)
}
