package main

import (
	"fmt"
	"github.com/jordanferenz/gooss2-demo/light"
)

func main() {
	fmt.Printf("Is light a particle or a wave?\n\n>>> It is a %s\n\n\n\n", light.Observe())
}
