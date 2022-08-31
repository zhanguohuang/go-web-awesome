package handler

import "fmt"

type Inter interface {
	Bar(x int) int
}

func SUT(i Inter) {
	fmt.Println("ss")
}
