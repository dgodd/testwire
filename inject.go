// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package main

import (
	"context"

	"github.com/google/go-cloud/wire"
)

func initFoo(ctx context.Context) (Foo, error) {
	wire.Build(SuperSet)
	return Foo{}, nil
}
