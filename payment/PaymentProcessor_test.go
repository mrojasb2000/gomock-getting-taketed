package payment

import (
	"errors"
	mock_payment "gomock-getting-taketed/payment/mocks"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestPaymentChangeReturnSuccessWhenCallbackWithValidParams(t *testing.T) {
	mockCtl := gomock.NewController(t)

	mockPaymentProcessor := mock_payment.NewMockPaymentProcessor(mockCtl)

	testPaymentProcessorClient := &PaymentProcessorClient{
		PaymentProcessor: mockPaymentProcessor,
	}

	defer mockCtl.Finish()

	mockPaymentProcessor.
		EXPECT().
		Charge(gomock.Any(), gomock.Any()).
		Return(nil).
		Times(1)

	err := testPaymentProcessorClient.Charge(100.0, "test_token")

	assert.Nil(t, err)
}

func TestPaymentChangeReturnFailedWhenCallbackWithValidParams(t *testing.T) {
	mockCtl := gomock.NewController(t)

	mockPaymentProcessor := mock_payment.NewMockPaymentProcessor(mockCtl)

	testPaymentProcessorClient := &PaymentProcessorClient{
		PaymentProcessor: mockPaymentProcessor,
	}

	defer mockCtl.Finish()

	mockPaymentProcessor.
		EXPECT().
		Charge(gomock.Any(), gomock.Any()).
		Return(errors.New("Charge too low")).
		AnyTimes()

	err := testPaymentProcessorClient.Charge(10.0, "test_token")

	assert.NotNil(t, err)
	assert.Equal(t, "Charge too low", err.Error())
}
