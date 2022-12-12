package payment_method

type paymentUsecase struct {
	paymentRepository Repository
}

func NewPaymentMethodUsecase(cr Repository) Usecase {
	return &paymentUsecase{
		paymentRepository: cr,
	}
}

func (cu *paymentUsecase) GetAll() []Domain {
	return cu.paymentRepository.GetAll()
}

func (cu *paymentUsecase) GetSpecificPayment(id string) Domain {
	return cu.paymentRepository.GetSpecificPayment(id)
}

func (cu *paymentUsecase) CreatePayment(paymentDomain *Domain) Domain {
	return cu.paymentRepository.CreatePayment(paymentDomain)
}

func (cu *paymentUsecase) UpdatePaymentByID(id string, paymentDomain *Domain) Domain {
	return cu.paymentRepository.UpdatePaymentByID(id, paymentDomain)
}

func (cu *paymentUsecase) DeletePayment(id string) bool {
	return cu.paymentRepository.DeletePayment(id)
}
