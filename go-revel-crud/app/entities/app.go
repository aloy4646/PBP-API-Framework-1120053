package entities

import "cobaRevel/go-revel-crud/app/models"

type (
	UsersResponse struct {
		Success bool          `json:"success"`
		Status  int           `json:"status"`
		Message string        `json:"message"`
		Data    []models.User `json:"data"`
	}

	UserResponse struct {
		Success bool        `json:"success"`
		Status  int         `json:"status"`
		Message string      `json:"message"`
		Data    models.User `json:"data"`
	}

	Response struct {
		Success bool   `json:"success"`
		Status  int    `json:"status"`
		Message string `json:"message"`
	}
)
