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