package main

import "fmt"

type Bank struct {
	Name string
	Users []*User

}

func (bank *Bank) Add(user *User)  {
	bank.Users = append(bank.Users, user)
}
func (bank *Bank) ShowUsers () {
	for _, eachUser := range bank.Users {
		fmt.Println(eachUser.Name, eachUser.Account.Price)
	}
}

func NewBank(name string) *Bank {
	return &Bank{Name: name, Users: []*User{}}
}
