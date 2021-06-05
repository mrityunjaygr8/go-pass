package users

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
}

func hashPassword(new_password string) string {
	return new_password
}

func CreateUser(username, password string, db *gorm.DB) (User, error) {
	var test User
	if err := db.Where(&User{Username: username}).First(&test).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Println("An error has occurred", err)
			return User{}, err
		} else {
			user := User{Username: username, Password: hashPassword(password)}
			result := db.Create(&user)
			if result.Error != nil {
				fmt.Println("An error has occurred", result.Error)
			}
			fmt.Println(user.ID, user.Username)
			return user, nil
		}
	}
	return User{}, fmt.Errorf("user already exists")
}

func ListUsers(db *gorm.DB) ([]User, error) {
	var users []User
	if err := db.Find(&users).Error; err != nil {
		fmt.Println("An error has occurred", err)
		return nil, err
	}

	return users, nil
}

func FetchUser(username string, db *gorm.DB) (User, error) {
	var user User
	if err := db.Where(&User{Username: username}).Limit(1).Find(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return User{}, fmt.Errorf("user does not exist")
		} else {
			fmt.Println("An error has occurred", err)
			return User{}, err
		}
	}

	return user, nil
}

func (u *User) UpdateUser(new_password string, db *gorm.DB) error {
	if err := db.Model(&u).Update("password", hashPassword(new_password)).Error; err != nil {
		return err
	}
	return nil
}

func (u *User) DeleteUser(db *gorm.DB) error {
	if err := db.Delete(&u).Error; err != nil {
		return err
	}
	return nil
}
