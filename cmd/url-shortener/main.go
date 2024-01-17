package main

import (
	"fmt"
	"url-shortner-2/internal/config"
)

func main() {
	cfg := config.MustLoad()

	fmt.Println(cfg.IdleTimeout)

	// todo init logger slog

	// todo init storage

	// todo init router

	// todo run server
}
