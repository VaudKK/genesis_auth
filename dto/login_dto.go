package dto

type LogInDto struct {
	Identifier string `json:"identifier" validate:"required" binding:"min=8,max=50"`
	Password   string `json:"password" validate:"required" binding:"min=8,max=50"`
}
