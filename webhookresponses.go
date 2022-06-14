package paystack

import (
	"encoding/json"
	"time"
)

type WebhookObject struct {
	Event string          `json:"event"`
	Data  json.RawMessage `json:"data"`
}
type WHCustomerIdentification struct {
	CustomerId     string `json:"customer_id"`
	CustomerCode   string `json:"customer_code"`
	Email          string `json:"email"`
	Identification struct {
		Country string `json:"country"`
		Type    string `json:"type"`
		Value   string `json:"value"`
	} `json:"identification"`
}
type WHCharge struct {
	Id              string      `json:"id"`
	Domain          string      `json:"domain"`
	Status          string      `json:"status"`
	Reference       string      `json:"reference"`
	Amount          int         `json:"amount"`
	Message         interface{} `json:"message"`
	GatewayResponse string      `json:"gateway_response"`
	PaidAt          time.Time   `json:"paid_at"`
	CreatedAt       time.Time   `json:"created_at"`
	Channel         string      `json:"channel"`
	Currency        string      `json:"currency"`
	IpAddress       string      `json:"ip_address"`
	Metadata        interface{} `json:"metadata"`
	Log             struct {
		TimeSpent      int           `json:"time_spent"`
		Attempts       int           `json:"attempts"`
		Authentication string        `json:"authentication"`
		Errors         int           `json:"errors"`
		Success        bool          `json:"success"`
		Mobile         bool          `json:"mobile"`
		Input          []interface{} `json:"input"`
		Channel        interface{}   `json:"channel"`
		History        []struct {
			Type    string `json:"type"`
			Message string `json:"message"`
			Time    int    `json:"time"`
		} `json:"history"`
	} `json:"log"`
	Fees     interface{} `json:"fees"`
	Customer struct {
		Id           int         `json:"id"`
		FirstName    string      `json:"first_name"`
		LastName     string      `json:"last_name"`
		Email        string      `json:"email"`
		CustomerCode string      `json:"customer_code"`
		Phone        string      `json:"phone"`
		Metadata     interface{} `json:"metadata"`
		RiskAction   string      `json:"risk_action"`
	} `json:"customer"`
	Authorization struct {
		AuthorizationCode         string `json:"authorization_code"`
		Bin                       string `json:"bin"`
		Last4                     string `json:"last4"`
		ExpMonth                  string `json:"exp_month"`
		ExpYear                   string `json:"exp_year"`
		CardType                  string `json:"card_type"`
		Bank                      string `json:"bank"`
		CountryCode               string `json:"country_code"`
		Brand                     string `json:"brand"`
		AccountName               string `json:"account_name"`
		Channel                   string `json:"channel"`
		Reusable                  bool   `json:"reusable"`
		Signature                 interface{}
		SenderBank                string `json:"sender_bank"`
		SenderBankAccountNumber   string `json:"sender_bank_account_number"`
		SenderCountry             string `json:"sender_country"`
		SenderName                string `json:"sender_name"`
		Narration                 string `json:"narration"`
		ReceiverBankAccountNumber string `json:"receiver_bank_account_number"`
		ReceiverBank              string `json:"receiver_bank"`
	} `json:"authorization"`
	Plan struct {
	} `json:"plan"`
}
type WHCustomField struct {
	DisplayName  string      `json:"display_name"`
	VariableName string      `json:"variable_name"`
	Value        interface{} `json:"value"`
}
type WHChargeMetaData struct {
	ReceiverAccountNumber string          `json:"receiver_account_number"`
	ReceiverBank          string          `json:"receiver_bank"`
	CustomFields          []WHCustomField `json:"custom_fields"`
	Scheduled             bool            `json:"scheduled"`
}
type WHTransfer struct {
	Amount      int         `json:"amount"`
	Currency    string      `json:"currency"`
	Domain      string      `json:"domain"`
	Failures    interface{} `json:"failures"`
	Id          string      `json:"id"`
	Integration struct {
		Id           string `json:"id"`
		IsLive       bool   `json:"is_live"`
		BusinessName string `json:"business_name"`
	} `json:"integration"`
	Reason        string      `json:"reason"`
	Reference     string      `json:"reference"`
	Source        string      `json:"source"`
	SourceDetails interface{} `json:"source_details"`
	Status        string      `json:"status"`
	TitanCode     interface{} `json:"titan_code"`
	TransferCode  string      `json:"transfer_code"`
	TransferredAt interface{} `json:"transferred_at"`
	Recipient     struct {
		Active        bool        `json:"active"`
		Currency      string      `json:"currency"`
		Description   string      `json:"description"`
		Domain        string      `json:"domain"`
		Email         interface{} `json:"email"`
		Id            string      `json:"id"`
		Integration   int         `json:"integration"`
		Metadata      interface{} `json:"metadata"`
		Name          string      `json:"name"`
		RecipientCode string      `json:"recipient_code"`
		Type          string      `json:"type"`
		IsDeleted     bool        `json:"is_deleted"`
		Details       struct {
			AccountNumber string      `json:"account_number"`
			AccountName   interface{} `json:"account_name"`
			BankCode      string      `json:"bank_code"`
			BankName      string      `json:"bank_name"`
		} `json:"details"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	} `json:"recipient"`
	Session struct {
		Provider interface{} `json:"provider"`
		Id       interface{} `json:"id"`
	} `json:"session"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (pwo *WebhookObject) UnmarshalAsCustomerIdent() (ci *WHCustomerIdentification, err error) {
	err = json.Unmarshal([]byte(pwo.Data), &ci)
	return
}
func (pwo *WebhookObject) UnmarshalAsCharge() (c *WHCharge, err error) {
	err = json.Unmarshal([]byte(pwo.Data), &c)
	return
}
func (pwo *WebhookObject) UnmarshalAsTransfer() (t *WHTransfer, err error) {
	err = json.Unmarshal([]byte(pwo.Data), &t)
	return
}
