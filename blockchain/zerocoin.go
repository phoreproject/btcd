package blockchain

import (
	"github.com/phoreproject/btcd/chaincfg/chainhash"
	"github.com/phoreproject/btcd/zerocoin"
)

func (b *BlockChain) CalculateAccumulatorCheckpoint(height int32) (*zerocoin.Accumulators, chainhash.Hash, bool, error) {
	if height < b.chainParams.ZerocoinStartHeight {
		return nil, *zeroHash, true, nil
	}

	if height%10 != 0 {
		return nil, b.bestChain.NodeByHeight(height - 1).accumulatorCheckpoint, true, nil
	}
}
