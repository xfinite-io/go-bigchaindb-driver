package interfaces

import (
	"github.com/xfinite-io/go-bigchaindb-driver/pkg/client"
	"github.com/xfinite-io/go-bigchaindb-driver/pkg/transaction"
)

type Transaction interface {
	New(operation string,
		asset transaction.Asset,
		metadata transaction.Metadata,
		inputs []transaction.Input,
		outputs []transaction.Output,
	) (*transaction.Transaction, error)
	NewCreateTransaction(
		asset transaction.Asset,
		metadata transaction.Metadata,
		outputs []transaction.Output,
		issuers []string,
	) (*transaction.Transaction, error)
	//NewTransferTransaction() (*transaction.Transaction, error)

	Sign(keyPairs []client.KeyPair) error
}
