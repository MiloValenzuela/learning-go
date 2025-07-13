package main

import (
	"log"
	"time"

	"github.com/MiloValenzuela/learning-go/helpers"
)

var s = "seven"

// var firstName string
// var lastName string
// var phoneNumber string
// var age int
// var birthDate time.Time

type User struct {
	FirstName   string
	LastName    string
	PhoneNumber string
	Age         int
	BirthDate   time.Time
}

func main() {
	user := User{
		FirstName:   "Trevor",
		LastName:    "Sawler",
		PhoneNumber: "1 555-555-1212",
	}

	log.Println(user.FirstName, user.LastName, "Birthdate:", user.BirthDate)

	changeUsingPointer(&user.LastName)

	var myVar helpers.SomeType

	myVar.TypeName = "Some Name"
	log.Println(myVar.TypeName)

}

func changeUsingPointer(s *string) {
	newValue := "red"
	*s = newValue
}
