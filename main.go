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
type Playment interface {
	Commissions() float64
	Debited() float64
	Processing() bool
}

type PaymentBase struct {
	Name       string
	Balance    float64
	FeePercent float64
}

// Structur
type KreditCard struct{ PaymentBase }
type PayPal struct{ PaymentBase }
type Cach struct{ PaymentBase }
type BankTransfer struct{ PaymentBase }

// Func for commission accounts
func (p PaymentBase) Commissions() float64 {
	return p.Balance * p.FeePercent / 100
}

// Func for payments with commissions
func (p PaymentBase) Debited() float64 {
	return p.Balance + p.Commissions()
}

// Func for verification of funds
func (p PaymentBase) Processing() bool {
	if p.Balance >= 10.00 && p.Balance <= 50000.00 {
		fmt.Println("✅ Платіж успішно оброблено!")
		return true
	}

	fmt.Println("📛 Сума платежу повинна бути від 10 до 50 000 грн!!!!!")
	os.Exit(0)
	return false
}

// Func for information output
func inputInfo(p PaymentBase) {
	fmt.Println("------------------------------")
	fmt.Printf("Ведіть вашу сумму на балансі:")
	fmt.Scanln(&p.Balance)
	fmt.Println("------------------------------")
	fmt.Printf("Обробляємо платіж на суму %.2f грн...\n", p.Balance)
	fmt.Println("==================================")
	fmt.Printf("💳 %s\n", p.Name)
	fmt.Printf("%t\n", p.Processing())
	fmt.Printf("💰 Сума: %.2f грн\n", p.Balance)
	fmt.Printf("💸 Комісія: %.2f грн (%.2f%%)\n", p.Commissions(), p.FeePercent)
	fmt.Printf("📊 До списання: %.2f грн\n", p.Debited())
	fmt.Println("Дякуємо за покупку!")
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
			inputNumber := getIntInput("Оберіть метод оплати (1-5): ")
			fmt.Println("==================================")
			switch inputNumber {
			case 1:
				kreditCard := KreditCard{PaymentBase{"Кередитна картка", 0, 1.5}}
				inputInfo(kreditCard.PaymentBase)
			case 2:
				payPal := KreditCard{PaymentBase{" Pay Pal", 0, 3.5}}
				inputInfo(payPal.PaymentBase)
			case 3:
				cach := KreditCard{PaymentBase{"Готівка", 0, 0}}
				inputInfo(cach.PaymentBase)
			case 4:
				bankTransfer := KreditCard{PaymentBase{"Банківський переказ", 0, 2.0}}
				inputInfo(bankTransfer.PaymentBase)
			case 5:
				fmt.Println("Допобачення!!!")
				os.Exit(0)
			default:
				fmt.Println("Ведіть лише чифри від 1-5, спробуйте ще раз")
			}
		} else {
			os.Exit(0)
		}

	}

}

func main() {
	systemPlayment()
}
