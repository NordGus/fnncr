package main

import (
	"context"
	"fmt"

	"github.com/NordGus/fnncr/authentication"
)

func main() {
	ctx, cancel := context.WithCancelCause(context.Background())

	auth := authentication.New(ctx, cancel)

	fmt.Println(auth)
}
