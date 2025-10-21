package main

import "fmt"

type Client struct {
	Name    string
	Age     int
	Enabled bool
	// Address compondo
	UserAddress Address //criando
}

type Address struct {
	StreetName string
	Number     int
	City       string
	State      string
}

func (client Client) Disabled() {
	client.Enabled = false

	fmt.Printf("The client %s was disabled\n", client.Name)
}

func main() {
	user := Client{
		Name:    "John Doe",
		Age:     30,
		Enabled: true,
	}

	user.Enabled = false

	fmt.Println(user.Name)
	fmt.Println(user.UserAddress.City)
	user.Disabled()
}
