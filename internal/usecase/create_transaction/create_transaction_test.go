package createtransaction

import (
	"ex_ms_walletcore/internal/entity"
	"ex_ms_walletcore/internal/event"
	"ex_ms_walletcore/pkg/events"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type AccountGatewayMock struct {
	mock.Mock
}

func (m *AccountGatewayMock) Save(account *entity.Account) error {
	args := m.Called(account)
	return args.Error(0)
}

func (m *AccountGatewayMock) FindById(id string) (*entity.Account, error) {
	args := m.Called(id)
	return args.Get(0).(*entity.Account), args.Error(1)
}

type TransactionGatewayMock struct {
	mock.Mock
}

func (m *TransactionGatewayMock) Create(transaction *entity.Transaction) error {
	args := m.Called(transaction)
	return args.Error(0)
}

func TestCreateTransactionUseCase_Execute(t *testing.T) {
	client1, _ := entity.NewClient("Client 1", "a@a.com")
	account1 := entity.NewAccount(client1)
	account1.Credit(1000)

	client2, _ := entity.NewClient("Client 2", "b@b.com")
	account2 := entity.NewAccount(client2)
	account2.Credit(1000)

	mockAccountGateway := &AccountGatewayMock{}
	mockAccountGateway.On("FindById", account1.ID).Return(account1, nil)
	mockAccountGateway.On("FindById", account2.ID).Return(account2, nil)

	mockTransactionGateway := &TransactionGatewayMock{}
	mockTransactionGateway.On("Create", mock.Anything).Return(nil)

	dispatcher := events.NewEventDispatcher()
	event := event.NewTransactionCreated()

	uc := NewCreateTransactionUseCase(mockTransactionGateway, mockAccountGateway, dispatcher, event)

	input := CreateTransactionInputDto{
		AccountIDFrom: account1.ID,
		AccountIDTo:   account2.ID,
		Amount:        100,
	}

	output, err := uc.Execute(input)

	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.NotEmpty(t, output.ID)
	mockAccountGateway.AssertExpectations(t)
	mockAccountGateway.AssertNumberOfCalls(t, "FindById", 2)
	mockTransactionGateway.AssertExpectations(t)
	mockTransactionGateway.AssertNumberOfCalls(t, "Create", 1)

}
