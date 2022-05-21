package amqprpc

import (
	"github.com/Jeanhwea/baliqiao/internal/usecase"
	"github.com/Jeanhwea/baliqiao/pkg/rabbitmq/rmq_rpc/server"
)

// NewRouter -.
func NewRouter(t usecase.Translation) map[string]server.CallHandler {
	routes := make(map[string]server.CallHandler)
	{
		newTranslationRoutes(routes, t)
	}

	return routes
}
