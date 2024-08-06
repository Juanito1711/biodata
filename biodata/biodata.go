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
	Firstname string
	Lastname  string
	Email     string
	Address   string
	BirthDate string
	Age       uint
	CreatedAt time.Time
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

func (bio *Bio) OutputDisplay() {
	fmt.Printf("First Name: %v\n Last Name: %v\n Email: %v\n Address: %v\n BirthDate: %v\n Age: %v\n", bio.Firstname, bio.Lastname, bio.Email, bio.Address, bio.BirthDate, bio.Age)
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
