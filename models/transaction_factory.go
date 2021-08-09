package models

import ()

func CreateDefaultTransaction() *Transaction {
	sender := CreateInsurer()
	receiver := CreateLender()
	amount := 1000.0

	return CreateTransaction(*sender, *receiver, amount)
}

func CreateTransaction(sender Actor, receiver Actor, amount float64) *Transaction {
	sender.UpdateBalance(-amount)
	receiver.UpdateBalance(amount)

	transaction := &Transaction{
		SenderID:     int(sender.ID),
		SenderType:   sender.Type,
		ReceiverID:   int(receiver.ID),
		ReceiverType: receiver.Type,
		Amount:       amount,
	}

	transaction.Save()
	return transaction
}

func CreateDepositTransaction(borrower Actor, amount float64) *Transaction {
	borrower.UpdateBalance(-amount)

	depositTransaction := &Transaction{
		SenderID:     int(borrower.ID),
		SenderType:   borrower.Type,
		ReceiverID:   0,
		ReceiverType: "deposit",
		Amount:       amount,
	}

	depositTransaction.Save()
	return depositTransaction
}

func CreateIncomeTransaction(borrower Actor, amount float64) *Transaction {
	borrower.UpdateBalance(amount)

	incomeTransaction := &Transaction{
		SenderID:     0,
		SenderType:   "income",
		ReceiverID:   int(borrower.ID),
		ReceiverType: borrower.Type,
		Amount:       amount,
	}

	incomeTransaction.Save()
	return incomeTransaction
}
