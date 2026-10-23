# Cipher Arsenal

Cipher Arsenal is a modular, production-minded Go service providing common blockchain and crypto utilities: deterministic wallet generation (Ethereum & Bitcoin), address validation, transaction serialization helpers, and lightweight price lookup — backed by Gin for HTTP and GORM for persistence.

Key features
- Generate wallets (Ethereum, Bitcoin) with example private key output (for dev/testing only).
- Validate Ethereum and Bitcoin addresses.
- Simple transaction serialization helper (mock/skeleton to be extended per-chain).
- Price lookup adapter (CoinGecko-based) with simple caching model.
- GORM models and migrations (SQLite by default, changeable via env DATABASE_URL).
- HTTP API built on Gin with clear routes and handlers.

Architecture

ASCII overview:

  +-----------------+       +--------------------+       +-----------+
  |  HTTP (Gin)     | <---- |  Handlers (pkg/api)| <---- |  Services |
  +-----------------+       +--------------------+       +-----------+
           |                         |                       |
           v                         v                       v
  +-----------------+       +--------------------+      +----------------+
  | cmd/server/main | ----> | internal/store     | ---> | Database (SQLite)|
  +-----------------+       +--------------------+      +----------------+
                                  ^   ^
                                  |   |
                              +-------+-----+
                              | pkg/wallet   |
                              | pkg/crypto   |
                              | pkg/price    |
                              +-------------+


Directories
- cmd/server: application entrypoint and wiring
- internal/config: environment config loader
- internal/store: GORM models and DB helpers
- pkg/wallet: wallet generation utilities
- pkg/crypto: address validation and helpers
- pkg/price: price retrieval adapter
- docs: usage examples and curl snippets

Getting started (development)

Prerequisites
- Go 1.20+ installed
- (Optional) Docker for containerized runs

Quick run

1. Clone

   git clone <this-repo>
   cd cipher-arsenal

2. Build & run locally

   go build ./cmd/server
   ./server

   By default the server listens on :8080 and uses SQLite file cipher_arsenal.db in the working directory.

Environment variables
- PORT: HTTP port (default: 8080)
- DATABASE_URL: DSN for GORM. Default: file cipher_arsenal.db using sqlite.
- CACHE_TTL: price cache TTL in seconds (default: 60)

Docker

  docker build -t cipher-arsenal:latest .
  docker run -p 8080:8080 -e PORT=8080 cipher-arsenal:latest

API (examples)

- POST /v1/wallets
  Body: { "network": "eth" }
  Generates a wallet and returns address and private key (development only).

- GET /v1/validate?network=eth&address=0x...
  Returns whether the supplied address is syntactically valid.

- GET /v1/price/:symbol
  Returns USD price for symbols like btc or eth (uses CoinGecko).

Curl examples
See docs/curl_examples.md for practical curl commands.

Security note
This project returns private keys for generated wallets for development convenience only. Do NOT expose these endpoints in production without strong access controls and hardware-backed key management. The repo is intended as a utilities reference and starting point.

Extending
- Add chain-specific transaction serializers in pkg/tx
- Plug external DBs by changing GORM driver in internal/store
- Add authentication and rate-limiting at API layer

Testing

Run unit tests with:

  go test ./...

License

MIT
