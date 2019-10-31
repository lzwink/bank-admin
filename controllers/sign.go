package controllers

import (
	"bank-admin/models"
	"time"
)

type SignController struct {
	BaseController
}

var (
	signModel = models.Sign{}
)

func (ctx *SignController) CreateSign() {
	userId := ctx.InterfaceToInt(ctx.GetSession("user_id"))
	signTime := ctx.DateToStr(time.Now().Unix())
	err := signModel.CreateSign(userId, signTime)
	if err != nil {
		ctx.JsonEncode(101, "failed", nil, 0)
	}
	err = userModel.UpdateUserIsSign(userId)
	if err != nil {
		ctx.JsonEncode(101, "failed", nil, 0)
	}
	ctx.JsonEncode(0, "success", nil, 0)
}
