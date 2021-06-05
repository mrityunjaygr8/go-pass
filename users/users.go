package users

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type User struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	ID        uint   `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func hashPassword(new_password string) string {
	return new_password
}

func (u *User) usernameExists(db *gorm.DB) bool {
	result := db.Where(&User{ID: u.ID}).Limit(1).Find(&u)
	if result.Error != nil {
		return false
	}

	if result.RowsAffected == 0 {
		return false
	}
	return true
}

func (u *User) CreateUser(db *gorm.DB) error {
	t := User{Username: u.Username}
	exists := t.usernameExists(db)
	if !exists {
		return fmt.Errorf(USER_EXISTS)
	}
	u.Password = hashPassword(u.Password)
	result := db.Create(&u)
	if result.Error != nil {
		fmt.Println("An error has occurred", result.Error)
	}
	return nil
}

func ListUsers(db *gorm.DB) ([]User, error) {
	var users []User
	if err := db.Find(&users).Error; err != nil {
		fmt.Println("An error has occurred", err)
		return nil, err
	}

	return users, nil
}

func (u *User) FetchUser(db *gorm.DB) error {
	result := db.Where(&User{ID: u.ID}).Limit(1).Find(&u)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf(USER_NOT_EXISTS)
	}
	return nil
}

func (u *User) UpdateUser(db *gorm.DB) error {
	if err := db.Model(&u).Update("password", hashPassword(u.Password)).Error; err != nil {
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
