package main

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// Scanner reads transactions from specified source.
type Scanner interface {
	ReadTransaction(ctx context.Context, hash common.Hash) (*Transaction, error)
	ReadFromBlock(ctx context.Context, block hexutil.Uint64) (Transactions, error)
}
