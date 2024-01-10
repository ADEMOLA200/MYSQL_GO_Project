package main

import (
	_"database/sql"
	_ "fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// type User struct {
// 	gorm.Model
// 	FirstName sql.NullString `gorm:"type:VARCHAR(30); null"`
//     LastName  sql.NullString `gorm:"size:100; default:'Smith'"`
//     Email     sql.NullString `gorm:"unique; not null"`
// 	Address		Address		 `gorm:"foreignKey:UserId"`
// }



type User struct {
	gorm.Model
	FirstName 	string 		`gorm:"type:VARCHAR(30); null"`
    LastName  	string 		`gorm:"size:100; default:'Smith'"`
    Email     	string 		`gorm:"unique; not null"`
	// Address	  	Address		`gorm:"foreignKey:UserId"`
	Books		[]Book		`gorm:"many2many:user_id"`
}

// type Address struct {
// 	gorm.Model
// 	UserId			uint
// 	NameOfAddress	string
// }

type Book struct {
	ID			uint
	Title		string
}



func main() {
	// connect to the database
	db, err := gorm.Open(mysql.Open("root:1234@/go-react-application"), &gorm.Config{})

	// check if err is not nil
	if err != nil {
		panic("cannot connect to database")
	}

	// db.Migrator().DropTable(&User{}, &Book{})
	// db.Migrator().CreateTable(&User{}, &Book{})


	// user := User {
	// 	Email: sql.NullString{
	// 		String: "johndoe@gmail.com",
	// 		Valid: true,
	// 	},
	// }


	db.AutoMigrate(&User{}, &Book{})

	user := User {
		FirstName: "a",
		LastName: "b",
		Email: "a@b.com",
		// Address: Address{
		// 	NameOfAddress: "Saint main str, 39",
		// },
		Books: []Book{
			{
				Title: "mysql_go",
			},
		},
	}
	db.Create(&user)


	// //passing records to our database
	// users := []User {
	// 	{
	// 		FirstName: "John",
	// 		LastName: "Doe",
	// 		Email: "johndoe@gmail.com",
	// 	},
	// 	{
	// 		FirstName: "Harley",
	// 		LastName: "Davids",
	// 		Email: "harleydavis@outlook.com",
	// 	},
	// 	{
	// 		FirstName: "Ryan",
	// 		LastName: "Stocker",
	// 		Email: "ryanstocker@hotmail.com",
	// 	},
	// }

	// for _, user := range users {
	// 	db.Create(&user)
	// }
}