package cuckoo

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/core/types"
)

func TestTestMode(t *testing.T) {
	head := &types.Header{
		Number:     big.NewInt(1),
		Difficulty: big.NewInt(100),
		Time:       big.NewInt(0),
	}

	cuckoo := NewTester()
	block, err := cuckoo.Seal(nil, types.NewBlockWithHeader(head), nil)
	if err != nil {
		t.Fatalf("failed to seal block: %v", err)
	}

	head.Nonce = types.EncodeNonce(block.Nonce())
	head.Solution = block.Solution()

	if err := cuckoo.VerifySeal(nil, head); err != nil {
		t.Fatalf("unexpected verification error: %v", err)
	}

}