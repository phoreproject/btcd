package zerocoin_test

import (
	"encoding/hex"
	"math/big"
	"math/rand"
	"testing"

	"github.com/phoreproject/btcd/zerocoin"
)

func TestBigConversion(t *testing.T) {
	x := big.NewInt(int64(rand.Uint64()))
	xHashed, err := zerocoin.BigToHash(x)
	if err != nil {
		t.Error(err)
	}

	xReversed := zerocoin.HashToBig(xHashed)
	if xReversed.Cmp(x) != 0 {
		t.Errorf("HashToBig(BigToHash(%x)) == %x != %x", x, xReversed, x)
	}
}

func TestCalculateParams(t *testing.T) {
	zerocoinModulusBytes, _ := hex.DecodeString("c7970ceedcc3b0754490201a7aa613cd73911081c790f5f1a8726f463550bb5b7ff0db8e1ea1189ec72f93d1650011bd721aeeacc2acde32a04107f0648c2813a31f5b0b7765ff8b44b4b6ffc93384b646eb09c7cf5e8592d40ea33c80039f35b4f14a04b51f7bfd781be4d1673164ba8eb991c2c4d730bbbe35f592bdef524af7e8daefd26c66fc02c479af89d64d373f442709439de66ceb955f3ea37d5159f6135809f85334b5cb1813addc80cd05609f10ac6a95ad65872c909525bdad32bc729592642920f24c61dc5b3c3b7923e56b16a4d9d373d8721f24a3fc0f1b3131f55615172866bccc30f95054c824e733a5eb6817f7bc16399d48c6361cc7e5")
	zerocoinModulus := new(big.Int).SetBytes(zerocoinModulusBytes)
	p, err := zerocoin.NewZerocoinParams(zerocoinModulus, 80)
	if err != nil {
		t.Error(err)
		return
	}

	if !p.AccumulatorParams.AccumulatorModulus.ProbablyPrime(80) {
		t.Error("Modulo is not prime")
	}
}
