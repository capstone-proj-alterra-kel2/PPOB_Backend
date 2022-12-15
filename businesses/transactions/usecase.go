package transactions

import (
	"PPOB_BACKEND/businesses/products"
	"PPOB_BACKEND/businesses/users"
	"time"

	"gorm.io/gorm"
)

type transactionUsecase struct {
	transactionRepository Repository
}

func NewTransactionUsecase(transactionRepo Repository) Usecase {
	return &transactionUsecase{
		transactionRepository: transactionRepo,
	}
}

func (tu *transactionUsecase) GetAll(Page int, Size int, Sort string, Search string) ([]Domain, *gorm.DB) {
	return tu.transactionRepository.GetAll(Page, Size, Sort, Search)
}
func (tu *transactionUsecase) GetDetail(transaction_id int) (Domain, bool) {
	return tu.transactionRepository.GetDetail(transaction_id)
}
func (tu *transactionUsecase) GetTransactionHistory(user_id int) []Domain {
	return tu.transactionRepository.GetTransactionHistory(user_id)
}

func (tu *transactionUsecase) Update(trasnactionDomain *Domain, transaction_id int) (Domain, bool) {
	return tu.transactionRepository.Update(trasnactionDomain, transaction_id)
}

func (tu *transactionUsecase) Create(productDomain *products.Domain, userDomain *users.Domain, totalAmount int, productDiscount int, targetPhoneNumber string) Domain {

	transaction := Domain{
		ProductID:         int(productDomain.ID),
		ProductName:       productDomain.Name,
		UserID:            int(userDomain.ID),
		UserEmail:         userDomain.Email,
		TargetPhoneNumber: targetPhoneNumber,
		ProductPrice:      productDomain.Price,
		ProductDiscount:   productDiscount,
		TotalPrice:        totalAmount,
		TransactionDate:   time.Now(),
	}

	return tu.transactionRepository.Create(&transaction)
}
func (tu *transactionUsecase) Delete(transaction_id int) (Domain, bool) {
	return tu.transactionRepository.Delete(transaction_id)
}
