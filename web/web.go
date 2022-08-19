package web

import (
	"net"

	"github.com/arden144/axiom/log"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

var app = fiber.New(fiber.Config{
	JSONEncoder:           json.Marshal,
	JSONDecoder:           json.Unmarshal,
	DisableStartupMessage: true,
})

func init() {
	ln, err := net.Listen(app.Config().Network, ":3000")
	if err != nil {
		log.L.Fatal("failed to create listener", zap.Error(err))
	}

	app.Hooks().OnListen(func() error {
		log.L.Info("web server ready", zap.String("address", ln.Addr().String()))
		return nil
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("hello world")
	})

	go func() {
		if err := app.Listener(ln); err != nil {
			log.L.Fatal("failed to start web server", zap.Error(err))
		}
	}()
}
