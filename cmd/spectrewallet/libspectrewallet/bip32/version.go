package bip32

import "github.com/pkg/errors"

// BitcoinMainnetPrivate is the version that is used for
// bitcoin mainnet bip32 private extended keys.
// Ecnodes to xprv in base58.
var BitcoinMainnetPrivate = [4]byte{
	0x04,
	0x88,
	0xad,
	0xe4,
}

// BitcoinMainnetPublic is the version that is used for
// bitcoin mainnet bip32 public extended keys.
// Ecnodes to xpub in base58.
var BitcoinMainnetPublic = [4]byte{
	0x04,
	0x88,
	0xb2,
	0x1e,
}

// SpectreMainnetPrivate is the version that is used for
// spectre mainnet bip32 private extended keys.
// Ecnodes to xprv in base58.
var SpectreMainnetPrivate = [4]byte{
	0x03,
	0x8f,
	0x2e,
	0xf4,
}

// SpectreMainnetPublic is the version that is used for
// spectre mainnet bip32 public extended keys.
// Ecnodes to kpub in base58.
var SpectreMainnetPublic = [4]byte{
	0x03,
	0x8f,
	0x33,
	0x2e,
}

// SpectreTestnetPrivate is the version that is used for
// spectre testnet bip32 public extended keys.
// Ecnodes to ktrv in base58.
var SpectreTestnetPrivate = [4]byte{
	0x03,
	0x90,
	0x9e,
	0x07,
}

// SpectreTestnetPublic is the version that is used for
// spectre testnet bip32 public extended keys.
// Ecnodes to ktub in base58.
var SpectreTestnetPublic = [4]byte{
	0x03,
	0x90,
	0xa2,
	0x41,
}

// SpectreDevnetPrivate is the version that is used for
// spectre devnet bip32 public extended keys.
// Ecnodes to kdrv in base58.
var SpectreDevnetPrivate = [4]byte{
	0x03,
	0x8b,
	0x3d,
	0x80,
}

// SpectreDevnetPublic is the version that is used for
// spectre devnet bip32 public extended keys.
// Ecnodes to xdub in base58.
var SpectreDevnetPublic = [4]byte{
	0x03,
	0x8b,
	0x41,
	0xba,
}

// SpectreSimnetPrivate is the version that is used for
// spectre simnet bip32 public extended keys.
// Ecnodes to ksrv in base58.
var SpectreSimnetPrivate = [4]byte{
	0x03,
	0x90,
	0x42,
	0x42,
}

// SpectreSimnetPublic is the version that is used for
// spectre simnet bip32 public extended keys.
// Ecnodes to xsub in base58.
var SpectreSimnetPublic = [4]byte{
	0x03,
	0x90,
	0x46,
	0x7d,
}

func toPublicVersion(version [4]byte) ([4]byte, error) {
	switch version {
	case BitcoinMainnetPrivate:
		return BitcoinMainnetPublic, nil
	case SpectreMainnetPrivate:
		return SpectreMainnetPublic, nil
	case SpectreTestnetPrivate:
		return SpectreTestnetPublic, nil
	case SpectreDevnetPrivate:
		return SpectreDevnetPublic, nil
	case SpectreSimnetPrivate:
		return SpectreSimnetPublic, nil
	}

	return [4]byte{}, errors.Errorf("unknown version %x", version)
}

func isPrivateVersion(version [4]byte) bool {
	switch version {
	case BitcoinMainnetPrivate:
		return true
	case SpectreMainnetPrivate:
		return true
	case SpectreTestnetPrivate:
		return true
	case SpectreDevnetPrivate:
		return true
	case SpectreSimnetPrivate:
		return true
	}

	return false
}
