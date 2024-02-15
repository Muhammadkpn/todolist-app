package model

type (
	CreateUserRequest struct {
		FullName string `json:"fullName" validate:"required"`
		Address  string `json:"address" validate:"required"`
	}

	GetUserDetailResponse struct {
		ID       uint64
		FullName string
		Address  string
	}
)
