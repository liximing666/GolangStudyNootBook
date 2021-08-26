package main

func main() {
	a := NewUser("aaa")
	a.AddMoney(666)
	a.ShowMoney()


	b := NewUser("bbb")
	b.AddMoney(999)
	b.ShowMoney()

	c := NewUser("ccc")
	c.AddMoney(1000)
	c.ShowMoney()

	bank := Bank{Name: "cbc"}
	bank.Add(a)
	bank.Add(b)
	bank.Add(c)

	bank.ShowUsers()



}



