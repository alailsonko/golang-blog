package models

import (
	"github.com/astaxie/beego/client/orm"
)

// User interface
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func init() {
	orm.RegisterModel(new(User))
}
