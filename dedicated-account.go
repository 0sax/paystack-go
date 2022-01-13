package paystack

import (
	"fmt"
)

// DedicatedNubanService handles operations related to Dedicated Accounts
// For more details see https://paystack.com/docs/api/#dedicated-nuban
type DedicatedNubanService service

type DedicatedNubanRequest struct {
	Customer      string `json:"customer,omitempty"`
	PreferredBank string `json:"preferred_bank,omitempty"`
	Subaccount    string `json:"subaccount,omitempty"`
	SplitCode     string `json:"split_code,omitempty"`
}

type DedicatedNuban struct {
	Bank struct {
		Name string `json:"name"`
		Id   int    `json:"id"`
		Slug string `json:"slug"`
	} `json:"bank"`
	AccountName   string      `json:"account_name"`
	AccountNumber string      `json:"account_number"`
	Assigned      bool        `json:"assigned"`
	Currency      string      `json:"currency"`
	Metadata      interface{} `json:"metadata"`
	Active        bool        `json:"active"`
	Id            int         `json:"id"`
	CreatedAt     string   `json:"created_at"`
	UpdatedAt     string   `json:"updated_at"`
	Assignment    struct {
		Integration  int         `json:"integration"`
		AssigneeId   int         `json:"assignee_id"`

		AssigneeType string      `json:"assignee_type"`
		Expired      bool        `json:"expired"`
		AccountType  string      `json:"account_type"`
		AssignedAt   string   `json:"assigned_at"`
		ExpiredAt    interface{} `json:"expired_at"`
	} `json:"assignment"`
	Customer struct {
		Id           int    `json:"id"`
		FirstName    string `json:"first_name"`
		LastName     string `json:"last_name"`
		Email        string `json:"email"`
		CustomerCode string `json:"customer_code"`
		Phone        string `json:"phone"`
		Metadata     struct {
		} `json:"metadata"`
		RiskAction               string      `json:"risk_action"`
		InternationalFormatPhone interface{} `json:"international_format_phone"`
	} `json:"customer"`
}

// Create creates a new dedicated nuban
// For more details see https://developers.paystack.co/v1.0/reference#create-customer
func (s *DedicatedNubanService) Create(customerCode, bank, subAccount, splitCode string) (dn *DedicatedNuban, err error) {
	u := fmt.Sprintf("/dedicated_account")
	dnReq := &DedicatedNubanRequest{
		Customer:      customerCode,
		PreferredBank: bank,
	}
	dn = &DedicatedNuban{}
	err = s.client.Call("POST", u, dnReq, dn)
	return
}
