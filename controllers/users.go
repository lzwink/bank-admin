package controllers

import (
	"bank-admin/models"
	"strconv"
)

type UsersController struct {
	BaseController
}

var (
	userModel      = models.Users{}
	scoreModel     = models.Score{}
	userGroupModel = models.UsersGroup{}
)

func (ctx *UsersController) Get() {
	ctx.TplName = "index.html"
}

// 获取全部用户信息
func (ctx *UsersController) GetAllUsers() {
	res, err := userModel.GetAllUsers()
	if err != nil {
		ctx.JsonEncode(101, "failed", nil, 0)
	}
	ctx.JsonEncode(0, "success", res, len(res))
}

// 根据用户id获取用户信息
func (ctx *UsersController) GetUserById() {
	id, _ := ctx.GetInt("id")
	res, err := userModel.GetUserById(id)
	if err != nil {
		ctx.JsonEncode(101, "failed", nil, 0)
	}
	ctx.JsonEncode(0, "success", res, 1)
}

// 获取登录用户信息
func (ctx *UsersController) GetLoginUser() {
	id := ctx.InterfaceToInt(ctx.GetSession("user_id"))
	user, err := userModel.GetUserById(id)
	if err != nil {
		ctx.JsonEncode(100, "failed", nil, 0)
	}
	ctx.JsonEncode(0, "success", user, 0)
}

// 根据用户名称获取用户信息
func (ctx *UsersController) GetUserByRealName() {
	realName := ctx.XssFilter(ctx.GetString("real_name"))
	res, err := userModel.GetUserByRealName(realName)
	if err != nil {
		ctx.JsonEncode(101, "failed", nil, 0)
	}
	ctx.JsonEncode(0, "success", res, len(res))
}

// 更新用户密码
func (ctx *UsersController) UpdateUserPwd() {
	oldPwd := ctx.XssFilter(ctx.GetString("oldPwd"))
	newPwd := ctx.XssFilter(ctx.GetString("newPwd"))
	userId := ctx.InterfaceToInt(ctx.GetSession("user_id"))
	err := userModel.UpdateUserPwd(userId, oldPwd, newPwd)
	if err != nil {
		ctx.JsonEncode(101, "failed", nil, 0)
	}
	ctx.JsonEncode(0, "success", nil, 0)
}

// 根据id查询对手列表
func (ctx *UsersController) GetOpponentList() {
	id := ctx.InterfaceToInt(ctx.GetSession("user_id"))
	userInfo, err := userModel.GetUserById(id)
	if err != nil {
		ctx.JsonEncode(101, "failed", nil, 0)
	}
	oppList := make([]models.Users, 0)
	userList, err := userModel.GetUserByGroupId(userInfo.GroupId)
	if err != nil {
		ctx.JsonEncode(101, "failed", nil, 0)
	}
	for _, v := range userList {
		if v.Id != id {
			oppList = append(oppList, v)
		}
	}
	ctx.JsonEncode(0, "success", oppList, len(oppList))
}

// 根据id选择对手绑定
func (ctx *UsersController) ChooseOpponent() {
	id := ctx.InterfaceToInt(ctx.GetSession("user_id"))
	oppId, _ := ctx.GetInt("oppId")
	err := userModel.UpdateUserOpp(id, oppId)
	if err != nil {
		ctx.JsonEncode(101, "failed", nil, 0)
	}
	ctx.JsonEncode(0, "success", nil, 0)
}

// 参赛人员创建比赛目标
func (ctx *UsersController) AddTargetScore() {
	userName := ctx.InterfaceToStr(ctx.GetSession("user_name"))
	eOneStr := ctx.XssFilter(ctx.GetString("eventOne"))
	eOneFloat, _ := strconv.ParseFloat(eOneStr, 64)
	eTwoStr := ctx.XssFilter(ctx.GetString("eventTwo"))
	eTwoFloat, _ := strconv.ParseFloat(eTwoStr, 64)
	eThreeStr := ctx.XssFilter(ctx.GetString("eventThree"))
	eThreeFloat, _ := strconv.ParseFloat(eThreeStr, 64)
	eFourStr := ctx.XssFilter(ctx.GetString("eventFour"))
	eFourFloat, _ := strconv.ParseFloat(eFourStr, 64)
	err := scoreModel.AddTargetScore(userName, eOneFloat, eTwoFloat, eThreeFloat, eFourFloat)
	if err != nil {
		ctx.JsonEncode(101, "failed", nil, 0)
	}
	ctx.JsonEncode(0, "success", nil, 0)
}

