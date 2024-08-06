package main

import (
	"fmt"

	"example.com/biodata/biodata"
)

func main() {

	//Description for Project Biodata
	printProjectTitle("HI, WELCOME TO THE BIODATA WEBSITE!\n")
	printProjectTitle("PLEASE ENTER YOUR FIRST NAME, LAST NAME, EMAIL, ADDRESS, AND AGE \n")

	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	email := getUserData("Please enter your email: ")
	address := getUserData("Please enter your address: ")
	birthdate := getUserData("Please enter your birthdate: ")
	age := getUserAgeData("Please enter your age: ")

	var userBio *biodata.Bio
	userBio, err := biodata.New(firstName, lastName, email, address, birthdate, age)
	if err != nil {
		fmt.Println(err)
	}

	userBio.OutputDisplay()
	userBio.Save()

}

func getUserAgeData(promptTextAge string) uint {
	fmt.Println(promptTextAge)
	var valueAge uint
	fmt.Scan(&valueAge)
	return valueAge
}

func getUserData(promptText string) string {
	fmt.Println(promptText)
	var value string
	fmt.Scan(&value)
	return value
}

func printProjectTitle(value interface{}) {
	fmt.Println(value)
}
