// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package main

import (
	"context"

	"github.com/google/go-cloud/wire"
)

func initBar(ctx context.Context) (Bar, error) {
	wire.Build(SuperSet)
	return Bar{}, nil
}

func initFooBar(ctx context.Context) (FooBar, error) {
	wire.Build(SuperSet)
	return FooBar{}, nil
}
