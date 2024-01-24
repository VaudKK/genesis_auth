package dto

type UserDto struct {
	FirstName      string `json:"firstName" validate:"required" binding:"max=30"`
	LastName       string `json:"lastName" validate:"required" binding:"max=30"`
	Phone          string `json:"phone" validate:"required"`
	Email          string `json:"email" validate:"required,email"`
	Password       string `json:"password" validate:"required"`
	OrganizationId string `json:"organizationId" validate:"required"`
	LastLogIn      uint64 `json:"lastLogIn"`
}
