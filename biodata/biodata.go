package biodata

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Bio struct {
	Firstname string    `json:"First name"`
	Lastname  string    `json:"Last name"`
	Email     string    `json:"Email"`
	Address   string    `json:"Address"`
	BirthDate string    `json:"Birthdate"`
	Age       uint      `json:"age"`
	CreatedAt time.Time `json:"Created At"`
}

type Admin struct {
	EmailAdmin    string
	PasswordAdmin string
	CreatedAt     time.Time
	Bio           Bio
}

func New(firstname string, lastname string, email string, address string, birthDate string, age uint) (*Bio, error) {
	if firstname == "" || lastname == "" || email == "" || address == "" || birthDate == "" || age == 0 {
		return &Bio{}, errors.New("Invalid input")
	}

	return &Bio{
		Firstname: firstname,
		Lastname:  lastname,
		Email:     email,
		Address:   address,
		BirthDate: birthDate,
		Age:       age,
		CreatedAt: time.Now(),
	}, nil
}

func NewAdmin(emailAdmin string, passwordAdmin string) (*Admin, error) {
	return &Admin{
		EmailAdmin:    " ",
		PasswordAdmin: " ",
		CreatedAt:     time.Now(),
	}, nil
}

func (bio *Bio) OutputDisplay() {
	fmt.Printf("First Name: %v\n Last Name: %v\n Email: %v\n Address: %v\n BirthDate: %v\n Age: %v\n", bio.Firstname, bio.Lastname, bio.Email, bio.Address, bio.BirthDate, bio.Age)
}

func (admin *Admin) Save() error {
	fileName := strings.ReplaceAll(admin.EmailAdmin, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"

	json, err := json.Marshal(admin)

	if err != nil {
		return err
	}

	return os.WriteFile(fileName, json, 0644)
}

func (bio *Bio) Save() error {
	fileName := strings.ReplaceAll(bio.Firstname, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"

	json, err := json.Marshal(bio)

	if err != nil {
		return err
	}

	return os.WriteFile(fileName, json, 0644)
}
