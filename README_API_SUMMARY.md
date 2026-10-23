API summary

POST /v1/wallets
- body: { network: "eth" | "btc" }
- response: { network, address, private_key }

GET /v1/validate?network=&address=
- validates address syntax

GET /v1/price/:symbol
- symbol: btc, eth

All responses are JSON.
