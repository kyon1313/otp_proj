package model

import (
	"sample/database"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id       uint   `json:"id" gorm:"primaryKey"`
	Name     string `json:"name"`
	Lastname string `json:"lastname"`
	Mobile   string `json:"mobile"`
	Username string `json:"username"`
	Password string `json:"password"`
	Status   int    `json:"status"`
}

func (u *User) ValidateUsername() string {
	rows := database.DB.Find(&u, "username=?", u.Username).RowsAffected
	if rows > 0 {
		return "username already exist"
	}
	database.DB.Create(&u)
	return "user created"
}

func (user *User) HashPassword() error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(bytes)
	return err
}
