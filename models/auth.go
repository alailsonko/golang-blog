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

	o := orm.NewOrm()

	dr := o.Driver()
	fmt.Println(dr.Name() == "default")    // true
	fmt.Println(dr.Type() == orm.DRSqlite) // true
	fmt.Println("data password:", user.Password)
	hashedPassword, err := HashPassword(user.Password)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("hashedPassword:", hashedPassword)
	u := User{Username: user.Username, Email: user.Email, Password: hashedPassword}
	// checkUniqueUsername := User{Username: "alailson"}
	// checkUniqueEmail := User{Email: "aalailson3@gmail.com"}
	var users []*User
	numUsername, errUsername := o.QueryTable("user").Filter("username", user.Username).All(&users)
	numEmail, errEmail := o.QueryTable("user").Filter("email", user.Email).All(&users)
	fmt.Println("queryTable username init")
	fmt.Printf("Returned Rows Num: %v, %s\n", numUsername, errUsername)
	fmt.Println("queryTable username end")
	fmt.Println("queryTable email init")
	fmt.Printf("Returned Rows Num: %v, %s\n", numEmail, errEmail)
	fmt.Println("queryTable email end")
	if numEmail == numUsername {
		fmt.Println("email", numEmail, "username", numUsername)
		id, err := o.Insert(&u)
		fmt.Printf("ID: %d, ERR: %v\n", id, err)
		fmt.Println("email", numEmail, "username", numUsername)

	} else {
		log.Println("user already exist")
	}

	msg := "working"

	return msg
}
