package test

import "log"

const Version = "v1.2"

func Add(x, y int) int {
	log.Println("You are using test.Add")
	return x + y
}
