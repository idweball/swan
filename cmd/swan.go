package main

import (
	"swan/internal/app"
	_ "swan/pkg/storage/redis"
)

func main() {
	app.Run()
}
