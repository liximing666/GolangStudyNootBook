package main

type User struct {
	Name string
	Account *Account

}

func (u *User) ShowMoney() {
	u.Account.ShowMoney()
}

func (u *User) AddMoney(count int) {
	u.Account.AddMoney(count)
}

func (u *User) RemoveMoney(count int) {
	u.Account.RemoveMoney(count)
}



func NewUser(s string) *User {
	return &User{Name: s, Account: &Account{Price: 0}}
}
