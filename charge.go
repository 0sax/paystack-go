package paystack

import (
	"fmt"
)

// ChargeService handles operations related to bulk charges
// For more details see https://developers.paystack.co/v1.0/reference#charge-tokenize
type ChargeService service

// Card represents a Card object
type Card struct {
	Number            string `json:"card_number,omitempty"`
	CVV               string `json:"card_cvc,omitempty"`
	ExpirtyMonth      string `json:"expiry_month,omitempty"`
	ExpiryYear        string `json:"expiry_year,omitempty"`
	AddressLine1      string `json:"address_line1,omitempty"`
	AddressLine2      string `json:"address_line2,omitempty"`
	AddressLine3      string `json:"address_line3,omitempty"`
	AddressCountry    string `json:"address_country,omitempty"`
	AddressPostalCode string `json:"address_postal_code,omitempty"`
	Country           string `json:"country,omitempty"`
}

// BankAccount is used as bank in a charge request
type BankAccount struct {
	Code          string `json:"code,omitempty"`
	AccountNumber string `json:"account_number,omitempty"`
}

// ChargeRequest represents a Paystack charge request
type ChargeRequest struct {
	Email             string       `json:"email,omitempty"`
	Amount            float32      `json:"amount,omitempty"`
	Birthday          string       `json:"birthday,omitempty"`
	Card              *Card        `json:"card,omitempty"`
	Bank              *BankAccount `json:"bank,omitempty"`
	AuthorizationCode string       `json:"authorization_code,omitempty"`
	Pin               string       `json:"pin,omitempty"`
	Metadata          *Metadata    `json:"metadata,omitempty"`
	Reference         string       `json:"reference,omitempty"`
}

type ChargeResponse struct {
	Amount          int       `json:"amount"`
	Currency        string    `json:"currency"`
	TransactionDate interface{} `json:"transaction_date"`
	Status          string    `json:"status"`
	Reference       string    `json:"reference"`
	DisplayText     string    `json:"display_text"`
	UssdCode        string    `json:"ussd_code"`
	CountryCode     string    `json:"country_code"`
	Url             string    `json:"url"`
	Domain          string    `json:"domain"`
	Metadata        interface{} `json:"metadata"`
	GatewayResponse string      `json:"gateway_response"`
	Message         interface{} `json:"message"`
	Channel         string      `json:"channel"`
	IpAddress       string      `json:"ip_address"`
	Log             interface{} `json:"log"`
	Fees            int         `json:"fees"`
	Authorization   struct {
		AuthorizationCode string `json:"authorization_code"`
		Bin               string `json:"bin"`
		Last4             string `json:"last4"`
		ExpMonth          string `json:"exp_month"`
		ExpYear           string `json:"exp_year"`
		Channel           string `json:"channel"`
		CardType          string `json:"card_type"`
		Bank              string `json:"bank"`
		CountryCode       string `json:"country_code"`
		Brand             string `json:"brand"`
		Reusable          bool   `json:"reusable"`
		Signature         string `json:"signature"`
		AccountName       string `json:"account_name"`
	} `json:"authorization"`
	Customer struct {
		Id           int         `json:"id"`
		FirstName    interface{} `json:"first_name"`
		LastName     interface{} `json:"last_name"`
		Email        string      `json:"email"`
		CustomerCode string      `json:"customer_code"`
		Phone        interface{} `json:"phone"`
		Metadata     interface{} `json:"metadata"`
		RiskAction   string      `json:"risk_action"`
	} `json:"customer"`
	Plan interface{} `json:"plan"`
}

// Create submits a charge request using card details or bank details or authorization code
// For more details see https://developers.paystack.co/v1.0/reference#charge
func (s *ChargeService) Create(req *ChargeRequest) (*ChargeResponse, error) {
	resp := &ChargeResponse{}
	err := s.client.Call("POST", "/charge", req, resp)
	return resp, err
}

// Tokenize tokenizes payment instrument before a charge
// For more details see https://developers.paystack.co/v1.0/reference#charge-tokenize
func (s *ChargeService) Tokenize(req *ChargeRequest) (Response, error) {
	resp := Response{}
	err := s.client.Call("POST", "/charge/tokenize", req, &resp)
	return resp, err
}

// SubmitPIN submits PIN to continue a charge
// For more details see https://developers.paystack.co/v1.0/reference#submit-pin
func (s *ChargeService) SubmitPIN(pin, reference string) (*ChargeResponse, error) {
	data := map[string]interface{}{
		"pin": pin,
		"reference": reference,
	}
	resp := &ChargeResponse{}
	err := s.client.Call("POST", "/charge/submit_pin", data, &resp)
	return resp, err
}

// SubmitOTP submits OTP to continue a charge
// For more details see https://developers.paystack.co/v1.0/reference#submit-pin
func (s *ChargeService) SubmitOTP(otp, reference string) (*ChargeResponse, error) {
	data := map[string]interface{}{
		"otp": otp,
		"reference": reference,
	}
	resp := &ChargeResponse{}
	err := s.client.Call("POST", "/charge/submit_otp", data, &resp)
	return resp, err
}

// SubmitPhone submits Phone when requested
// For more details see https://developers.paystack.co/v1.0/reference#submit-pin
func (s *ChargeService) SubmitPhone(phone, reference string) (*ChargeResponse, error) {
	data := map[string]interface{}{
		"phone": phone,
		"reference": reference,
	}
	resp := &ChargeResponse{}
	err := s.client.Call("POST", "/charge/submit_phone", data, &resp)
	return resp, err
}

// SubmitBirthday submits Birthday when requested
// For more details see https://developers.paystack.co/v1.0/reference#submit-pin
func (s *ChargeService) SubmitBirthday(birthday, reference string) (*ChargeResponse, error) {
	data := map[string]interface{}{
		"birthday": birthday,
		"reference": reference,
	}
	resp := &ChargeResponse{}
	err := s.client.Call("POST", "/charge/submit_birthday", data, &resp)
	return resp, err
}

// CheckPending returns pending charges
// When you get "pending" as a charge status, wait 30 seconds or more,
// then make a check to see if its status has changed. Don't call too early as you may get a lot more pending than you should.
// For more details see https://developers.paystack.co/v1.0/reference#check-pending-charge
func (s *ChargeService) CheckPending(reference string) (Response, error) {
	u := fmt.Sprintf("/charge/%s", reference)
	resp := Response{}
	err := s.client.Call("GET", u, nil, &resp)
	return resp, err
}
