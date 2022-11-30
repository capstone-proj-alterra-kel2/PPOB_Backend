package stocks

type stockUsecase struct {
	stockRepository Repository
}

func NewStockUseCase(sr Repository) Usecase {
	return &stockUsecase{
		stockRepository: sr,
	}
}

func (su *stockUsecase) GetAll() []Domain {
	return su.stockRepository.GetAll()
}

func (su *stockUsecase) Get(stock_id int) Domain {
	return su.stockRepository.Get(stock_id)
}
func (su *stockUsecase) Create(stockDomain *Domain) Domain {
	return su.stockRepository.Create(stockDomain)
}
func (su *stockUsecase) Update(stockDomain *Domain, stock_id int) Domain {
	return su.stockRepository.Update(stockDomain, stock_id)
}
func (su *stockUsecase) Delete(stock_id int) Domain {
	return su.stockRepository.Delete(stock_id)
}
