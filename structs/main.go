package main

import "fmt"

type emailInfo struct {
	emailAddress string
}

type phoneInfo struct {
	localCode   string
	stateCode   string
	phoneNumber string
}

type person struct {
	firstName string
	lastName  string
	email     emailInfo
	phone     phoneInfo
}

func main() {

	luiz := person{
		firstName: "Luiz",
		lastName:  "Silveira",
		email: emailInfo{
			emailAddress: "luizfelipinho@teste.com",
		},
		phone: phoneInfo{
			localCode:   "55",
			stateCode:   "11",
			phoneNumber: "9876543210",
		}}
	luiz.print()

	luiz.updateFirstName("LuizFelipe")
	luiz.updateLastName("Rosilveira")
	luiz.print()
}

func (pointerToPerson *person) updateFirstName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (pointerToPerson *person) updateLastName(newLastName string) {
	(*pointerToPerson).lastName = newLastName
}

func (p person) print() {
	fmt.Printf("%+v \n", p)
}
