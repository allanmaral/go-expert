package usecase

import (
	"github.com/allanmaral/go-expert/20-clean-arch/internal/entity"
	"github.com/allanmaral/go-expert/20-clean-arch/pkg/events"
)

type CreateOrderInput struct {
	ID    string  `json:"id"`
	Price float64 `json:"price"`
	Tax   float64 `json:"tax"`
}

type CreateOrderOutput struct {
	ID         string  `json:"id"`
	Price      float64 `json:"price"`
	Tax        float64 `json:"tax"`
	FinalPrice float64 `json:"finalPrice"`
}

type CreateOrderUseCase struct {
	OrderRepository entity.OrderRepository
	OrderCreated    events.Event
	EventDispatcher events.EventDispatcher
}

func NewCreateOrderUseCase(
	orderRespository entity.OrderRepository,
	orderCreated events.Event,
	eventDispatcher events.EventDispatcher,
) *CreateOrderUseCase {
	return &CreateOrderUseCase{
		OrderRepository: orderRespository,
		OrderCreated:    orderCreated,
		EventDispatcher: eventDispatcher,
	}
}

func (uc *CreateOrderUseCase) Execute(input CreateOrderInput) (CreateOrderOutput, error) {
	order, err := entity.NewOrder(input.ID, input.Price, input.Tax)
	if err != nil {
		return CreateOrderOutput{}, err
	}

	if err := order.CalculateFinalPrice(); err != nil {
		return CreateOrderOutput{}, err
	}

	if err := uc.OrderRepository.Save(order); err != nil {
		return CreateOrderOutput{}, err
	}

	output := CreateOrderOutput{
		ID:         order.ID,
		Price:      order.Price,
		Tax:        order.Tax,
		FinalPrice: order.FinalPrice,
	}

	uc.OrderCreated.SetPayload(output)
	uc.EventDispatcher.Dispatch(uc.OrderCreated)

	return output, nil
}
