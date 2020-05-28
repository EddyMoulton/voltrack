package transactions

// TransactionListDto is for transfering a list of TransactionDto
type TransactionListDto struct {
	Transactions []TransactionDTO `json:"transactions" binding:"required"` // List of transactions
}
