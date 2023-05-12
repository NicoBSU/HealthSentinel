package app

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

const idleTimeout = 30 * time.Second
const readTimeout = 10 * time.Second

func Setup() {
	//Fiber app init
	fiberApp := fiber.New(fiber.Config{
		IdleTimeout: idleTimeout,
		ReadTimeout: readTimeout,
	})

	go func() {
		listen(fiberApp)
	}()

	time.Sleep(1000)
}

func listen(fiberApp *fiber.App) {
	fiberApp.Listen(":8080")
}
