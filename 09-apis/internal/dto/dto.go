package dto

type CreateProductInput struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type UpdateProductInput struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type ProductOutput struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type CreateUserInput struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthOutput struct {
	AccessToken string `json:"access_token"`
}

type ErrorOutput struct {
	Message string `json:"message"`
}
