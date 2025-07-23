# Go Payment Gateway Abstraction Example

## Overview

This project demonstrates how to use Go interfaces to build a flexible payment gateway system that can support multiple providers (e.g., Dummy, Razorpay, Stripe, PhonePe, etc.) with a unified API. The architecture allows you to easily add or swap payment providers without changing your business logic or API endpoints.

---

## How It Works

### 1. **The PaymentGateway Interface**

At the core is the `PaymentGateway` interface, which defines the contract for all payment providers:

```go
// payment/gateway.go

// PaymentGateway defines the contract for all payment providers
// Any provider (Razorpay, Stripe, Dummy, etc.) must implement these methods
// to be used interchangeably in the application.
type PaymentGateway interface {
    Init(config map[string]string) error
    CreatePayment(amount int64, currency, receipt string) (string, error)
    VerifyPayment(paymentID, signature string) (bool, error)
    Refund(paymentID string, amount int64) (string, error)
    GenerateInvoice(paymentID string, details map[string]interface{}) (string, error)
}
```

### 2. **Implementing the Interface**

Each payment provider implements the interface in its own file:

- **DummyGateway** (for testing, no real payments)
- **RazorpayGateway** (real or stubbed Razorpay integration)

Example (Dummy):
```go
// payment/dummy.go
func (d *DummyGateway) CreatePayment(amount int64, currency, receipt string) (string, error) {
    return "dummy_payment_id", nil
}
func (d *DummyGateway) Refund(paymentID string, amount int64) (string, error) {
    return "dummy_refund_id", nil
}
```

Example (Razorpay):
```go
// payment/razorpay.go
func (r *RazorpayGateway) CreatePayment(amount int64, currency, receipt string) (string, error) {
    // Real Razorpay API call or stub
    return "razorpay_order_id", nil
}
func (r *RazorpayGateway) Refund(paymentID string, amount int64) (string, error) {
    // Real Razorpay refund logic or stub
    return "razorpay_refund_id", nil
}
```

### 3. **Factory Pattern for Gateway Selection**

The `GetGateway` function returns the correct implementation based on a string name (e.g., from a query parameter):

```go
// payment/gateway.go
func GetGateway(name string) PaymentGateway {
    switch name {
    case "razorpay":
        return &RazorpayGateway{}
    default:
        return &DummyGateway{}
    }
}
```

### 4. **Contextual Example: Dynamic Provider Selection in API**

In `main.go`, the API handler reads the `provider` query parameter and uses the factory to get the right gateway:

```go
provider := r.URL.Query().Get("provider")
if provider == "" {
    provider = "dummy"
}
gateway := payment.GetGateway(provider)
```

Now, all payment actions (`CreatePayment`, `Refund`, etc.) are called on the `gateway` variable, which could be any provider implementing the interface.

---

## Extending the System

To add a new payment provider (e.g., Stripe):
1. Create a new file (e.g., `stripe.go`) and implement all methods of `PaymentGateway`.
2. Add a case to `GetGateway`:
   ```go
   case "stripe":
       return &StripeGateway{}
   ```
3. Now you can use `provider=stripe` in your API calls.

---

## Benefits of This Approach
- **Loose coupling:** Business logic and API code do not depend on any specific payment provider.
- **Easy extensibility:** Add or swap providers with minimal changes.
- **Unified API:** All providers expose the same methods, so the rest of your code stays the same.

---

## Example API Usage

- `POST /pay?provider=razorpay` — Create a payment with Razorpay
- `POST /pay?provider=dummy` — Create a payment with the dummy gateway
- `POST /refund?provider=razorpay` — Refund a Razorpay payment
- `POST /invoice?provider=dummy` — Generate a dummy invoice

See `API_DOCS.md` for full endpoint and request/response details.

---

## Summary

This project is a practical demonstration of Go interfaces and the factory pattern to build a scalable, maintainable, and extensible payment gateway system. 