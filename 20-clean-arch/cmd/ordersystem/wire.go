//go:build wireinject
// +build wireinject

package main

import (
	"database/sql"

	"github.com/allanmaral/go-expert/20-clean-arch/internal/entity"
	"github.com/allanmaral/go-expert/20-clean-arch/internal/infra/database"
	"github.com/allanmaral/go-expert/20-clean-arch/internal/infra/eventhandlers"
	"github.com/allanmaral/go-expert/20-clean-arch/internal/infra/web"
	"github.com/allanmaral/go-expert/20-clean-arch/internal/usecase"
	"github.com/allanmaral/go-expert/20-clean-arch/pkg/events"
	"github.com/google/wire"
)

var setOrderRepositoryDependency = wire.NewSet(
	database.NewOrderRepository,
	wire.Bind(new(entity.OrderRepository), new(*database.OrderRepositorySQL)),
)

var setOrderCreatedEvent = wire.NewSet(
	eventhandlers.NewOrderCreatedEvent,
	wire.Bind(new(events.Event), new(*eventhandlers.OrderCreatedEvent)),
)

var setCreateOrderUseCase = wire.NewSet(
	setOrderRepositoryDependency,
	setOrderCreatedEvent,
	usecase.NewCreateOrderUseCase,
)

func NewCreateOrderUseCase(db *sql.DB, eventDispatcher events.EventDispatcher) *usecase.CreateOrderUseCase {
	wire.Build(setCreateOrderUseCase)
	return &usecase.CreateOrderUseCase{}
}

func NewOrderHandlerWeb(db *sql.DB, eventDispatcher events.EventDispatcher) *web.OrderHandlerWeb {
	wire.Build(
		setCreateOrderUseCase,
		web.NewOrderHandlerWeb,
	)
	return &web.OrderHandlerWeb{}
}
