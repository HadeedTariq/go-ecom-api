package types

type RegisterUserPayload struct {
	FirstName string `json:"firstName" validate:"required,trim,min=2,max=50,alpha"`
	LastName  string `json:"lastName" validate:"required,trim,min=2,max=50,alpha"`
	Email     string `json:"email" validate:"required,trim,email,max=255"`
	Password  string `json:"password" validate:"required,min=8,max=72"`
}
