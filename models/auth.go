package models

import (
	"fmt"
	"log"

	"github.com/astaxie/beego/client/orm"
	"golang.org/x/crypto/bcrypt"
)

// HashPassword hashing
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	return string(bytes), err
}

// SaveUser to database
func SaveUser(user *User) string {
	// define package
	o := orm.NewOrm()
	// verify connection to database
	dr := o.Driver()
	fmt.Println(dr.Name() == "default")    // true
	fmt.Println(dr.Type() == orm.DRSqlite) // true
	fmt.Println("data password:", user.Password)
	// hash password
	hashedPassword, err := HashPassword(user.Password)

	if err != nil {
		log.Println(err)
	}

	fmt.Println("hashedPassword:", hashedPassword)
	// prepare data for insert to database
	u := User{Username: user.Username, Email: user.Email, Password: hashedPassword}
	id, err := o.Insert(&u)
	fmt.Printf("ID: %d, ERR: %v\n", id, err)

	msg := "user saved in database"

	return msg
}
