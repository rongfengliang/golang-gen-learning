package main

import "fmt"

// +gen set
type MyUser struct {
	Name string
	Age  int
}

func (u MyUser) String() string {
	return fmt.Sprintf("this is myuser print: %s , %d", u.Name, u.Age)
}
