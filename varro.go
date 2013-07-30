package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"time"
)

type SomePosition struct {
	start int
	end   int
}

func main() {
	file, err := ioutil.ReadFile(os.Args[1])

	positions := make([]SomePosition, 0, 100)

	if err != nil {
		fmt.Println("Error when trying to open file was:", err)
		return
	}

	start, end := -1, 0

	for index, somebyte := range file {
		if somebyte == 10 {
			end = index
			if end > start {
				positions = append(positions, SomePosition{start + 1, end})
				start = end
			}
		}
	}

	rand.Seed(time.Now().UTC().UnixNano())
	some_random := rand.Intn(len(positions))

	fmt.Println(string(file[positions[some_random].start:positions[some_random].end]))
}
