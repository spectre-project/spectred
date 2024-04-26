/*
Package util provides spectre-specific convenience functions and types.

# Block Overview

A Block defines a spectre block that provides easier and more efficient
manipulation of raw blocks. It also memoizes hashes for the
block and its transactions on their first access so subsequent accesses don't
have to repeat the relatively expensive hashing operations.

# Tx Overview

A Tx defines a spectre transaction that provides more efficient manipulation of
raw transactions. It memoizes the hash for the transaction on its
first access so subsequent accesses don't have to repeat the relatively
expensive hashing operations.

# Address Overview

The Address interface provides an abstraction for a spectre address. While the
most common type is a pay-to-pubkey, spectre already supports others and
may well support more in the future. This package currently provides
implementations for the pay-to-pubkey, and pay-to-script-hash address
types.

To decode/encode an address:

	addrString := "spectre:qqj9fg59mptxkr9j0y53j5mwurcmda5mtza9n6v9pm9uj8h0wgk6uma5pvumr"
	defaultPrefix := util.Bech32PrefixSpectre
	addr, err := util.DecodeAddress(addrString, defaultPrefix)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(addr.EncodeAddress())
*/
package util
