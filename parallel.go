package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	go one()
	go two()
	three()
	fmt.Println(t.Sub(time.Now()))
}

func one() {
	time.Sleep(1 * time.Second)
	fmt.Println("one")
}

func two() {
	time.Sleep(2 * time.Second)
	fmt.Println("two")
}

func three() {
	time.Sleep(3 * time.Second)
	fmt.Println("three")
}
