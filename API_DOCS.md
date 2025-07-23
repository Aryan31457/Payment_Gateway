# Payment Gateway API Documentation

Base URL:  
```
http://localhost:8080
```

---

## 1. Create Payment

**Endpoint:**
```
POST /pay?provider={provider}
```
- `{provider}` can be `dummy` or `razorpay` (default is `dummy` if not specified)

**Request:**
```json
{
  "amount": 1000,
  "currency": "INR",
  "receipt": "receipt#1"
}
```

**Response (Dummy):**
```json
{
  "payment_id": "dummy_payment_id",
  "error": ""
}
```

**Response (Razorpay):**
```json
{
  "payment_id": "order_Jv1X2Y3Z4A5B6C",
  "error": ""
}
```

---

## 2. Verify Payment

**Endpoint:**
```
POST /verify?provider={provider}
```
- `{provider}` can be `dummy` or `razorpay` (default is `dummy` if not specified)

**Request:**
```json
{
  "payment_id": "order_Jv1X2Y3Z4A5B6C",
  "signature": "razorpay_signature_here"
}
```

**Response:**
```json
{
  "verified": true,
  "error": ""
}
```

---

## Error Handling
If an error occurs, the `error` field will be populated and the HTTP status code will be 400 or 500.

---

## Example Usage (cURL)

**Create Payment (Dummy):**
```sh
curl -X POST "http://localhost:8080/pay?provider=dummy" \
  -H "Content-Type: application/json" \
  -d '{"amount":1000,"currency":"INR","receipt":"receipt#1"}'
```

**Create Payment (Razorpay):**
```sh
curl -X POST "http://localhost:8080/pay?provider=razorpay" \
  -H "Content-Type: application/json" \
  -d '{"amount":1000,"currency":"INR","receipt":"receipt#1"}'
```

**Verify Payment (Razorpay):**
```sh
curl -X POST "http://localhost:8080/verify?provider=razorpay" \
  -H "Content-Type: application/json" \
  -d '{"payment_id":"order_Jv1X2Y3Z4A5B6C","signature":"razorpay_signature_here"}'
``` 

---

## 1. **Interface Design**

Let’s extend your `PaymentGateway` interface:

```go
type PaymentGateway interface {
    Init(config map[string]string) error
    CreatePayment(amount int64, currency, receipt string) (string, error)
    VerifyPayment(paymentID, signature string) (bool, error)
    Refund(paymentID string, amount int64) (string, error)
    GenerateInvoice(paymentID string, details map[string]interface{}) (string, error)
}
```

### **Explanation:**
- `Refund(paymentID string, amount int64) (string, error)`
  - Initiates a refund for a given payment ID and amount.
  - Returns a refund ID or status message.
- `GenerateInvoice(paymentID string, details map[string]interface{}) (string, error)`
  - Generates an invoice for a payment.
  - `details` can include customer info, line items, etc.
  - Returns an invoice ID or a link to the invoice.

---

## 2. **Dummy Implementation Example**

```go
// payment/dummy.go

func (d *DummyGateway) Refund(paymentID string, amount int64) (string, error) {
    return "dummy_refund_id", nil
}

func (d *DummyGateway) GenerateInvoice(paymentID string, details map[string]interface{}) (string, error) {
    return "dummy_invoice_id", nil
}
```

---

## 3. **Razorpay Implementation Example (Stub)**

```go
// payment/razorpay.go

func (r *RazorpayGateway) Refund(paymentID string, amount int64) (string, error) {
    // Use Razorpay's Refund API here
    return "razorpay_refund_id", nil
}

func (r *RazorpayGateway) GenerateInvoice(paymentID string, details map[string]interface{}) (string, error) {
    // Use Razorpay's Invoice API here
    return "razorpay_invoice_id", nil
}
```
*(You can later fill in the real API calls as needed.)*

---

## 4. **API Endpoints (Suggested)**

- `POST /refund?provider=...`
  - Body: `{ "payment_id": "...", "amount": 1000 }`
  - Response: `{ "refund_id": "...", "error": "" }`
- `POST /invoice?provider=...`
  - Body: `{ "payment_id": "...", "details": { ... } }`
  - Response: `{ "invoice_id": "...", "error": "" }`

---

## 5. **Summary**

- You can now support refunds and invoice generation for any payment provider.
- Each provider implements these methods according to its own API.
- Your API can expose `/refund` and `/invoice` endpoints for clients to use.

---

**Would you like me to:**
- Add these methods to your interface and dummy/razorpay implementations?
- Scaffold the new API endpoints as well?
- Or explain how to call the real Razorpay/Stripe APIs for these features?

Let me know your preference! 