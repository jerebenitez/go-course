package main

import (
	"errors"
)

type customer struct {
	id      int
	balance float64
}

type transactionType string

const (
	transactionDeposit    transactionType = "deposit"
	transactionWithdrawal transactionType = "withdrawal"
)

type transaction struct {
	customerID      int
	amount          float64
	transactionType transactionType
}

// Don't touch above this line

func updateBalance(c *customer, t transaction) error {
    if c.balance < t.amount {
        return errors.New("insufficient funds")
    }

    switch t.transactionType {
        case transactionDeposit:
            c.balance += t.amount
        case transactionWithdrawal:
            c.balance -= t.amount
        default:
            return errors.New("unknown transaction type")
    }

    return nil
}

