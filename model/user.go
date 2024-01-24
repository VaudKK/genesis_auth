package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID             primitive.ObjectID  `bson:"_id" json:"_id"`
	FirstName      string              `json:"firstName" validate:"required" binding:"max=30"`
	LastName       string              `json:"lastName" validate:"required" binding:"max=30"`
	Phone          string              `json:"phone" validate:"regexp=^[+|254|0][17]{1}[0-9]{8}$"`
	Email          string              `json:"email" validate:"required,email"`
	Password       string              `json:"password" validate:"required"`
	OrganizationId string              `json:"organizationId" validate:"required"`
	CreatedAt      primitive.Timestamp `json:"createdAt"`
	LastLogIn      uint64              `json:"lastLogIn"`
}
