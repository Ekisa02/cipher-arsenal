package wallet

import (
	"testing"
)

func TestGenerateEthereumKeypair(t *testing.T) {
	priv, addr, err := Generate("eth")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(priv) == 0 {
		t.Fatalf("private key empty")
	}
	if len(addr) != 42 || addr[:2] != "0x" {
		t.Fatalf("invalid eth address format: %s", addr)
	}
}

func TestGenerateBitcoinKeypair(t *testing.T) {
	priv, addr, err := Generate("btc")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(priv) == 0 {
		t.Fatalf("private key empty")
	}
	if len(addr) < 26 || len(addr) > 36 {
		t.Fatalf("unexpected btc address length: %s", addr)
	}
}
