package main

import "fmt"

func main() {
	fmt.Println("welcome to method in golang")

	ashish := User{"Aashish", "tiwari@go.in", true, 19}
	fmt.Println(ashish)

	fmt.Printf("aashish details are %+v\n", ashish) //+ add they give more details
	fmt.Printf("Name is %v and Email is %v \n.", ashish.Name, ashish.Email)
	fmt.Println()
	ashish.getData()
	ashish.NewMail()
	fmt.Printf("Name is %v and Email is %v \n.", ashish.Name, ashish.Email)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) getData() {
	fmt.Println("is User is online", u.Status)

}

func (u User) NewMail() {
	u.Email = "tesla@gmail.com"
	fmt.Println("Email of this user is: ", u.Email)
}

/*
1. A method is a function attached to a type, such as a struct.
2. A method has a receiver that specifies the type it belongs to.
3. Methods can access and manipulate struct fields.
4. Go passes values by copy. A value receiver gets a copy of the struct.
5. A pointer receiver can modify the original struct.
6. If the method name starts with a CAPITAL letter, it is exported
   and can be accessed from another package.

7. If the method name starts with a lowercase letter, it is unexported
   and can only be accessed within its own package.

8. Methods can have parameters and return values.
9. Go does not have classes, but methods can be defined on user-defined types.
10. The receiver can be a value receiver or a pointer receiver.
*/
