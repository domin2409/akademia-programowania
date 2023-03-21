package main

import (
	// "fmt"
	// "pjatk_project/token"

	"log"

	"github.com/google/uuid"
	"github.com/grupawp/appdispatcher"
)

func main() {
	// uniqueToken := token.Generate(16)
	// fmt.Printf("Twój unikalny token to: %s", uniqueToken)

	status, err := appdispatcher.Submit(Student{"Dominik", "Dudzik", uuid.New()})
	log.Print(err)
	log.Print(status)
}

type Student struct {
	FirstName     string
	LastName      string
	applicationID uuid.UUID
}

func (s Student) ApplicationID() string {
	return s.applicationID.String()
	// return s.applicationID
}

func (s Student) FullName() string {
	return s.FirstName + " " + s.LastName
}
