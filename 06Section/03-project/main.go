package main

import (
	"errors"
	"fmt"
)

type Account struct {
	AccountNumber string
	Balance       float64
	OwnerName     string
}

func (account *Account) Deposit(amount float64) error {
	if amount <= 0 {
		return errors.New("deposit amount must be greater than zero")
	}
	account.Balance += amount
	fmt.Printf("Deposited : %.2f to %s, New Balance: %.2f\n", amount, account.AccountNumber, account.Balance)
	return nil
}

func (account *Account) Withdraw(amount float64) error {
	if amount <= 0 {
		return errors.New("withdraw amount must be greater than zero")
	}
	if account.Balance < amount {
		return errors.New("withdraw amount must be greater than zero")
	}
	account.Balance -= amount
	fmt.Printf("Withdraw : %.2f to %s, New Balance: %.2f\n", amount, account.AccountNumber, account.Balance)
	return nil
}

func (account *Account) GetBalance() float64 {
	return account.Balance
}

func (account *Account) String() string {
	return fmt.Sprintf("AccountNumber: %s, Owner: %s, Balance: %.2f",
		account.AccountNumber, account.OwnerName, account.Balance)
}

type SavingAccount struct {
	Account      // embedded Account struct (anonymous field)
	InterestRate float64
}

func (acc *SavingAccount) AddInterest() {
	interest := acc.InterestRate * acc.Balance
	fmt.Printf("Adding interest : %.2f\n", interest)
	err := acc.Deposit(interest)
	if err != nil {
		fmt.Printf("Error depositing interest : %v", err)
	}
}

type OverDraftAccount struct {
	Account
	OverDraftLimit float64
}

func (acc *OverDraftAccount) Withdraw(amount float64) error {
	if amount <= 0 {
		return errors.New("withdraw amount must be greater than zero")
	}
	// Allow withdrawal up to Balance + OverDraftLimit

	if acc.Balance+acc.OverDraftLimit < amount {
		return errors.New("withdraw amount must be less than or equal to available balance including overdraft")
	}
	acc.Balance -= amount
	fmt.Printf("Withdraw : %.2f to %s, New Balance: %.2f\n", amount, acc.AccountNumber, acc.Balance)
	return nil
}

func main() {
	fmt.Println("---- Bank account system")
	savAcc := SavingAccount{
		Account: Account{
			AccountNumber: "123",
			Balance:       1000,
			OwnerName:     "John Doe",
		},
		InterestRate: 0.02,
	}
	fmt.Println("------ Saving Account Operations ------")
	fmt.Println(savAcc.Account.String())

	err := savAcc.Deposit(200)
	if err != nil {
		fmt.Printf("Error depositing %.2f to saving account. %s", 200.00, err)
	}
	savAcc.AddInterest()

	err = savAcc.Withdraw(50)
	if err != nil {
		fmt.Printf("Error withdrawing %.2f to saving account. %s", 50.00, err)
	}
	fmt.Println("Final Savings Details:", savAcc.Account.String())

	ovdAcc := OverDraftAccount{
		Account: Account{
			AccountNumber: "OVD002",
			Balance:       100.00,
			OwnerName:     "Bob Spender",
		},
		OverDraftLimit: 200.00,
	}
	fmt.Println("\n--- Overdraft Account Operations ---")
	fmt.Println(ovdAcc.Account.String())

	err = ovdAcc.Deposit(50.00)
	if err != nil {
		fmt.Println("Error:", err)
	}
	err = ovdAcc.Withdraw(200.00)
	if err != nil {
		fmt.Println("Error:", err)
	}

	err = ovdAcc.Withdraw(100.00)
	if err != nil {
		fmt.Println("Error:", err)
	}
	err = ovdAcc.Withdraw(100.00)
	if err != nil {
		fmt.Println("Error (expected for overdraft limit):", err)
	}

	fmt.Println("Final Overdraft Details:", ovdAcc.Account.String())
}
