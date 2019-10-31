package test

import (
	"bank-admin/models"
	_ "bank-admin/routers"
	"log"
	"testing"
)

func TestUsers(t *testing.T) {
	log.Println((&models.Users{}).GetUserById(2))
}
