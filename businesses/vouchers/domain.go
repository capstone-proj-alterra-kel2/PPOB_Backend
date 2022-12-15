package vouchers

import (
	"time"

	"gorm.io/gorm"
)

type Voucher struct {
  ID int ``
  Code string ``
  
  CreatedAt time.Time `json:"created_at"`
  ExpiredAt time.Time `json:"expire_at,omitempty"`
  UpdatedAt time.Time`json:"updated_at"`
  DeletedAt gorm.DeletedAt`json:"deleted_at"`
}

