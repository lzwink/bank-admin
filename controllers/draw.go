package controllers

import (
	"bank-admin/models"
	"crypto/rand"
	"math/big"
	"time"
)

type DrawController struct {
	BaseController
}

var (
	drawModel = models.Draw{}
)

func (ctx *DrawController) CreateDraw() {
	userId := ctx.InterfaceToInt(ctx.GetSession("user_id"))
	// TODO: 抽奖逻辑暂未需要限制数量
	max := big.NewInt(8)
	r, _ := rand.Int(rand.Reader, max)
	awardId := int(r.Int64())
	createTime := ctx.DateToStr(time.Now().Unix())
	err := drawModel.CreateDraw(userId, awardId, createTime)
	if err != nil {
		ctx.JsonEncode(101, "failed", nil, 0)
	}
	err = userModel.UpdateUserIsDraw(userId)
	if err != nil {
		ctx.JsonEncode(101, "failed", nil, 0)
	}
	ctx.JsonEncode(0, "success", awardId, 0)
}
