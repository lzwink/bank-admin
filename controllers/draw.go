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
	drawModel  = models.Draw{}
	awardModel = models.Award{}
)

// 创建抽奖记录
func (ctx *DrawController) CreateDraw() {
	userId := ctx.InterfaceToInt(ctx.GetSession("user_id"))
	// 查看抽奖条件
	signDays, err := signModel.CheckSignDays(userId)
	if err != nil {
		ctx.JsonEncode(101, "failed", nil, 0)
	}
	if signDays < 5 {
		ctx.JsonEncode(102, "failed", nil, 0)
	}
	// 生成可发放的奖品id
	allow := 0
	awardId := 0
	for allow != 1 {
		max := big.NewInt(6)
		r, _ := rand.Int(rand.Reader, max)
		awardId = int(r.Int64())
		surplus, err := awardModel.GetSurplusById(awardId)
		if err != nil {
			ctx.JsonEncode(101, "failed", nil, 0)
		}
		if surplus > 0 {
			allow = 1
		}
	}
	// 创建抽奖记录
	createTime := ctx.DateToStr(time.Now().Unix())
	err = drawModel.CreateDraw(userId, awardId, createTime)
	if err != nil {
		ctx.JsonEncode(101, "failed", nil, 0)
	}
	// 更新用户抽奖状态
	err = userModel.UpdateUserIsDraw(userId)
	if err != nil {
		ctx.JsonEncode(101, "failed", nil, 0)
	}
	// 减少奖品库存
	if awardId != 5 {
		err = awardModel.SubNumById(awardId)
		if err != nil {
			ctx.JsonEncode(101, "failed", nil, 0)
		}
	}
	// 返回具体奖品名称
	awardName, err := awardModel.GetNameById(awardId)
	if err != nil {
		ctx.JsonEncode(101, "failed", nil, 0)
	}
	ctx.JsonEncode(0, "success", awardName, 0)
}
