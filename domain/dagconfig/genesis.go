// Copyright (c) 2014-2016 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package dagconfig

import (
	"math/big"

	"github.com/spectre-project/go-muhash"
	"github.com/spectre-project/spectred/domain/consensus/model/externalapi"
	"github.com/spectre-project/spectred/domain/consensus/utils/blockheader"
	"github.com/spectre-project/spectred/domain/consensus/utils/subnetworks"
	"github.com/spectre-project/spectred/domain/consensus/utils/transactionhelper"
)

var genesisTxOuts = []*externalapi.DomainTransactionOutput{}

var genesisTxPayload = []byte{
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // Blue score
	0x00, 0xE1, 0xF5, 0x05, 0x00, 0x00, 0x00, 0x00, // Subsidy
	0x00, 0x00, // Script version
	0x01,                                           // Varint
	0x00,                                           // OP-FALSE
	0x27, 0x18, 0x28, 0x18, 0x28, 0x45, 0x90, 0x45, // Euler's number = 2.718281828459045
}

// genesisCoinbaseTx is the coinbase transaction for the genesis blocks for
// the main network.
var genesisCoinbaseTx = transactionhelper.NewSubnetworkTransaction(0, []*externalapi.DomainTransactionInput{}, genesisTxOuts,
	&subnetworks.SubnetworkIDCoinbase, 0, genesisTxPayload)

// genesisHash is the hash of the first block in the block DAG for the main
// network (genesis block).
var genesisHash = externalapi.NewDomainHashFromByteArray(&[externalapi.DomainHashSize]byte{
	0x2A, 0xFA, 0x63, 0xE3, 0xAC, 0x16, 0x65, 0x62,
	0x97, 0xBE, 0xF7, 0x75, 0x23, 0x79, 0x91, 0xC3,
	0x9E, 0xED, 0x10, 0xF5, 0x23, 0x84, 0xEE, 0x9D,
	0x94, 0x20, 0x2C, 0x80, 0x1C, 0x76, 0xF5, 0x5D,
})

// genesisMerkleRoot is the hash of the first transaction in the genesis block
// for the main network.
var genesisMerkleRoot = externalapi.NewDomainHashFromByteArray(&[externalapi.DomainHashSize]byte{
	0x55, 0x29, 0xA6, 0xB3, 0xB8, 0x7F, 0xC2, 0x09,
	0x12, 0xA6, 0xE6, 0xD7, 0x9E, 0xFF, 0x9B, 0x92,
	0x49, 0xF2, 0x4F, 0xF9, 0xED, 0xDA, 0x4D, 0xEC,
	0x40, 0x59, 0xEF, 0x9E, 0xD7, 0xC5, 0xBD, 0xCB,
})

// genesisBlock defines the genesis block of the block DAG which serves as the
// public transaction ledger for the main network.
var genesisBlock = externalapi.DomainBlock{
	Header: blockheader.NewImmutableBlockHeader(
		0,
		[]externalapi.BlockLevelParents{},
		genesisMerkleRoot,
		&externalapi.DomainHash{},
		externalapi.NewDomainHashFromByteArray(muhash.EmptyMuHashHash.AsArray()),
		1714369615432,
		536999497, // Prime number
		271828,    // Euler's number
		0,
		0,
		big.NewInt(0),
		&externalapi.DomainHash{},
	),
	Transactions: []*externalapi.DomainTransaction{genesisCoinbaseTx},
}

var devnetGenesisTxOuts = []*externalapi.DomainTransactionOutput{}

var devnetGenesisTxPayload = []byte{
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // Blue score
	0x00, 0xE1, 0xF5, 0x05, 0x00, 0x00, 0x00, 0x00, // Subsidy
	0x00, 0x00, // Script version
	0x01,                                           // Varint
	0x00,                                           // OP-FALSE
	0x24, 0x14, 0x21, 0x35, 0x62, 0x37, 0x30, 0x95, // Silver ratio
}

// devnetGenesisCoinbaseTx is the coinbase transaction for the genesis blocks for
// the development network.
var devnetGenesisCoinbaseTx = transactionhelper.NewSubnetworkTransaction(0,
	[]*externalapi.DomainTransactionInput{}, devnetGenesisTxOuts,
	&subnetworks.SubnetworkIDCoinbase, 0, devnetGenesisTxPayload)

