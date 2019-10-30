package test

import (
	"bank-admin/controllers"
	_ "bank-admin/routers"
	"testing"
)

func TestUsers(t *testing.T) {
	//log.Println((&models.Users{}).GetUserByCombatGroupId(1))
	(&controllers.UsersController{}).GetAllUsers()
}
