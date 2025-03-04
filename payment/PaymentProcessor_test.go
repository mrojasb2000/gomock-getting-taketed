package payment

import (
	mock_payment "gomock-getting-taketed/payment/mocks"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func testChange(t *testing.T) {
	mockCtl := gomock.NewController(t)
	defer mockCtl.Finish()

	mockPaymentProcessor := mock_payment.NewMockPaymentProcessor(mockCtl)

	testPaymentProcessorClient := PaymentProcessorClient{
		PaymentProcessor: mockPaymentProcessor,
	}

	mockPaymentProcessor.
		EXPECT().
		Charge(100, "test_token").
		Return(nil).
		Times(1)

	err := testPaymentProcessorClient.Charge(100, "test_token")

	assert.Nil(t, err)
}
