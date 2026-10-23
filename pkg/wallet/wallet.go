package wallet

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"

	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcutil/base58"
	"golang.org/x/crypto/ripemd160"
	"golang.org/x/crypto/sha3"
)

// GenerateEthereumKeyPair returns a private key hex and an ethereum address (0x...)
func GenerateEthereumKeyPair() (privHex string, address string, err error) {
	priv, err := btcec.NewPrivateKey()
	if err != nil {
		return "", "", err
	}
	privBytes := priv.Serialize()
	pub := priv.PubKey().SerializeUncompressed()
	// Keccak256 hash of the public key (skip 0x04 prefix)
	h := sha3.NewLegacyKeccak256()
	h.Write(pub[1:])
	hh := h.Sum(nil)
	addr := hh[len(hh)-20:]
	return hex.EncodeToString(privBytes), "0x" + hex.EncodeToString(addr), nil
}

// GenerateBitcoinKeyPair returns a WIF-like private key hex and a P2PKH address
func GenerateBitcoinKeyPair() (privHex string, address string, err error) {
	priv, err := btcec.NewPrivateKey()
	if err != nil {
		return "", "", err
	}
	privBytes := priv.Serialize()
	pub := priv.PubKey().SerializeCompressed()
	// HASH160 (RIPEMD160(SHA256(pubkey)))
	h := sha256.Sum256(pub)
	r := ripemd160.New()
	r.Write(h[:])
	hash160 := r.Sum(nil)
	// Prepend version byte 0x00 for mainnet P2PKH
	payload := append([]byte{0x00}, hash160...)
	checksum := sha256.Sum256(payload)
	checksum = sha256.Sum256(checksum[:])
	final := append(payload, checksum[0:4]...)
	addr := base58.Encode(final)
	return hex.EncodeToString(privBytes), addr, nil
}

// ErrInvalidNetwork indicates unsupported network parameter
var ErrInvalidNetwork = errors.New("unsupported network")

// Generate generates keys for network: "eth" or "btc"
func Generate(network string) (privHex string, address string, err error) {
	switch network {
	case "eth", "ethereum":
		return GenerateEthereumKeyPair()
	case "btc", "bitcoin":
		return GenerateBitcoinKeyPair()
	default:
		return "", "", ErrInvalidNetwork
	}
}
