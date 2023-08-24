package users

import "fmt"

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func NewUser() *User {
	return &User{}
}

func (u *User) ScanUser() {
	var name, password string
	fmt.Println("Your Name")
	fmt.Scan(&name)
	fmt.Println("Your Password")
	fmt.Scan(&password)
	u.Name = name
	u.Password = password
}
