package payment

type DummyGateway struct{}

func (d *DummyGateway) Init(config map[string]string) error { return nil }
func (d *DummyGateway) CreatePayment(amount int64, currency, receipt string) (string, error) {
	return "dummy_payment_id", nil
}
func (d *DummyGateway) VerifyPayment(paymentID, signature string) (bool, error) {
	return true, nil
}
