package usecase

import (
	"loket-app/modules/transaction/query"
)

type transactionUseCaseImpl struct {
	TransactionQuery query.TransactionQuery
}

func NewTransactionUseCase(txQuery query.TransactionQuery) TransactionUseCase {
	return &transactionUseCaseImpl{
		TransactionQuery: txQuery,
	}
}

type TransactionUseCase interface {
}
