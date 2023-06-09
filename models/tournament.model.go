package models

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type Tournament struct {
	ID              uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name            string
	AddressLineOne  string
	AddressLineTwo  string
	AddressPostCode string
	AddressCity     string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

var validate = validator.New()

type ErrorResponse struct {
	Field string `json:"field"`
	Tag   string `json:"tag"`
	Value string `json:"value,omitempty"`
}

func ValidateStruct[T any](payload T) []*ErrorResponse {
	var errors []*ErrorResponse
	err := validate.Struct(payload)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.Field = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}

type CreateTournamentSchema struct {
	Name            string `json:"name" validate:"required"`
	AddressLineOne  string `json:"address_line_one" validate:"required"`
	AddressLineTwo  string `json:"address_line_two,omitempty"`
	AddressCity     string `json:"address_city" validate:"required"`
	AddressPostCode string `json:"address_post_code" validate:"required"`
}

type UpdateTournamentSchema struct {
	Name            string `json:"name" validate:"required"`
	AddressLineOne  string `json:"address_line_one" validate:"required"`
	AddressLineTwo  string `json:"address_line_two,omitempty"`
	AddressCity     string `json:"address_city" validate:"required"`
	AddressPostCode string `json:"address_post_code" validate:"required"`
}
