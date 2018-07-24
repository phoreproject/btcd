// Copyright (c) 2014-2016 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package chaincfg

import (
	"bytes"
	"testing"

	"github.com/davecgh/go-spew/spew"
)

// TestGenesisBlock tests the genesis block of the main network for validity by
// checking the encoded bytes and hashes.
func TestGenesisBlock(t *testing.T) {
	// Encode the genesis block to raw bytes.
	var buf bytes.Buffer
	err := MainNetParams.GenesisBlock.Serialize(&buf)
	if err != nil {
		t.Fatalf("TestGenesisBlock: %v", err)
	}

	// Ensure the encoded block matches the expected bytes.
	if !bytes.Equal(buf.Bytes(), genesisBlockBytes) {
		t.Fatalf("TestGenesisBlock: Genesis block does not appear valid - "+
			"got %v, want %v", spew.Sdump(buf.Bytes()),
			spew.Sdump(genesisBlockBytes))
	}

	// Check hash of the block against expected hash.
	hash := MainNetParams.GenesisBlock.BlockHash()
	if !MainNetParams.GenesisHash.IsEqual(&hash) {
		t.Fatalf("TestGenesisBlock: Genesis block hash does not "+
			"appear valid - got %v, want %v", spew.Sdump(hash),
			spew.Sdump(MainNetParams.GenesisHash))
	}
}

// TestRegTestGenesisBlock tests the genesis block of the regression test
// network for validity by checking the encoded bytes and hashes.
// func TestRegTestGenesisBlock(t *testing.T) {
// 	// Encode the genesis block to raw bytes.
// 	var buf bytes.Buffer
// 	err := RegressionNetParams.GenesisBlock.Serialize(&buf)
// 	if err != nil {
// 		t.Fatalf("TestRegTestGenesisBlock: %v", err)
// 	}

// 	// Ensure the encoded block matches the expected bytes.
// 	if !bytes.Equal(buf.Bytes(), regTestGenesisBlockBytes) {
// 		t.Fatalf("TestRegTestGenesisBlock: Genesis block does not "+
// 			"appear valid - got %v, want %v",
// 			spew.Sdump(buf.Bytes()),
// 			spew.Sdump(regTestGenesisBlockBytes))
// 	}

// 	// Check hash of the block against expected hash.
// 	hash := RegressionNetParams.GenesisBlock.BlockHash()
// 	if !RegressionNetParams.GenesisHash.IsEqual(&hash) {
// 		t.Fatalf("TestRegTestGenesisBlock: Genesis block hash does "+
// 			"not appear valid - got %v, want %v", spew.Sdump(hash),
// 			spew.Sdump(RegressionNetParams.GenesisHash))
// 	}
// }

// TestTestNet3GenesisBlock tests the genesis block of the test network (version
// 3) for validity by checking the encoded bytes and hashes.
// func TestTestNet3GenesisBlock(t *testing.T) {
// 	// Encode the genesis block to raw bytes.
// 	var buf bytes.Buffer
// 	err := TestNet3Params.GenesisBlock.Serialize(&buf)
// 	if err != nil {
// 		t.Fatalf("TestTestNet3GenesisBlock: %v", err)
// 	}

// 	// Ensure the encoded block matches the expected bytes.
// 	if !bytes.Equal(buf.Bytes(), testNet3GenesisBlockBytes) {
// 		t.Fatalf("TestTestNet3GenesisBlock: Genesis block does not "+
// 			"appear valid - got %v, want %v",
// 			spew.Sdump(buf.Bytes()),
// 			spew.Sdump(testNet3GenesisBlockBytes))
// 	}

// 	// Check hash of the block against expected hash.
// 	hash := TestNet3Params.GenesisBlock.BlockHash()
// 	if !TestNet3Params.GenesisHash.IsEqual(&hash) {
// 		t.Fatalf("TestTestNet3GenesisBlock: Genesis block hash does "+
// 			"not appear valid - got %v, want %v", spew.Sdump(hash),
// 			spew.Sdump(TestNet3Params.GenesisHash))
// 	}
// }

// TestSimNetGenesisBlock tests the genesis block of the simulation test network
// for validity by checking the encoded bytes and hashes.
// func TestSimNetGenesisBlock(t *testing.T) {
// 	// Encode the genesis block to raw bytes.
// 	var buf bytes.Buffer
// 	err := SimNetParams.GenesisBlock.Serialize(&buf)
// 	if err != nil {
// 		t.Fatalf("TestSimNetGenesisBlock: %v", err)
// 	}

// 	// Ensure the encoded block matches the expected bytes.
// 	if !bytes.Equal(buf.Bytes(), simNetGenesisBlockBytes) {
// 		t.Fatalf("TestSimNetGenesisBlock: Genesis block does not "+
// 			"appear valid - got %v, want %v",
// 			spew.Sdump(buf.Bytes()),
// 			spew.Sdump(simNetGenesisBlockBytes))
// 	}

// 	// Check hash of the block against expected hash.
// 	hash := SimNetParams.GenesisBlock.BlockHash()
// 	if !SimNetParams.GenesisHash.IsEqual(&hash) {
// 		t.Fatalf("TestSimNetGenesisBlock: Genesis block hash does "+
// 			"not appear valid - got %v, want %v", spew.Sdump(hash),
// 			spew.Sdump(SimNetParams.GenesisHash))
// 	}
// }

// genesisBlockBytes are the wire encoded bytes for the genesis block of the
// main network as of protocol version 70003.
var genesisBlockBytes = []byte{
	0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* |........| */
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* |........| */
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* |........| */
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* |........| */
	0x00, 0x00, 0x00, 0x00, 0xa3, 0xb3, 0x36, 0x07, /* |......6.| */
	0xcb, 0x85, 0x1a, 0x7c, 0xfc, 0xde, 0x34, 0x8f, /* |...|..4.| */
	0x54, 0x53, 0x8a, 0x20, 0x85, 0xfc, 0xe7, 0x95, /* |TS. ....| */
	0xd3, 0x9d, 0xd8, 0xfe, 0x2c, 0x95, 0x45, 0x7a, /* |....,.Ez| */
	0x13, 0x77, 0x41, 0x89, 0x60, 0xe8, 0xb7, 0x59, /* |.wA.`..Y| */
	0xff, 0xff, 0x7f, 0x20, 0x39, 0x30, 0x00, 0x00, /* |... 90..| */
	0x01, 0x01, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, /* |........| */
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* |........| */
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* |........| */
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* |........| */
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xff, 0xff, /* |........| */
	0xff, 0xff, 0x19, 0x04, 0xff, 0xff, 0x00, 0x1d, /* |........| */
	0x01, 0x04, 0x11, 0x31, 0x32, 0x20, 0x53, 0x65, /* |...12 Se| */
	0x70, 0x74, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x20, /* |ptember | */
	0x32, 0x30, 0x31, 0x37, 0xff, 0xff, 0xff, 0xff, /* |2017....| */
	0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* |........| */
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* |......| */
}
