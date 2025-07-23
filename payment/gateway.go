package payment

type PaymentGateway interface {
	Init(config map[string]string) error
	CreatePayment(amount int64, currency, receipt string) (string, error)
	VerifyPayment(paymentID, signature string) (bool, error)
}

func GetGateway(name string) PaymentGateway {
	switch name {
	case "razorpay":
		return &RazorpayGateway{}
	// case "stripe":
	//     return &StripeGateway{}
	// case "phonepe":
	//     return &PhonePeGateway{}
	default:
		return &DummyGateway{}
	}
}
