package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo // or we can simple use contactInfo
}

func main() {
	//p := person{firstName: "Hem", lastName: "Kant"}
	//fmt.Println(p)

	var p1 person
	p1.firstName = "Hem"
	p1.lastName = "Kant"

	//fmt.Println(p1)
	//fmt.Printf("%+v", p1)

	p2 := person{
		firstName: "Hem",
		lastName:  "Kant",
		contact: contactInfo{
			email:   "hemkant.india@gmail.com",
			zipCode: 110011,
		},
	}
	//fmt.Printf("%+v", p2)

	// update fname based on pointer
	p2pointer := &p2
	p2pointer.updateName("hem")
	//end
	//p2.updateName("H")

	p2.print()
}

//func (p person) updateName(newFirstName string) {
//	p.firstName = newFirstName
//}

// update fname based on pointer *pointerTpPerson (memoey address where hem exist) give me the value of this memoery address
func (pointerTpPerson *person) updateName(newFirstName string) {
	(*pointerTpPerson).firstName = newFirstName
}
func (p person) print() {
	fmt.Println("%+v", p)
}
