package payment

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// Defines the struct for StripePaymentProcessor, which will implement the PaymentProcessor interface
type StripePaymentProcessor struct{}

// Implementation of the Charge function defined in the PaymentProcessor interface
func (s *StripePaymentProcessor) Charge(amount float64, token string) error {

	// Define empty json object for use in HTTP request
	data := map[string]string{}
	json_data, err := json.Marshal(data)
	if err != nil {
		return err
	}

	// execute HTTP request
	resp, err := http.Post("https://api.stripe.com/v1/charges", "", bytes.NewBuffer(json_data))

	// Return error if any exists
	if err != nil {
		return err
	}

	// Decode the response to json
	// Note: this isn't used anywhere, as this function definition only returns an error
	// This is done purely to simplify the example
	var res map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&res)

	return nil
}
