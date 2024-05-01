package entity

type OrderRepository interface {
	Save(*Order) error
}
