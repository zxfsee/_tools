package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	problems := 794
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)
	fmt.Print(random.Intn(problems) + 1)
}