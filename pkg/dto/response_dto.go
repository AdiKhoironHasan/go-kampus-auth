package dto

type ResponseDTO struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type UserResponse struct {
	ID    int64  `json:"id_user"`
	Name  string `json:"name"`
	Token string `json:"token"`
}
