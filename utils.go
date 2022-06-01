package test

import "log"

const Version = "v1.0"

func Add(x, y int) int {
	log.Println("You are using test.Add")
	return x + y
}
