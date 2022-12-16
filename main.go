package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"math/rand"

	"github.com/go-vgo/robotgo"
)

func main() {
	fmt.Println("mousemove")

	var sec time.Duration = 5
	secVal, ok := os.LookupEnv("MOUSEMOVE_INTERVAL")
	if ok {
		temp, err := strconv.Atoi(secVal)
		if err != nil {
			panic(err)
		}
		sec = time.Duration(temp)
	}
	fmt.Printf("interval: %d\n", sec)

	for {
		x := rand.Intn(10000)
		y := rand.Intn(10000)
		fmt.Printf("move mouse x: %d y: %d\n", x, y)
		robotgo.Move(x, y)
		time.Sleep(time.Duration(sec) * time.Second)
	}
}
