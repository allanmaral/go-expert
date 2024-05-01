package entity

import "errors"

type Order struct {
	ID         string
	Price      float64
	Tax        float64
	FinalPrice float64
}

var (
	ErrInvalidID    = errors.New("invalid id")
	ErrInvalidPrice = errors.New("invalid price")
	ErrInvalidTax   = errors.New("invalid tax")
)

func NewOrder(id string, price float64, tax float64) (*Order, error) {
	order := &Order{
		ID:    id,
		Price: price,
		Tax:   tax,
	}

	if err := order.IsValid(); err != nil {
		return nil, err
	}

	return order, nil
}

func (o *Order) IsValid() error {
	if o.ID == "" {
		return ErrInvalidID
	}
	if o.Price <= 0 {
		return ErrInvalidPrice
	}
	if o.Tax <= 0 {
		return ErrInvalidTax
	}
	return nil
}

func (o *Order) CalculateFinalPrice() error {
	o.FinalPrice = o.Price + o.Tax
	if err := o.IsValid(); err != nil {
		return err
	}
	return nil
}
