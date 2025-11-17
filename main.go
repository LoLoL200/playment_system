package main

import (
	"fmt"
	"os"
)

// func int input
func getIntInput(promt string) int {
	var input int
	fmt.Print(promt)
	fmt.Scanln(&input)
	return input
}

// Interfice playment
type Payment interface {
	Commission() float64
	Total() float64
	Validate() bool
	GetName() string
	GetBalance() float64
	SetBalance(float64)
	GetFeePercent() float64
}

type paymentBase struct {
	Name       string
	Balance    float64
	FeePercent float64
}

func (p *paymentBase) GetName() string        { return p.Name }
func (p *paymentBase) GetBalance() float64    { return p.Balance }
func (p *paymentBase) SetBalance(b float64)   { p.Balance = b }
func (p *paymentBase) GetFeePercent() float64 { return p.FeePercent }

// Structur
type KreditCard struct{ paymentBase }
type PayPal struct{ paymentBase }
type Cach struct{ paymentBase }
type BankTransfer struct{ paymentBase }

// Func for commission accounts
func (p *paymentBase) Commission() float64 {
	return p.Balance * p.FeePercent / 100
}

// Func for payments with commissions
func (p *paymentBase) Total() float64 {
	return p.Balance + p.Commission()
}

// Func for verification of funds
func (p *paymentBase) Validate() bool {
	if p.Balance >= 10.00 && p.Balance <= 50000.00 {
		fmt.Println("✅ Платіж успішно оброблено!")
		return true
	}

	fmt.Println("📛 Сума платежу повинна бути від 10 до 50 000 грн!!!!!")
	os.Exit(0)
	return false
}

// Func for information output
func processPayment(p Payment) {
	fmt.Println("------------------------------")
	fmt.Printf("Ведіть вашу сумму на балансі: ")
	var balance float64
	fmt.Scanln(&balance)
	p.SetBalance(balance)

	fmt.Println("------------------------------")
	fmt.Printf("Обробляємо платіж на суму %.2f грн...\n", p.GetBalance())
	fmt.Println("==================================")
	fmt.Printf("💳 %s\n", p.GetName())
	fmt.Printf("%t\n", p.Validate())
	fmt.Printf("💰 Сума: %.2f грн\n", p.GetBalance())
	fmt.Printf("💸 Комісія: %.2f грн (%.2f%%)\n", p.Commission(), p.GetFeePercent())
	fmt.Printf("📊 До списання: %.2f грн\n", p.Total())
	fmt.Println("Дякуємо за покупку! )))))")
	fmt.Println("==================================")
}

// Func Playment system
func systemPlayment() {
	var isAktive bool
	isAktive = true
	for {
		if isAktive == true {
			fmt.Println("========== Система платежів =============")
			fmt.Printf("Доступні методи оплати:\n1. Кредитна картка\n2. PayPal\n3. Готівка\n4. Банківський переказ\n5. Вихід\n")
			fmt.Println("==================================")
			choice := getIntInput("Оберіть метод оплати (1-5): ")
			fmt.Println("==================================")

			switch choice {
			case 1:
				processPayment(&KreditCard{paymentBase{"Кредитна картка", 0, 1.5}})
			case 2:
				processPayment(&PayPal{paymentBase{"PayPal", 0, 3.5}})
			case 3:
				processPayment(&Cach{paymentBase{"Готівка", 0, 0}})
			case 4:
				processPayment(&BankTransfer{paymentBase{"Банківський переказ", 0, 2.0}})
			case 5:
				fmt.Println("До побачення!")
				return
			default:
				fmt.Println("Ведіть лише цифри від 1 до 5")
			}
		}

	}

}

func main() {
	systemPlayment()
}
