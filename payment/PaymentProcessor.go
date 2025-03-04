package payment

//go:generate mockgen -destination=mocks/mocks.go -source=PaymentProcessor.go
type PaymentProcessor interface {
	Charge(amount float64, token string) error
}
