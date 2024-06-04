package dto

type NewUser struct {
	Name     string `json:"name"`
	Lastname string `json:"lastname"`
	UType    string `json:"type"`
	Document string `json:"document"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
