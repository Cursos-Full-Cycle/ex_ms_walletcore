package createclient

import (
	"ex_ms_walletcore/internal/entity"
	"ex_ms_walletcore/internal/gateway"
	"time"
)

type CreateClientInputDto struct {
	Name  string
	Email string
}

type CreateClientOutputDto struct {
	ID        string
	Name      string
	Email     string
	CreatedAt time.Time
	UpdateAt  time.Time
}

type CreateClientUseCase struct {
	ClientGateway gateway.ClientGateway
}

func NewCreateClientUseCase(clientGateway gateway.ClientGateway) *CreateClientUseCase {
	return &CreateClientUseCase{
		ClientGateway: clientGateway,
	}
}

func (uc *CreateClientUseCase) Execute(input CreateClientInputDto) (*CreateClientOutputDto, error) {
	client, err := entity.NewClient(input.Name, input.Email)
	if err != nil {
		return nil, err
	}
	err = uc.ClientGateway.Save(client)
	if err != nil {
		return nil, err
	}
	return &CreateClientOutputDto{
		ID:        client.ID,
		Name:      client.Name,
		Email:     client.Email,
		CreatedAt: client.CreatedAt,
		UpdateAt:  client.UpdatedAt,
	}, nil
}
