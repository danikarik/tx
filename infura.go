package main

import (
	"context"
	"net/url"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/rpc"
)

// NewInfura creates a new instance of infura scanner.
func NewInfura(src string) (Scanner, error) {
	url, err := url.Parse(src)
	if err != nil {
		return nil, err
	}
	client, err := rpc.DialWebsocket(context.Background(), url.String(), "")
	if err != nil {
		return nil, err
	}
	return &infura{client}, nil
}

type infura struct {
	client *rpc.Client
}

func (i *infura) ReadFromBlock(ctx context.Context, block hexutil.Uint64) (Transactions, error) {
	var result BlockTransactions
	err := i.client.CallContext(ctx, &result, "eth_getBlockByNumber", block, true)
	if err != nil {
		return nil, err
	}
	return result.Transactions, nil
}

func (i *infura) ReadTransaction(ctx context.Context, hash common.Hash) (*Transaction, error) {
	var result *Transaction
	err := i.client.CallContext(ctx, &result, "eth_getTransactionByHash", hash)
	if err != nil {
		return nil, err
	}
	return result, nil
}
