// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

import (
	"context"
)

// Injectors from inject.go:

func initFoo(ctx context.Context) (Foo, error) {
	foo := ProvideFoo()
	return foo, nil
}
