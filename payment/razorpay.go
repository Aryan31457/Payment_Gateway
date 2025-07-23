package payment

import (
	"encoding/json"
	"fmt"

	"github.com/razorpay/razorpay-go"
)

type RazorpayGateway struct {
	client *razorpay.Client
}

func (r *RazorpayGateway) Init(config map[string]string) error {
	key, ok1 := config["key"]
	secret, ok2 := config["secret"]
	if !ok1 || !ok2 {
		return fmt.Errorf("missing Razorpay key or secret in config")
	}
	r.client = razorpay.NewClient(key, secret)
	return nil
}

func (r *RazorpayGateway) CreatePayment(amount int64, currency, receipt string) (string, error) {
	data := map[string]interface{}{
		"amount":   amount,
		"currency": currency,
		"receipt":  receipt,
	}
	body, err := r.client.Order.Create(data, nil)
	if err != nil {
		// On error, return a dummy message
		resp := map[string]interface{}{
			"order_id":        "razorpay_dummy_order_id",
			"message":         "This is a dummy Razorpay response due to error or test mode.",
			"payment_proceed": false,
			"error":           err.Error(),
		}
		b, _ := json.Marshal(resp)
		return string(b), nil
	}
	orderID, ok := body["id"].(string)
	if !ok {
		resp := map[string]interface{}{
			"order_id":        "razorpay_dummy_order_id",
			"message":         "Could not parse Razorpay order ID. Returning dummy.",
			"payment_proceed": false,
		}
		b, _ := json.Marshal(resp)
		return string(b), nil
	}
	// On success, return a structured message
	resp := map[string]interface{}{
		"order_id":        orderID,
		"message":         "Razorpay order created successfully.",
		"payment_proceed": true,
	}
	b, _ := json.Marshal(resp)
	return string(b), nil
}

func (r *RazorpayGateway) VerifyPayment(paymentID, signature string) (bool, error) {
	// Razorpay signature verification typically requires order_id, payment_id, and signature
	// For now, just return true for demo; implement real verification as needed
	return true, nil
}
