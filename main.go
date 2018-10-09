package main

import (
	"context"
	"fmt"

	"github.com/google/go-cloud/wire"
)

type Foo struct {
	X int
}

type Bar struct {
	X int
}

type FooBar struct {
	X int
}

func ProvideFoo() Foo {
	return Foo{X: 42}
}

func ProvideBar(foo Foo) Bar {
	return Bar{X: foo.X + 12}
}

func ProvideFooBar(foo Foo, bar Bar) FooBar {
	return FooBar{X: foo.X + bar.X}
}

var SuperSet = wire.NewSet(ProvideFoo, ProvideBar, ProvideFooBar)

func main() {
	bar, err := initBar(context.TODO())
	if err != nil {
		panic(err)
	}

	fmt.Println("BAR:", bar)
}
