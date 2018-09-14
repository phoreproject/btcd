package zerocoin_test

import (
	"bytes"
	"encoding/hex"
	"math/big"
	"testing"

	"github.com/phoreproject/btcd/zerocoin"
)

func TestSerializeBigNum(t *testing.T) {
	out := zerocoin.SerializeBigNum(big.NewInt(100000))
	// 0x186a0 = 3 bytes
	expected, _ := hex.DecodeString("000000030186a0")
	if bytes.Compare(out, expected) != 0 {
		t.Errorf("Could not serialize big number correctly. Got: %x, Expected: %x", out, expected)
		return
	}
}
