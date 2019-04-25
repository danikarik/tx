package main

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// NewEtherscan creates a new instance of etherscan.
func NewEtherscan(src, key string) (Scanner, error) {
	url, err := url.Parse(src)
	if err != nil {
		return nil, err
	}
	return &etherscan{
		baseURL: url,
		apiKey:  key,
		client:  &http.Client{},
	}, nil
}

type etherscan struct {
	baseURL *url.URL
	apiKey  string
	client  *http.Client
}

type pair struct {
	key string
	val string
}

func (e *etherscan) getURL(pairs ...pair) string {
	u := e.baseURL
	for _, p := range pairs {
		q := e.baseURL.Query()
		q.Set(p.key, p.val)
		u.RawQuery = q.Encode()
	}
	return u.String()
}

func (e *etherscan) sendRequest(url string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	res, err := e.client.Do(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (e *etherscan) ReadFromBlock(ctx context.Context, block hexutil.Uint64) (Transactions, error) {
	url := e.getURL(
		pair{"module", "proxy"},
		pair{"action", "eth_getBlockByNumber"},
		pair{"tag", block.String()},
		pair{"boolean", "true"},
		pair{"apiKey", e.apiKey},
	)

	res, err := e.sendRequest(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var rpcJSON struct {
		Version string `json:"jsonrpc"`
		Result  struct {
			Transactions Transactions `json:"transactions"`
		} `json:"result"`
	}

	err = json.NewDecoder(res.Body).Decode(&rpcJSON)
	if err != nil {
		return nil, err
	}

	return rpcJSON.Result.Transactions, nil
}

func (e *etherscan) ReadTransaction(ctx context.Context, hash common.Hash) (*Transaction, error) {
	url := e.getURL(
		pair{"module", "proxy"},
		pair{"action", "eth_getTransactionByHash"},
		pair{"txhash", hash.String()},
		pair{"apiKey", e.apiKey},
	)

	res, err := e.sendRequest(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var rpcJSON struct {
		ID      uint64       `json:"id"`
		Version string       `json:"jsonrpc"`
		Result  *Transaction `json:"result"`
	}

	err = json.NewDecoder(res.Body).Decode(&rpcJSON)
	if err != nil {
		return nil, err
	}

	return rpcJSON.Result, nil
}
