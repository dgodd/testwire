package main

import (
	"context"
	"fmt"

	"github.com/google/go-cloud/wire"
)

type Foo struct {
	X int
}

func ProvideFoo() Foo {
	return Foo{X: 42}
}

var SuperSet = wire.NewSet(ProvideFoo)

func main() {
	foo, err := initFoo(context.TODO())
	if err != nil {
		panic(err)
	}

	fmt.Println("FOO:", foo)
}
