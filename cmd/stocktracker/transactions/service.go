package transactions

type TransactionService struct {
	transactionRepository TransactionRepository
}

func ProvideTransactionService(r TransactionRepository) TransactionService {
	return TransactionService{r}
}

func (service *TransactionService) GetAll() []StockTransaction {
	return service.transactionRepository.GetAll()
}
