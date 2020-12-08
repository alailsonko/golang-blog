package auth

import (
	"fmt"
	"log"

	"github.com/astaxie/beego/client/orm"
	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	return string(bytes), err
}

// User types
type User struct {
	Username string
	Email    string
	Password string
}

// SaveUser to database
func (user *User) SaveUser() string {

	o := orm.NewOrm()

	dr := o.Driver()
	fmt.Println(dr.Name() == "default")    // true
	fmt.Println(dr.Type() == orm.DRSqlite) // true
	fmt.Println("data password:", user.Password)
	hashedPassword, err := hashPassword(user.Password)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("hashedPassword:", hashedPassword)
	msg := "working"

	return msg
}
