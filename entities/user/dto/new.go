package dto

type NewUser struct {
	Name     string `json:"name" validate:"required"`
	Lastname string `json:"lastname" validate:"required"`
	UType    string `json:"type" validate:"required"`
	Document string `json:"document" validate:"required,min=11,max=14"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}
