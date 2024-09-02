package main

import (
	"fmt"
	"projectEuler/euler"
	"time"
)

func main() {
	fmt.Println(time.Now())
	fmt.Println(euler.PrintNthPrime(5))
	fmt.Println(time.Since(time.Now()))

}
