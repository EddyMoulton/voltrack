package stocks

import "github.com/jinzhu/gorm"

type StocksRepository struct {
	db *gorm.DB
}

func ProvideStocksRepository(db *gorm.DB) StocksRepository {
	return StocksRepository{db}
}

func (r *StocksRepository) getAll() []Stock {
	allStocks := []Stock{}
	r.db.Find(&allStocks)
	return allStocks
}

func (r *StocksRepository) find(code string) Stock {
	stock := Stock{}
	r.db.Where(&Stock{Code: code}).Find(&stock)
	return stock
}

func (r *StocksRepository) add(stock Stock) {
	r.db.Create(&stock)
}
