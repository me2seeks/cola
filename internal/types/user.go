package types

import "github.com/me2seeks/cola/internal/models"

type LoginReq struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type LoginResp struct{}

type RegisterReq struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	Sex      int8   `json:"sex" binding:"required,oneof=0 1 2"` // Assuming 0, 1, 2 are valid values
	Avatar   string `json:"avatar" binding:"omitempty,url"`
	Info     string `json:"info" binding:"omitempty,max=500"`
}

type RegisterResp struct{}

type UpdateUserReq struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Sex      int8   `json:"sex"`
	Avatar   string `json:"avatar"`
	Info     string `json:"info"`
}

type UpdateUserResp struct{}

type ListUserByPageReq struct {
	PageNum  int `json:"pageNum" binding:"required,min=1" default:"1"`
	PageSize int `json:"pageSize" binding:"required,min=1" default:"10"`
}

type ListUserByPageResp struct {
	Users []*models.User `json:"users"`
}
