package myserver

//import of used packages
import (
	"fmt"
	"strconv"
)

// All thoses Functions insteract with a global variable.
// But what if we want to start 2 servers ?
// Each servers should have a different user list and a secure way to treat them.

//Function to check if a user exist
func ChkUser(id string) bool {
	var Chk bool = false
	ConvertID, err := strconv.Atoi(id)
	if err != nil {
		panic("Error()")
	}

	for i, a := range users {
		if users == nil {
			panic(i)
		}
		if a.ID == ConvertID {
			Chk = true
		} else {
			Chk = false
		}
	}
	return Chk
}

//function du remove a user
func remove(i int) {
	users[len(users)-1], users[i] = users[i], users[len(users)-1]
	users = users[:len(users)-1]
	MkOrder()
}

//A function to be sure that a user was created
func printSlice(s []User) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

//Function to order users by ID
func MkOrder() {
	var User_Mem User
	var Ctl int = 0

	for Ctl < len(users) {
		var Mem int = Ctl
		var h int = Ctl
		for h, a := range users {
			if users[h].ID > users[Mem].ID {
				User_Mem = users[Mem]
				users[Mem] = users[h]
				users[h] = User_Mem
				if users == nil {
					panic(a)
				}
			}
		}
		if users == nil {
			panic(h)
		}
		Ctl = Ctl + 1
	}
}
