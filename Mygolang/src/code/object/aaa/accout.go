package main

import "fmt"

type Account struct {
	Price int
}



func (account *Account) AddMoney(count int)  {
	account.Price += count
}
func (account *Account) RemoveMoney(count int) {
	if account.Price >= count {
		account.Price -= count
	}else {
		fmt.Println("余额不足")
	}
}

func (account *Account) ShowMoney() {
	fmt.Println(account.Price)
}