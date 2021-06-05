package main

import (
	"fmt"
	"log"

	"github.com/mrityunjaygr8/go-pass/app"
	"github.com/mrityunjaygr8/go-pass/users"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	fmt.Println("yo")

	app := app.App{}
	dsn := "host=localhost user=mgr8 password=dr0w.Ssap dbname=pass"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("An error has occurred")
	}

	app.DB = db
	app.Initialize()

	all_user, err := users.ListUsers(db)
	if err != nil {
		fmt.Println("An error has occurred", err)
	}

	fmt.Println(len(all_user))
	for _, u := range all_user {
		fmt.Println(u)
	}

	n_user, err := users.CreateUser("mgr8", "pass", app.DB)
	if err != nil {
		fmt.Println("An error has occurred", err)
	}
	if n_user != (users.User{}) {
		fmt.Println(n_user)
	}

	user, err := users.FetchUser("mgr", app.DB)
	if err != nil {
		fmt.Println("An error has occurred", err)
	}
	fmt.Println(user.Username, user.Password, user.ID)

	err = user.UpdateUser("newpass111", app.DB)
	if err != nil {
		fmt.Println("An error has occurred", err)
	}

	fmt.Println(user.Password)

	err = user.DeleteUser(db)
	if err != nil {
		fmt.Println("An error has occurred", err)
	}

	all_user_again, err := users.ListUsers(db)
	if err != nil {
		fmt.Println("An error has occurred", err)
	}

	fmt.Println(len(all_user_again))
	for _, u := range all_user_again {
		fmt.Println(u)
	}
}
