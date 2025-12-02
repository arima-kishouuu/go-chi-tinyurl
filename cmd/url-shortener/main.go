package main

import (
	"fmt"

	config "example.com/urlshortner/internal/comfig"
)

func main() {

	cfg := config.MustLoad()

	fmt.Println(cfg)
}
