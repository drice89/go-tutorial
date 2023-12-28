package main

import "fmt"

type contactInfo struct {
	email string
	zip int
}
type person struct {
	firstName string
	lastName string
	contactInfo //or contactInfo contactInfo
}


func main() {
	jim := person{
		firstName: "Jim",
		lastName: "Bob",
		contactInfo: contactInfo{
			zip: 12345,
			email: "jimbob@gmail.com",
		},
	}

	jim.firstName = "Joe"
	jim.print() // {firstName:Joe lastName:Bob contactInfo:{email:jimbob@gmail.com zip:12345}}%
	
	
	// jim.updateName("joe")
	// jim.print()
}

func (p *person) updateName(n string) {
	(*p).firstName = n
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

	