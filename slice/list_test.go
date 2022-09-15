package slice

import (
	"fmt"
	"testing"
)

type Foo struct {
}

func TestAppendNil(t *testing.T) {
	FooList := []*Foo{{}}
	FooList = append(FooList, nil)
	fmt.Println(FooList)
}
