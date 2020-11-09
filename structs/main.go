package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func (pp *person) updateFirstName(firstName string) {
	(*pp).firstName = firstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func main() {
	emmanouilk := person{
		firstName: "Emmanouil",
		lastName:  "Katefidis",
		contactInfo: contactInfo{
			email:   "emmanouil.katefidis@citrix.com",
			zipCode: 26225,
		},
	}

	emmanouilk.updateFirstName("Manos")
	emmanouilk.print()
}
