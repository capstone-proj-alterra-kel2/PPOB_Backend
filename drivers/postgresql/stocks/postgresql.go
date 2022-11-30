package stocks

import (
	"PPOB_BACKEND/businesses/stocks"

	"gorm.io/gorm"
)

type stockRepository struct {
	conn *gorm.DB
}

func NewPostgreSQLRepository(conn *gorm.DB) stocks.Repository {
	return &stockRepository{
		conn: conn,
	}
}

func (sr *stockRepository) GetAll() []stocks.Domain {
	var stockData []Stock

	sr.conn.Preload("Products").Find(&stockData)

	stockDomain := []stocks.Domain{}
	for _, stock := range stockData {
		stockDomain = append(stockDomain, stock.ToDomain())
	}

	return stockDomain
}

func (sr *stockRepository) Get(stock_id int) stocks.Domain {
	var stockData Stock

	sr.conn.Preload("Products").First(&stockData, stock_id)
	return stockData.ToDomain()
}

func (sr *stockRepository) Create(stockDomain *stocks.Domain) stocks.Domain {
	stockData := FromDomain(stockDomain)

	sr.conn.Create(&stockData)
	return stockData.ToDomain()
}

func (sr *stockRepository) Update(stockDomain *stocks.Domain, stock_id int) stocks.Domain {
	stockData := FromDomain(stockDomain)
	sr.conn.Model(&stockData).Where("id = ?", stock_id).Updates(
		Stock{
			Quantity: stockDomain.Quantity,
		},
	)

	return stockData.ToDomain()
}

func (sr *stockRepository) Delete(stock_id int) stocks.Domain {
	var stockData Stock

	sr.conn.Unscoped().Where("id = ?", stock_id).Delete(&stockData)
	return stockData.ToDomain()
}
