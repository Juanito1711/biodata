package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"example.com/biodata/biodata"
)

func main() {

	//Description for Project Biodata
	printProjectTitle("HI, WELCOME TO THE BIODATA WEBSITE!\n")
	printProjectTitle("PLEASE ENTER YOUR FIRST NAME, LAST NAME, EMAIL, ADDRESS, AND AGE \n")

	firstName := getUserData("Please enter your first name: ", "").(string)
	lastName := getUserData("Please enter your last name: ", "").(string)
	email := getUserData("Please enter your email: ", "").(string)
	address := getUserData("Please enter your address: ", "").(string)
	birthdate := getUserData("Please enter your birthdate: ", "").(string)
	age := getUserData("Please enter your age: ", uint(0)).(uint)

	var userBio *biodata.Bio
	userBio, err := biodata.New(firstName, lastName, email, address, birthdate, age)
	if err != nil {
		fmt.Println(err)
	}

	userBio.OutputDisplay()
	userBio.Save()

}

func getUserData(promptText string, defaultValue interface{}) interface{} {
	fmt.Println(promptText)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	value := scanner.Text()

	switch defaultValue.(type) {
	case string:
		return value

	case uint:
		input, err := strconv.ParseUint(value, 10, 32)
		if err != nil {
			fmt.Println("Invalid input, expected to be an integer")
			return defaultValue
		}
		return uint(input)

	case float64:
		input, err := strconv.ParseFloat(value, 64)
		if err != nil {
			fmt.Println("Invalid input, expected to be an float")
			return defaultValue
		}
		return input

	case int:
		input, err := strconv.Atoi(value)
		if err != nil {
			fmt.Println("Invalid input, expected to be an integer")
			return defaultValue
		}
		return input

	default:
		return defaultValue
	}

}

func printProjectTitle(value interface{}) {
	fmt.Println(value)
}
