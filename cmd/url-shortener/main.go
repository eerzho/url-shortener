package main

import (
	"fmt"

	"url-shortener-2/internal/config"
)

func main() {
	cfg := config.MustLoad()

	fmt.Println(cfg)

	// todo init logger slog

	// todo init storage

	// todo init router

	// todo run server
}
