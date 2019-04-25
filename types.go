package main

import (
	"encoding/json"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
)

// Transactions is an alias for transactions.
type Transactions []*Transaction

// Transaction is a tx wrapper with extra fields.
type Transaction struct {
	*types.Transaction
}

// UnmarshalJSON decodes data into object.
func (t *Transaction) UnmarshalJSON(data []byte) error {
	var tx types.Transaction
	err := json.Unmarshal(data, &tx)
	if err != nil {
		return err
	}
	t.Transaction = &tx
	return nil
}

// MarshalJSON encodes data into json.
func (t *Transaction) MarshalJSON() ([]byte, error) {
	type txdata struct {
		AccountNonce hexutil.Uint64  `json:"nonce"`
		Price        *hexutil.Big    `json:"gasPrice"`
		GasLimit     hexutil.Uint64  `json:"gas"`
		Recipient    *common.Address `json:"to"`
		From         *common.Address `json:"from"`
		Amount       *hexutil.Big    `json:"value"`
		Payload      hexutil.Bytes   `json:"input"`
		V            *hexutil.Big    `json:"v"`
		R            *hexutil.Big    `json:"r"`
		S            *hexutil.Big    `json:"s"`
		Hash         *common.Hash    `json:"hash"`
		ChainID      *hexutil.Big    `json:"chainId"`
	}
	var enc txdata
	enc.AccountNonce = hexutil.Uint64(t.Nonce())
	enc.Price = (*hexutil.Big)(t.GasPrice())
	enc.GasLimit = hexutil.Uint64(t.Gas())
	enc.Recipient = t.To()
	enc.Amount = (*hexutil.Big)(t.Value())
	enc.Payload = t.Data()

	V, R, S := t.RawSignatureValues()
	enc.V = (*hexutil.Big)(V)
	enc.R = (*hexutil.Big)(R)
	enc.S = (*hexutil.Big)(S)

	hash := t.Hash()
	enc.Hash = &hash
	enc.ChainID = (*hexutil.Big)(t.ChainId())

	return json.Marshal(&enc)
}

// BlockTransactions holds transaction within pending block.
type BlockTransactions struct {
	Transactions Transactions `json:"transactions"`
}
