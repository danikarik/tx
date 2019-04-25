package main

import (
	"context"
	"net/url"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/rpc"
)

// NewEndpass creates a new instance of endpass scanner.
func NewEndpass(src string) (Scanner, error) {
	url, err := url.Parse(src)
	if err != nil {
		return nil, err
	}
	client, err := rpc.DialContext(context.Background(), url.String())
	if err != nil {
		return nil, err
	}
	return &infura{client}, nil
}

type endpass struct {
	client *rpc.Client
}

func (e *endpass) ReadFromBlock(ctx context.Context, block hexutil.Uint64) (Transactions, error) {
	var result BlockTransactions
	err := e.client.CallContext(ctx, &result, "eth_getBlockByNumber", block, true)
	if err != nil {
		return nil, err
	}
	return result.Transactions, nil
}

func (e *endpass) ReadTransaction(ctx context.Context, hash common.Hash) (*Transaction, error) {
	var result *Transaction
	err := e.client.CallContext(ctx, &result, "eth_getTransactionByHash", hash)
	if err != nil {
		return nil, err
	}
	return result, nil
}
