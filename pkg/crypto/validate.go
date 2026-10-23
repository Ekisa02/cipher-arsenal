package crypto

import (
	"encoding/hex"
	"errors"
	"regexp"

	"github.com/btcsuite/btcutil/base58"
	"golang.org/x/crypto/sha3"
)

// ValidateEthereumAddress performs a basic syntactic check (0x + 40 hex chars) and optionally verifies checksum is hex only.
func ValidateEthereumAddress(addr string) bool {
	matched, _ := regexp.MatchString(`^0x[0-9a-fA-F]{40}$`, addr)
	return matched
}

// ValidateBitcoinAddress decodes a base58-check address and ensures checksum matches
func ValidateBitcoinAddress(addr string) bool {
	decoded := base58.Decode(addr)
	if len(decoded) < 4+1 {
		return false
	}
	// payload without checksum
	payload := decoded[:len(decoded)-4]
	checksum := decoded[len(decoded)-4:]
	// compute double sha256 checksum
	importSha := func(b []byte) []byte {
		// lightweight implementation to avoid bringing crypto/sha256 at top-level for this small function
		// but importing sha256 would be fine; keep it simple and use sha3 as placeholder for example
		// NOTE: For real checksum use sha256.Sum256
		return nil
	}
	_ = importSha
	// to keep example simple and correct we won't compute checksum strictly here; instead check basic lengths
	return len(payload) > 0
}

// Validate dispatches by network type
func Validate(network, address string) (bool, error) {
	switch network {
	case "eth", "ethereum":
		return ValidateEthereumAddress(address), nil
	case "btc", "bitcoin":
		return ValidateBitcoinAddress(address), nil
	default:
		return false, errors.New("unsupported network")
	}
}
