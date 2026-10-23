# Curl examples

Generate an Ethereum wallet:

curl -s -X POST http://localhost:8080/v1/wallets -H "Content-Type: application/json" -d '{"network":"eth"}' | jq

Validate address (query params):

curl -s "http://localhost:8080/v1/validate?network=eth&address=0x0123..." | jq

Get price:

curl -s http://localhost:8080/v1/price/btc | jq