// 管理员更新实际得分
func (ctx *UsersController) AddRealScore() {
	// 验证身份权限
	id := ctx.InterfaceToInt(ctx.GetSession("user_id"))
	user, err := userModel.GetUserById(id)
	if err != nil {
		ctx.JsonEncode(100, "failed", nil, 0)
	}
	if user.AuthGroupId != 1 {
		ctx.JsonEncode(100, "failed", nil, 0)
	}
	// 读取Excel表数据
	// 暂时通过get方式测试
	userName := ctx.XssFilter(ctx.GetString("userName"))
	eOneStr := ctx.XssFilter(ctx.GetString("eventOne"))
	eOneFloat, _ := strconv.ParseFloat(eOneStr, 64)
	eTwoStr := ctx.XssFilter(ctx.GetString("eventTwo"))
	eTwoFloat, _ := strconv.ParseFloat(eTwoStr, 64)
	eThreeStr := ctx.XssFilter(ctx.GetString("eventThree"))
	eThreeFloat, _ := strconv.ParseFloat(eThreeStr, 64)
	eFourStr := ctx.XssFilter(ctx.GetString("eventFour"))
	eFourFloat, _ := strconv.ParseFloat(eFourStr, 64)
	err = scoreModel.AddRealScore(userName, eOneFloat, eTwoFloat, eThreeFloat, eFourFloat)
	if err != nil {
		ctx.JsonEncode(101, "failed", nil, 0)
	}
	ctx.JsonEncode(0, "success", nil, 0)
}

// 获取用户分数信息
func (ctx *UsersController) GetUserScore() {
	userName := ctx.InterfaceToStr(ctx.GetSession("user_name"))
	score, err := scoreModel.GetScoreByUserName(userName)
	if err != nil {
		ctx.JsonEncode(101, "failed", nil, 0)
	}
	ctx.JsonEncode(0, "success", score, 0)
}

// 获取对手分数信息
func (ctx *UsersController) GetOppScore() {
	id := ctx.InterfaceToInt(ctx.GetSession("user_id"))
	user, err := userModel.GetUserById(id)
	if err != nil {
		ctx.JsonEncode(101, "failed", nil, 0)
	}
	opp, err := userModel.GetUserById(user.OpponentId)
	oppScore, err := scoreModel.GetScoreByUserName(opp.UserName)
	if err != nil {
		ctx.JsonEncode(101, "failed", nil, 0)
	}
	ctx.JsonEncode(0, "success", oppScore, 0)
}

// 根据用户获取对手信息
func (ctx *UsersController) GetOppInfo() {
	id := ctx.InterfaceToInt(ctx.GetSession("user_id"))
	user, err := userModel.GetUserById(id)
	if err != nil {
		ctx.JsonEncode(101, "failed", nil, 0)
	}
	opp, err := userModel.GetUserById(user.OpponentId)
	if err != nil {
		ctx.JsonEncode(101, "failed", nil, 0)
	}
	ctx.JsonEncode(0, "success", opp, 0)
}

// 获取用户分组列表
func (ctx *UsersController) GetAllUsersGroup() {
	res, err := userGroupModel.GetAllUsersGroup()
	if err != nil {
		ctx.JsonEncode(101, "failed", nil, 0)
	}
	ctx.JsonEncode(0, "success", res, len(res))
}

// 根据分组编号获取用户列表
func (ctx *UsersController) GetUsersByGroupId() {
	groupId, _ := ctx.GetInt("groupId")
	res, err := userModel.GetUserByGroupId(groupId)
	if err != nil {
		ctx.JsonEncode(101, "failed", nil, 0)
	}
	ctx.JsonEncode(0, "success", res, 0)
}

// 根据id获取信息
func (ctx *UsersController) GetAllInfoById() {
	id, _ := ctx.GetInt("id")
	res := make(map[string]interface{}, 0)
	user, err := userModel.GetUserById(id)
	if err != nil {
		ctx.JsonEncode(101, "failed", nil, 0)
	}
	res["user"] = user
	opp, err := userModel.GetUserById(user.OpponentId)
	if err != nil {
		ctx.JsonEncode(101, "failed", nil, 0)
	}
	res["opp"] = opp
	uScore, err := scoreModel.GetScoreByUserName(user.UserName)
	if err != nil {
		ctx.JsonEncode(101, "failed", nil, 0)
	}
	res["uScore"] = uScore
	oScore, err := scoreModel.GetScoreByUserName(opp.UserName)
	if err != nil {
		ctx.JsonEncode(101, "failed", nil, 0)
	}
	res["oScore"] = oScore
	ctx.JsonEncode(0, "success", res, 0)
}
