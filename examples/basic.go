package main

import (
	jackalope "github.com/go-mesquite/Jackalope"
)

type User struct {
	ID    uint32 `json:"id" jackalope:"primary_key"`
	Email string `json:"email" jackalope:"unique,nil,default:0"`
	Name  string `json:"name"`
	Age   uint8  `json:"age" jackalope:"nil,default:0"`
}

func main() {
	// Run a function from the other file

	_, err := jackalope.NewDB("jackalope.db")
	if err != nil {
		panic(err)
	}
	//db.AddTable(&User{})

	//db.Create(&User{ID: 1, Name: "D42", Age: 100})

	/*
		err = db.Insert(user)
		if err != nil {
			fmt.Println("Error inserting user:", err)
			return
		}
	*/

	/*
		// Retrieve all users
		allUsers := db.GetAll()
		fmt.Println("All Users:")
		fmt.Println(allUsers)
	*/
}
