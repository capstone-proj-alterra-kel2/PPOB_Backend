package transactions

import (
	"PPOB_BACKEND/businesses/transactions"
	"strings"

	"gorm.io/gorm"
)

type transactionRepository struct {
	conn *gorm.DB
}

func NewTransactionRepository(conn *gorm.DB) *transactionRepository {
	return &transactionRepository{
		conn: conn,
	}
}

func (tr *transactionRepository) GetAll(Page int, Size int, Sort string, Search string) ([]transactions.Domain, *gorm.DB) {
	var tx []Transaction

	var sort string
	var search string
	var model *gorm.DB

	if strings.HasPrefix(Sort, "-") {
		sort = Sort[1:] + " DESC"
	} else {
		sort = Sort[0:] + " ASC"
	}

	model = tr.conn.Order(sort).Find(&tx).Model(&tx)

	if Search != "" {
		search = "%" + Search + "%"
		model = tr.conn.Order(sort).Model(&tx).Where("name LIKE ?", search)
	}

	tr.conn.Offset(Page).Limit(Size).Order(sort).Find(&tx)
	transactionDomain := []transactions.Domain{}

	for _, trasaction := range tx {
		transactionDomain = append(transactionDomain, trasaction.ToDomain())
	}

	return transactionDomain, model
}

func (tr *transactionRepository) GetDetail(transaction_id int) (transactions.Domain, bool) {
	var tx Transaction

	if checkTransaction := tr.conn.First(&tx, transaction_id).Error; checkTransaction != gorm.ErrRecordNotFound {
		return tx.ToDomain(), true
	}

	return tx.ToDomain(), false
}

func (tr *transactionRepository) GetTransactionHistory(user_id int) []transactions.Domain {
	var tx []Transaction

	tr.conn.Where("user_id = ?", user_id).Order("transaction_date DESC").Find(&tx)
	transactions := []transactions.Domain{}

	for _, transaction := range tx {
		transactions = append(transactions, transaction.ToDomain())
	}

	return transactions
}

func (tr *transactionRepository) Create(transactionDomain *transactions.Domain) transactions.Domain {
	transactionData := FromDomain(transactionDomain)

	tr.conn.Create(&transactionData)
	transactionDomain.ID = transactionData.ID

	return *transactionDomain
}

func (tr *transactionRepository) Update(transactionDomain *transactions.Domain, transaction_id int) (transactions.Domain, bool) {
	transactionData := FromDomain(transactionDomain)
	var tx Transaction

	if checkTransaction := tr.conn.First(&tx, transaction_id).Error; checkTransaction == gorm.ErrRecordNotFound {
		return transactionData.ToDomain(), false
	}

	tr.conn.Model(&transactionData).Where("id = ?", transaction_id).Updates(
		Transaction{
			TargetPhoneNumber: transactionData.TargetPhoneNumber,
		},
	)

	return transactionData.ToDomain(), true
}

func (tr *transactionRepository) Delete(transaction_id int) (transactions.Domain, bool) {
	var tx Transaction

	if checkTransaction := tr.conn.First(&tx, transaction_id).Error; checkTransaction == gorm.ErrRecordNotFound {
		return tx.ToDomain(), false
	}

	tr.conn.Unscoped().Delete(&tx).Where("id = ?", transaction_id)

	return tx.ToDomain(), true
}
