package payment_method

import (
	"PPOB_BACKEND/businesses/payment_method"

	"gorm.io/gorm"
)

type paymentRepository struct {
	conn *gorm.DB
}

func NewPostgreSQLRepository(conn *gorm.DB) payment_method.Repository {
	return &paymentRepository{
		conn: conn,
	}
}

func (pr *paymentRepository) GetAll() []payment_method.Domain {
	var rec []Payment_Method
	pr.conn.Preload("Payment").Find(&rec)

	paymentDomain := []payment_method.Domain{}
	for _, payment := range rec {
		paymentDomain = append(paymentDomain, payment.ToDomain())
	}
	return paymentDomain
}

func (pr *paymentRepository) GetSpecificPayment(id string) payment_method.Domain {
	var payment Payment_Method

	pr.conn.Preload("Payment").First(&payment, "id=?", id)
	return payment.ToDomain()
}

func (pr *paymentRepository) CreatePayment(paymentDomain *payment_method.Domain) payment_method.Domain {
	rec := FromDomain(paymentDomain)

	result := pr.conn.Create(&rec)

	result.Last(&rec)

	return rec.ToDomain()
}

func (pr *paymentRepository) UpdatePaymentByID(id string, paymentDomain *payment_method.Domain) payment_method.Domain {
	var payment payment_method.Domain = pr.GetSpecificPayment(id)

	updatedPayment := FromDomain(&payment)

	updatedPayment.Payment_Name = paymentDomain.Payment_Name
	updatedPayment.Url_Payment = paymentDomain.Url_Payment
	updatedPayment.Icon = paymentDomain.Icon

	pr.conn.Save(&updatedPayment)

	return updatedPayment.ToDomain()
}

func (pr *paymentRepository) DeletePayment(id string) bool {
	var payment payment_method.Domain = pr.GetSpecificPayment(id)

	deletedPayment := FromDomain(&payment)

	result := pr.conn.Delete(&deletedPayment)

	if result.RowsAffected == 0 {
		return false
	}
	return true
}
