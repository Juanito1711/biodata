package main

import (
	"fmt"
)

func main() {

	printProjectTitle("HI, WELCOME TO THE BIODATA WEBSITE!\n")
	printProjectTitle("PLEASE ENTER YOUR FIRST NAME, LAST NAME, EMAIL, ADDRESS, AND AGE \n")

}

func printProjectTitle(value interface{}) {
	fmt.Println(value)
}
