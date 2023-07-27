package gateway

import "ex_ms_walletcore/internal/entity"

type TransactionGateway interface {
	Create(transaction *entity.Transaction) error
}
