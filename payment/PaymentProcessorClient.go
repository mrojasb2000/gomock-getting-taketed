package payment

import "errors"

// Define the PaymentProcessorClient that is going to be tested
type PaymentProcessorClient struct {
	PaymentProcessor PaymentProcessor
}

// Define the Charge function for the client, which is going to be tested
func (c *PaymentProcessorClient) Charge(amount float64, token string) error {

	// Define the lowest possible charge
	if amount < 20 {
		return errors.New("Charge too low")
	}

	return c.PaymentProcessor.Charge(amount, token)

}
