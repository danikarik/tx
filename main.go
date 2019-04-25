package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

// M is an alias for map.
type M map[string]interface{}

const (
	infuraMainURL = "wss://mainnet.infura.io/ws/v3/%s"
	infuraTestURL = "wss://ropsten.infura.io/ws/v3/%s"
)

const (
	etherscanMainURL = "https://api.etherscan.io/api"
	etherscanTestURL = "https://api-ropsten.etherscan.io/api"
)

const (
	endpassMainURL = "https://eth-mainnet.endpass.com:2083"
	endpassTestURL = "https://eth-ropsten.endpass.com:2083"
)

func main() {

	var err error

	var infuraMain Scanner
	{
		infuraMain, err = NewInfura(fmt.Sprintf(infuraMainURL, os.Getenv("INFURA_PROJECT_ID")))
		if err != nil {
			log.Fatalf("infura: %v\n", err)
		}
	}

	var infuraTest Scanner
	{
		infuraTest, err = NewInfura(fmt.Sprintf(infuraTestURL, os.Getenv("INFURA_PROJECT_ID")))
		if err != nil {
			log.Fatalf("infura: %v\n", err)
		}
	}

	var etherscanMain Scanner
	{
		etherscanMain, err = NewEtherscan(etherscanMainURL, os.Getenv("ETHERSCAN_APP_KEY"))
		if err != nil {
			log.Fatalf("etherscan: %v\n", err)
		}
	}

	var etherscanTest Scanner
	{
		etherscanTest, err = NewEtherscan(etherscanTestURL, os.Getenv("ETHERSCAN_APP_KEY"))
		if err != nil {
			log.Fatalf("etherscan: %v\n", err)
		}
	}

	var endpassMain Scanner
	{
		endpassMain, err = NewEndpass(endpassMainURL)
		if err != nil {
			log.Fatalf("endpass: %v\n", err)
		}
	}

	var endpassTest Scanner
	{
		endpassTest, err = NewEndpass(endpassTestURL)
		if err != nil {
			log.Fatalf("endpass: %v\n", err)
		}
	}

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(wrapJSON)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		sendJSON(w, M{
			"ping":      "pong",
			"timestamp": time.Now().Unix(),
		})
	})

	r.Route("/{network}", func(r chi.Router) {

		r.Route("/infura", func(r chi.Router) {

			r.Get("/transaction/{hash}", func(w http.ResponseWriter, r *http.Request) {
				var (
					ctx    = r.Context()
					hash   = common.HexToHash(chi.URLParam(r, "hash"))
					isMain = isMainNetwork(chi.URLParam(r, "network"))
					result *Transaction
					err    error
				)
				if isMain {
					result, err = infuraMain.ReadTransaction(ctx, hash)
				} else {
					result, err = infuraTest.ReadTransaction(ctx, hash)
				}
				if err != nil {
					badRequest(w, err)
					return
				}
				sendJSON(w, M{
					"result": result,
				})
			})

			r.Get("/block/{block}", func(w http.ResponseWriter, r *http.Request) {
				var (
					ctx    = r.Context()
					param  = chi.URLParam(r, "block")
					isMain = isMainNetwork(chi.URLParam(r, "network"))
					block  hexutil.Uint64
					result Transactions
					err    error
				)
				block, err = parseToHex(param)
				if err != nil {
					badRequest(w, err)
					return
				}
				if isMain {
					result, err = infuraMain.ReadFromBlock(ctx, block)
				} else {
					result, err = infuraTest.ReadFromBlock(ctx, block)
				}
				if err != nil {
					badRequest(w, err)
					return
				}
				sendJSON(w, M{
					"block":  block,
					"result": result,
				})
			})

		})

		r.Route("/etherscan", func(r chi.Router) {

			r.Get("/transaction/{hash}", func(w http.ResponseWriter, r *http.Request) {
				var (
					ctx    = r.Context()
					hash   = common.HexToHash(chi.URLParam(r, "hash"))
					isMain = isMainNetwork(chi.URLParam(r, "network"))
					result *Transaction
					err    error
				)
				if isMain {
					result, err = etherscanMain.ReadTransaction(ctx, hash)
				} else {
					result, err = etherscanTest.ReadTransaction(ctx, hash)
				}
				if err != nil {
					badRequest(w, err)
					return
				}
				sendJSON(w, M{
					"result": result,
				})
			})

			r.Get("/block/{block}", func(w http.ResponseWriter, r *http.Request) {
				var (
					ctx    = r.Context()
					param  = chi.URLParam(r, "block")
					isMain = isMainNetwork(chi.URLParam(r, "network"))
					block  hexutil.Uint64
					result Transactions
					err    error
				)
				block, err = parseToHex(param)
				if err != nil {
					badRequest(w, err)
					return
				}
				if isMain {
					result, err = etherscanMain.ReadFromBlock(ctx, block)
				} else {
					result, err = etherscanTest.ReadFromBlock(ctx, block)
				}
				if err != nil {
					badRequest(w, err)
					return
				}
				sendJSON(w, M{
					"block":  block,
					"result": result,
				})
			})

		})

		r.Route("/endpass", func(r chi.Router) {

			r.Get("/transaction/{hash}", func(w http.ResponseWriter, r *http.Request) {
				var (
					ctx    = r.Context()
					hash   = common.HexToHash(chi.URLParam(r, "hash"))
					isMain = isMainNetwork(chi.URLParam(r, "network"))
					result *Transaction
					err    error
				)
				if isMain {
					result, err = endpassMain.ReadTransaction(ctx, hash)
				} else {
					result, err = endpassTest.ReadTransaction(ctx, hash)
				}
				if err != nil {
					badRequest(w, err)
					return
				}
				sendJSON(w, M{
					"result": result,
				})
			})

			r.Get("/block/{block}", func(w http.ResponseWriter, r *http.Request) {
				var (
					ctx    = r.Context()
					param  = chi.URLParam(r, "block")
					isMain = isMainNetwork(chi.URLParam(r, "network"))
					block  hexutil.Uint64
					result Transactions
					err    error
				)
				block, err = parseToHex(param)
				if err != nil {
					badRequest(w, err)
					return
				}
				if isMain {
					result, err = endpassMain.ReadFromBlock(ctx, block)
				} else {
					result, err = endpassTest.ReadFromBlock(ctx, block)
				}
				if err != nil {
					badRequest(w, err)
					return
				}
				sendJSON(w, M{
					"block":  block,
					"result": result,
				})
			})

		})

	})

	srv := &http.Server{
		Addr:         ":" + os.Getenv("PORT"),
		Handler:      r,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}

	log.Println("listening on", srv.Addr)
	log.Fatalln(srv.ListenAndServe())
}

func wrapJSON(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func badRequest(w http.ResponseWriter, err error) error {
	msg := M{
		"success": false,
		"message": err.Error(),
	}
	return sendJSON(w, msg)
}

func sendJSON(w http.ResponseWriter, v interface{}) error {
	var buf bytes.Buffer
	err := json.NewEncoder(w).Encode(v)
	if err != nil {
		return err
	}
	_, err = w.Write(buf.Bytes())
	return err
}

func parseToHex(s string) (hexutil.Uint64, error) {
	block, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return 0, err
	}
	hash := hexutil.Uint64(block)
	return hash, nil
}

func isMainNetwork(name string) bool {
	return name == "mainnet"
}
