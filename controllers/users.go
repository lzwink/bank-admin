package controllers

import "bank-admin/models"

type UsersController struct {
	BaseController
}

var (
	userModel = models.Users{}
)

func (ctx *UsersController) Get() {
	ctx.JsonEncode(0, "success", "hello world", 0)
}

func (ctx *UsersController) GetAllUsers() {
	res, err := userModel.GetAllUsers()
	if err != nil {
		ctx.JsonEncode(101, "failed", nil, 0)
	}
	ctx.JsonEncode(0, "success", res, len(res))
}

func (ctx *UsersController) GetUserById() {
	id, _ := ctx.GetInt("id")
	res, err := userModel.GetUserById(id)
	if err != nil {
		ctx.JsonEncode(101, "failed", nil, 0)
	}
	ctx.JsonEncode(0, "success", res, 1)
}

func (ctx *UsersController) UpdateUserPwd() {

}
