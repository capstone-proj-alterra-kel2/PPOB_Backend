package payment_method

import (
	paymentUsecase "PPOB_BACKEND/businesses/payment_method"
	"time"

	"gorm.io/gorm"
)

type Payment_Method struct {
	ID           uint           `json:"id" gorm:"size:100;primaryKey"`
	Payment_Name string         `json:"payment_name"`
	Url_Payment  string         `json:"url_payment"`
	Icon         string         `json:"icon"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at"`
}

func FromDomain(domain *paymentUsecase.Domain) *Payment_Method {
	return &Payment_Method{
		ID:           domain.ID,
		Payment_Name: domain.Payment_Name,
		Url_Payment:  domain.Payment_Name,
		Icon:         domain.Icon,
		CreatedAt:    domain.CreatedAt,
		UpdatedAt:    domain.UpdatedAt,
		DeletedAt:    domain.DeletedAt,
	}
}

func (rec *Payment_Method) ToDomain() paymentUsecase.Domain {
	return paymentUsecase.Domain{
		ID:           rec.ID,
		Payment_Name: rec.Payment_Name,
		Url_Payment:  rec.Url_Payment,
		CreatedAt:    rec.CreatedAt,
		UpdatedAt:    rec.UpdatedAt,
		DeletedAt:    rec.DeletedAt,
	}
}