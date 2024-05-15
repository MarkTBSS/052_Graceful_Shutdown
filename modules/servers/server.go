package servers

import (
	"encoding/json"
	"log"
	"os"
	"os/signal"

	_pkgConfig "github.com/MarkTBSS/052_Graceful_Shutdown/config"
	"github.com/gofiber/fiber/v2"
)

type IServer interface {
	Start()
}

type server struct {
	app *fiber.App
	cfg _pkgConfig.IConfig
}

func (s *server) Start() {
	// Graceful Shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		log.Println("server is shutting down...")
		_ = s.app.Shutdown()
	}()

	// Listen to host:port
	log.Printf("server is starting on %v", s.cfg.App().Url())
	s.app.Listen(s.cfg.App().Url())
}

func NewServer(cfg _pkgConfig.IConfig) IServer {
	return &server{
		cfg: cfg,
		app: fiber.New(fiber.Config{
			JSONEncoder: json.Marshal,
			JSONDecoder: json.Unmarshal,
		}),
	}
}
