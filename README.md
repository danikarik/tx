# TX

## Supported networks

- `mainnet`
- `ropsten`

## Endpoints

### Transactions by block

#### GET [https://tx-demo.herokuapp.com/{NETWORK}/etherscan/block/{BLOCK}](https://tx-demo.herokuapp.com/ropsten/etherscan/block/5478009)

Request

```sh
curl https://tx-demo.herokuapp.com/ropsten/etherscan/block/5478009
```

Response

```json
{
  "block": "0x539679",
  "result": [
    {
      "nonce": "0x269fa",
      "gasPrice": "0x2540be400",
      "gas": "0xf4240",
      "to": "0xb2395289b27aaf2011425cb0f84b338550605493",
      "from": null,
      "value": "0x0",
      "input": "0xc84cda0f00000000000000000000000000000000000000000000000000000000000000400000000000000000000000000000000000000000000000000000000217b9ff0000000000000000000000000000000000000000000000000000000000000000034555520000000000000000000000000000000000000000000000000000000000",
      "v": "0x1c",
      "r": "0x61e034c07504c8161d1a734e496c6986c8e138721596b83b566a9bb1eacd9266",
      "s": "0x10ad1f5bc65a2879bc6c510146309400f5668abaec9515512e302401ab8543d1",
      "hash": "0xbae9a7865235fccf70d7336352bf5c3a8c75aabee31e8d55e4234fc6e7e60a15",
      "chainId": "0x0"
    }
  ]
}
```

#### GET [https://tx-demo.herokuapp.com/ropsten/etherscan/transaction/{HASH}](https://tx-demo.herokuapp.com/ropsten/etherscan/transaction/0xbae9a7865235fccf70d7336352bf5c3a8c75aabee31e8d55e4234fc6e7e60a15)

Request

```sh
curl https://tx-demo.herokuapp.com/ropsten/etherscan/block/5478009
```

Response

```json
{
  "result": {
    "nonce": "0x269fa",
    "gasPrice": "0x2540be400",
    "gas": "0xf4240",
    "to": "0xb2395289b27aaf2011425cb0f84b338550605493",
    "from": null,
    "value": "0x0",
    "input": "0xc84cda0f00000000000000000000000000000000000000000000000000000000000000400000000000000000000000000000000000000000000000000000000217b9ff0000000000000000000000000000000000000000000000000000000000000000034555520000000000000000000000000000000000000000000000000000000000",
    "v": "0x1c",
    "r": "0x61e034c07504c8161d1a734e496c6986c8e138721596b83b566a9bb1eacd9266",
    "s": "0x10ad1f5bc65a2879bc6c510146309400f5668abaec9515512e302401ab8543d1",
    "hash": "0xbae9a7865235fccf70d7336352bf5c3a8c75aabee31e8d55e4234fc6e7e60a15",
    "chainId": "0x0"
  }
}
```