// devGenesisHash is the hash of the first block in the block DAG for the development
// network (genesis block).
var devnetGenesisHash = externalapi.NewDomainHashFromByteArray(&[externalapi.DomainHashSize]byte{
	0x6C, 0x34, 0x89, 0xBF, 0xB5, 0x92, 0xCA, 0x0A,
	0x0C, 0x12, 0xED, 0xB7, 0xAD, 0x86, 0x2D, 0x62,
	0x27, 0x92, 0x3E, 0xC2, 0xD2, 0x77, 0x7E, 0x0D,
	0xFD, 0x93, 0xF3, 0xC5, 0xB8, 0xA5, 0x5C, 0x35,
})

// devnetGenesisMerkleRoot is the hash of the first transaction in the genesis block
// for the devopment network.
var devnetGenesisMerkleRoot = externalapi.NewDomainHashFromByteArray(&[externalapi.DomainHashSize]byte{
	0x45, 0x7F, 0x6D, 0xF5, 0x76, 0x25, 0xCF, 0xC9,
	0x4A, 0x63, 0x16, 0x9E, 0xBA, 0xC8, 0xE1, 0x86,
	0xCF, 0x1B, 0x5F, 0x1E, 0xF6, 0x8D, 0x1A, 0xEF,
	0x3B, 0x8D, 0x3F, 0xFC, 0xC2, 0x6C, 0x01, 0xE4,
})

// devnetGenesisBlock defines the genesis block of the block DAG which serves as the
// public transaction ledger for the development network.
var devnetGenesisBlock = externalapi.DomainBlock{
	Header: blockheader.NewImmutableBlockHeader(
		0,
		[]externalapi.BlockLevelParents{},
		devnetGenesisMerkleRoot,
		&externalapi.DomainHash{},
		externalapi.NewDomainHashFromByteArray(muhash.EmptyMuHashHash.AsArray()),
		1713884849877,
		541034453, // Prime number
		241421,    // Silver ratio
		0,
		0,
		big.NewInt(0),
		&externalapi.DomainHash{},
	),
	Transactions: []*externalapi.DomainTransaction{devnetGenesisCoinbaseTx},
}

var simnetGenesisTxOuts = []*externalapi.DomainTransactionOutput{}

var simnetGenesisTxPayload = []byte{
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // Blue score
	0x00, 0xE1, 0xF5, 0x05, 0x00, 0x00, 0x00, 0x00, // Subsidy
	0x00, 0x00, // Script version
	0x01,                                           // Varint
	0x00,                                           // OP-FALSE
	0x54, 0x36, 0x56, 0x36, 0x56, 0x91, 0x80, 0x90, // Euler's number * 2 = 5.436563656918090
}

// simnetGenesisCoinbaseTx is the coinbase transaction for the simnet genesis block.
var simnetGenesisCoinbaseTx = transactionhelper.NewSubnetworkTransaction(0,
	[]*externalapi.DomainTransactionInput{}, simnetGenesisTxOuts,
	&subnetworks.SubnetworkIDCoinbase, 0, simnetGenesisTxPayload)

// simnetGenesisHash is the hash of the first block in the block DAG for
// the simnet (genesis block).
var simnetGenesisHash = externalapi.NewDomainHashFromByteArray(&[externalapi.DomainHashSize]byte{
	0x56, 0xBB, 0x87, 0xCF, 0x18, 0x77, 0x7B, 0x76,
	0x35, 0x8E, 0xEE, 0xF0, 0x20, 0xA9, 0x01, 0xCD,
	0xDD, 0xDC, 0x0B, 0xA4, 0x46, 0xC0, 0x99, 0x2D,
	0xE2, 0x7C, 0xC2, 0xA8, 0x9E, 0xC7, 0xA1, 0x30,
})

