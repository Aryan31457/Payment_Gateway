package main

import (
	"encoding/json"
	"log"
	"net/http"
	"payment/payment"
)

type CreatePaymentRequest struct {
	Amount   int64  `json:"amount"`
	Currency string `json:"currency"`
	Receipt  string `json:"receipt"`
}

type CreatePaymentResponse struct {
	PaymentID string `json:"payment_id"`
	Error     string `json:"error,omitempty"`
}

type VerifyPaymentRequest struct {
	PaymentID string `json:"payment_id"`
	Signature string `json:"signature"`
}

type VerifyPaymentResponse struct {
	Verified bool   `json:"verified"`
	Error    string `json:"error,omitempty"`
}

func createPaymentHandler(w http.ResponseWriter, r *http.Request) {
	provider := r.URL.Query().Get("provider")
	if provider == "" {
		provider = "dummy" // default
	}
	gateway := payment.GetGateway(provider)
	// For Razorpay, you should use real keys in production
	if provider == "razorpay" {
		gateway.Init(map[string]string{
			"key":    "YOUR_RAZORPAY_KEY",
			"secret": "YOUR_RAZORPAY_SECRET",
		})
	} else {
		gateway.Init(map[string]string{})
	}

	var req CreatePaymentRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	paymentID, err := gateway.CreatePayment(req.Amount, req.Currency, req.Receipt)
	resp := CreatePaymentResponse{PaymentID: paymentID}
	if err != nil {
		resp.Error = err.Error()
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func verifyPaymentHandler(w http.ResponseWriter, r *http.Request) {
	provider := r.URL.Query().Get("provider")
	if provider == "" {
		provider = "dummy" // default
	}
	gateway := payment.GetGateway(provider)
	if provider == "razorpay" {
		gateway.Init(map[string]string{
			"key":    "YOUR_RAZORPAY_KEY",
			"secret": "YOUR_RAZORPAY_SECRET",
		})
	} else {
		gateway.Init(map[string]string{})
	}

	var req VerifyPaymentRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	verified, err := gateway.VerifyPayment(req.PaymentID, req.Signature)
	resp := VerifyPaymentResponse{Verified: verified}
	if err != nil {
		resp.Error = err.Error()
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func main() {
	http.HandleFunc("/pay", createPaymentHandler)
	http.HandleFunc("/verify", verifyPaymentHandler)

	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
