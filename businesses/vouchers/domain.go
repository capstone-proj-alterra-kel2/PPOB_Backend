package vouchers

import (
	"time"

	"gorm.io/gorm"
)

type Voucher struct {
	ID                  uint           `json:"id"`
	Voucher_Code        string         `gorm:"NOT NULL; UNIQUE_INDEX " json:"voucher_code"`
	Voucher_Description string         `json:"voucher_description"`
	Voucher_Balance     int            `json:"voucher_balance"`
	Quantity_Stock      int            `json:"quantity_stock"`
	CreatedAt           time.Time      `json:"created_at"`
	ExpiredAt           string         `json:"expired_at"`
	UpdatedAt           time.Time      `json:"updated_at"`
	DeletedAt           gorm.DeletedAt `json:"deleted_at"`
}

type VoucherInputCreate struct {
	Voucher_Code        string `gorm:"NOT NULL; UNIQUE_INDEX " json:"voucher_code"`
	Voucher_Description string `json:"voucher_description"`
	Voucher_Balance     int    `json:"voucher_balance"`
	Quantity_Stock      int    `json:"quantity_stock"`
	ExpiredAt           string `json:"expired_at"`
}

type Usecase interface {
	Create(voucherDomain *VoucherInputCreate) VoucherInputCreate
	Update(voucherDomain *Voucher) error
	Delete(voucherID int) error
	GetByCode(voucherCode string) (*Voucher, error)
	GetAll(voucherID int) []Voucher
	UseVoucher(voucherCode string) (*Voucher, error)
}

type Repository interface {
	Create(voucherDomain *VoucherInputCreate) VoucherInputCreate
	Update(voucherDomain *Voucher) error
	Delete(voucherID int) error
	GetAll(voucherID int) []Voucher
	GetByCode(voucherCode string) (*Voucher, error)
	UseVoucher(voucherCode string) (*Voucher, error)
}
