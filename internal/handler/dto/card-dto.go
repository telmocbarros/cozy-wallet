package dto

type CreateCardRequest struct {
	CardholderName string `json:"cardholder_name"`
	CardNumber     string `json:"card_number"`
	ExpiryMonth    int    `json:"expiry_month"`
	ExpiryYear     int    `json:"expiry_year"`
}

type UpdateCardRequest struct {
	CardholderName *string `json:"cardholder_name,omitempty"`
	CardNumber     *string `json:"card_number,omitempty"`
	ExpiryMonth    *int    `json:"expiry_month,omitempty"`
	ExpiryYear     *int    `json:"expiry_year,omitempty"`
}
