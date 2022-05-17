package nft

import (
	"github.com/spikeekips/mitum-currency/currency"
	"github.com/spikeekips/mitum/base/operation"
	"github.com/spikeekips/mitum/util/hint"
)

func OperationHinter(ht hint.Hint) currency.BaseOperation {
	return currency.BaseOperation{BaseOperation: operation.EmptyBaseOperation(ht)}
}
