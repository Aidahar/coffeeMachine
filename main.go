package main

import (
	"fmt"
)

type Machine struct {
	water int
	milk  int
	beans int
	cups  int
	money int
}

func newMachine() *Machine {
	return &Machine{
		water: 400,
		milk:  540,
		beans: 120,
		cups:  9,
		money: 550,
	}
}

func (m *Machine) Info() {
	fmt.Println("The coffee machine has:")
	fmt.Printf("%d water\n", m.water)
	fmt.Printf("%d milk\n", m.milk)
	fmt.Printf("%d beans\n", m.beans)
	fmt.Printf("%d cups\n", m.cups)
	fmt.Printf("%d money\n", m.money)
}

func selectChoose(a string, m *Machine) {
	switch a {
	case "buy":
		buyCoffee(m)
	case "fill":
		fillMachine(m)
	case "take":
		takeCofe(m)
	case "remaining":
		remaining(m)
	}
}

func buyCoffee(m *Machine) {
	var choose int
	fmt.Println("What do you want to buy? 1 - espresso, 2 - latte, 3 - cappuccino:")
	fmt.Scan(&choose)
	switch choose {
	case 1:
		fmt.Println("Espresso")
		if m.water < 250 {
			fmt.Println("Sorry, not enough water!")
			return
		}
		if m.beans < 15 {
			fmt.Println("Sorry, not enough beans!")
			return
		}
		if m.cups < 1 {
			fmt.Println("Sorry, not enough cups!")
			return
		}
		fmt.Println("I have enough resources, making you a coffee!")
		m.water -= 250
		m.beans -= 16
		m.cups -= 1
		m.money += 4
		m.Info()
	case 2:
		fmt.Println("Latte")
		if m.water < 350 {
			fmt.Println("Sorry, not enough water!")
			return
		}
		if m.milk < 75 {
			fmt.Println("Sorry, not enough milk!")
			return
		}
		if m.beans < 20 {
			fmt.Println("Sorry, not enough beans!")
			return
		}
		if m.cups < 1 {
			fmt.Println("Sorry, not enough cups!")
			return
		}
		fmt.Println("I have enough resources, making you a coffee!")
		m.water -= 350
		m.milk -= 75
		m.beans -= 20
		m.cups -= 1
		m.money += 7
		m.Info()
	case 3:
		fmt.Println("Cappuccino")
		if m.water < 200 {
			fmt.Println("Sorry, not enough water!")
			return
		}
		if m.milk < 100 {
			fmt.Println("Sorry, not enough milk!")
			return
		}
		if m.beans < 12 {
			fmt.Println("Sorry, not enough beans!")
			return
		}
		if m.cups < 1 {
			fmt.Println("Sorry, not enough cups!")
			return
		}
		fmt.Println("I have enough resources, making you a coffee!")
		m.water -= 200
		m.milk -= 100
		m.beans -= 12
		m.cups -= 1
		m.money += 6
		m.Info()
	}
}

func fillMachine(m *Machine) {
	var water, milk, beans, cups int
	fmt.Println("Write how many ml of water you want to add:")
	fmt.Scan(&water)
	fmt.Println("Write how many ml of milk you want to add:")
	fmt.Scan(&milk)
	fmt.Println("Write how many grams of coffee beans you want to add:")
	fmt.Scan(&beans)
	fmt.Println("Write how many disposable coffee cups you want to add:")
	fmt.Scan(&cups)
	m.water += water
	m.milk += milk
	m.beans += beans
	m.cups += cups
}

func takeCofe(m *Machine) {
	fmt.Printf("I gave you %d\n", m.money)
	m.money = 0
}

func remaining(m *Machine) {
	m.Info()
}

func main() {
	var action string
	var m *Machine = newMachine()
	m.Info()

	for {
		fmt.Println("Write action (buy, fill, take, remaining, exit):")
		fmt.Scan(&action)
		selectChoose(action, m)
		if action == "exit" {
			return
		}
	}
}
