package amqprpc

import (
	"context"
	"fmt"

	"github.com/Jeanhwea/baliqiao/internal/entity"
	"github.com/Jeanhwea/baliqiao/internal/usecase"
	"github.com/Jeanhwea/baliqiao/pkg/rabbitmq/rmq_rpc/server"
	"github.com/streadway/amqp"
)

type translationRoutes struct {
	translationUseCase usecase.Translation
}

func newTranslationRoutes(routes map[string]server.CallHandler, t usecase.Translation) {
	r := &translationRoutes{t}
	{
		routes["getHistory"] = r.getHistory()
	}
}

type historyResponse struct {
	History []entity.Translation `json:"history"`
}

func (r *translationRoutes) getHistory() server.CallHandler {
	return func(d *amqp.Delivery) (interface{}, error) {
		translations, err := r.translationUseCase.History(context.Background())
		if err != nil {
			return nil, fmt.Errorf("amqp_rpc - translationRoutes - getHistory - r.translationUseCase.History: %w", err)
		}

		response := historyResponse{translations}

		return response, nil
	}
}
