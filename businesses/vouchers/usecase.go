package vouchers



type voucherUsecase struct {
	voucherRepository Repository
}

func NewVoucherUsecase(voucherRepo Repository) Usecase {
	return &voucherUsecase{
		voucherRepository: voucherRepo,
	}
}

func (vu *voucherUsecase) Create(voucherDomain *VoucherInputCreate) VoucherInputCreate {
	return vu.voucherRepository.Create(voucherDomain)
}

func (vu *voucherUsecase) Update(voucherDomain *Voucher) error {
	return vu.voucherRepository.Update(voucherDomain)
}

func (vu *voucherUsecase) Delete(voucherID int) error {
	return vu.voucherRepository.Delete(voucherID)
}

func (vu *voucherUsecase) GetAll(voucherID int) []Voucher {
	return vu.voucherRepository.GetAll(voucherID)
}

func (vu *voucherUsecase) GetByCode(voucherCode string) (*Voucher, error) {
	return vu.voucherRepository.GetByCode(voucherCode)
}

func (vu *voucherUsecase) UseVoucher(voucherCode string) (*Voucher, error) {
	return vu.voucherRepository.UseVoucher(voucherCode)
}
