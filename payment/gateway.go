package payment

type PaymentGateway interface {
	Init(config map[string]string) error
	CreatePayment(amount int64, currency, receipt string) (string, error)
	VerifyPayment(paymentID, signature string) (bool, error)
	Refund(paymentID string, amount int64) (string, error)
	GenerateInvoice(paymentID string, details map[string]interface{}) (string, error)
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

// DummyGateway implementation for Refund and GenerateInvoice
func (d *DummyGateway) Refund(paymentID string, amount int64) (string, error) {
	return "dummy_refund_id", nil
}

func (d *DummyGateway) GenerateInvoice(paymentID string, details map[string]interface{}) (string, error) {
	return "dummy_invoice_id", nil
}

// RazorpayGateway implementation for Refund and GenerateInvoice
func (r *RazorpayGateway) Refund(paymentID string, amount int64) (string, error) {
	return "razorpay_refund_id", nil
}

func (r *RazorpayGateway) GenerateInvoice(paymentID string, details map[string]interface{}) (string, error) {
	return "razorpay_invoice_id", nil
}
