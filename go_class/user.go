package main

type User struct {
	firstName string
	lastName  string
	id        int
}

func (u User) fullName() string {
	return u.firstName + " " + u.lastName
}

func main() {
	user := User{firstName: "John", lastName: "Doe", id: 1}
	println(user.fullName())
}
