package main

import (
	"context"
	"fmt"
	"os"
)

func main() {
	component := hello("World")
	component.Render(context.Background(), os.Stdout)
	fmt.Println()
}