// simnetGenesisMerkleRoot is the hash of the first transaction in the genesis block
// for the development network.
var simnetGenesisMerkleRoot = externalapi.NewDomainHashFromByteArray(&[externalapi.DomainHashSize]byte{
	0x85, 0x81, 0x84, 0xD0, 0x98, 0x16, 0x40, 0x4F,
	0xD7, 0xD7, 0x96, 0xFB, 0xDE, 0x60, 0xAC, 0x4B,
	0x99, 0x29, 0xB9, 0x18, 0x63, 0x39, 0xDA, 0x23,
	0x08, 0x3C, 0xDF, 0xC3, 0x5F, 0x13, 0x8F, 0xC6,
})

// simnetGenesisBlock defines the genesis block of the block DAG which serves as the
// public transaction ledger for the development network.
var simnetGenesisBlock = externalapi.DomainBlock{
	Header: blockheader.NewImmutableBlockHeader(
		0,
		[]externalapi.BlockLevelParents{},
		simnetGenesisMerkleRoot,
		&externalapi.DomainHash{},
		externalapi.NewDomainHashFromByteArray(muhash.EmptyMuHashHash.AsArray()),
		1713885012324,
		543656363, // Prime number
		2,         // Two
		0,
		0,
		big.NewInt(0),
		&externalapi.DomainHash{},
	),
	Transactions: []*externalapi.DomainTransaction{simnetGenesisCoinbaseTx},
}

var testnetGenesisTxOuts = []*externalapi.DomainTransactionOutput{}

var testnetGenesisTxPayload = []byte{
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // Blue score
	0x00, 0xE1, 0xF5, 0x05, 0x00, 0x00, 0x00, 0x00, // Subsidy
	0x00, 0x00, // Script version
	0x01,                                           // Varint
	0x00,                                           // OP-FALSE
	0x31, 0x41, 0x59, 0x26, 0x53, 0x58, 0x97, 0x93, // Pi = 3.141592653589793
}

// testnetGenesisCoinbaseTx is the coinbase transaction for the testnet genesis block.
var testnetGenesisCoinbaseTx = transactionhelper.NewSubnetworkTransaction(0,
	[]*externalapi.DomainTransactionInput{}, testnetGenesisTxOuts,
	&subnetworks.SubnetworkIDCoinbase, 0, testnetGenesisTxPayload)

// testnetGenesisHash is the hash of the first block in the block DAG for the test
// network (genesis block).
var testnetGenesisHash = externalapi.NewDomainHashFromByteArray(&[externalapi.DomainHashSize]byte{
	0x48, 0x44, 0xDF, 0x54, 0x95, 0x72, 0x66, 0x0E,
	0xAF, 0xDC, 0x9A, 0xA0, 0xBC, 0x1D, 0x2B, 0xEE,
	0xB8, 0xCA, 0x14, 0x0A, 0x5B, 0x5D, 0x63, 0x15,
	0xDC, 0x41, 0xBA, 0x42, 0x9B, 0xD2, 0x44, 0x00,
})

// testnetGenesisMerkleRoot is the hash of the first transaction in the genesis block
// for testnet.
var testnetGenesisMerkleRoot = externalapi.NewDomainHashFromByteArray(&[externalapi.DomainHashSize]byte{
	0xC5, 0xAE, 0xEF, 0x98, 0xF3, 0xE4, 0xF2, 0xBA,
	0x2C, 0xB4, 0xAF, 0x00, 0xC1, 0x6F, 0xEC, 0x3D,
	0x59, 0x9A, 0xF8, 0x03, 0x4E, 0xE1, 0xE0, 0x15,
	0xBC, 0x20, 0xCA, 0x60, 0xC9, 0x3E, 0x99, 0x1C,
})

// testnetGenesisBlock defines the genesis block of the block DAG which serves as the
// public transaction ledger for testnet.
var testnetGenesisBlock = externalapi.DomainBlock{
	Header: blockheader.NewImmutableBlockHeader(
		0,
		[]externalapi.BlockLevelParents{},
		testnetGenesisMerkleRoot,
		&externalapi.DomainHash{},
		externalapi.NewDomainHashFromByteArray(muhash.EmptyMuHashHash.AsArray()),
		1713884672545,
		511699987, // Prime number
		314159,    // Pi number
		0,
		0,
		big.NewInt(0),
		&externalapi.DomainHash{},
	),
	Transactions: []*externalapi.DomainTransaction{testnetGenesisCoinbaseTx},
}
