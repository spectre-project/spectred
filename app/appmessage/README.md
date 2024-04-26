# wire

[![ISC License](http://img.shields.io/badge/license-ISC-blue.svg)](https://choosealicense.com/licenses/isc/)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg)](http://godoc.org/github.com/spectre-project/spectred/wire)

Package wire implements the spectre wire protocol.

## Spectre Message Overview

The spectre protocol consists of exchanging messages between peers.
Each message is preceded by a header which identifies information
about it such as which spectre network it is a part of, its type, how
big it is, and a checksum to verify validity. All encoding and
decoding of message headers is handled by this package.

To accomplish this, there is a generic interface for spectre messages
named `Message` which allows messages of any type to be read, written,
or passed around through channels, functions, etc. In addition,
concrete implementations of most all spectre messages are provided.
All of the details of marshalling and unmarshalling to and from the
wire using spectre encoding are handled so the  caller doesn't have
to concern themselves with the specifics.

## Reading Messages Example

In order to unmarshal spectre messages from the wire, use the
`ReadMessage` function. It accepts any `io.Reader`, but typically
this will be a `net.Conn` to a remote node running a spectre peer.
Example syntax is:

```Go
// Use the most recent protocol version supported by the package and the
// main spectre network.
pver := wire.ProtocolVersion
spectrenet := wire.Mainnet

// Reads and validates the next spectre message from conn using the
// protocol version pver and the spectre network spectrenet. The returns
// are a appmessage.Message, a []byte which contains the unmarshalled
// raw payload, and a possible error.
msg, rawPayload, err := wire.ReadMessage(conn, pver, spectrenet)
if err != nil {
	// Log and handle the error
}
```

See the package documentation for details on determining the message
type.

## Writing Messages Example

In order to marshal spectre messages to the wire, use the
`WriteMessage` function. It accepts any `io.Writer`, but typically
this will be a `net.Conn` to a remote node running a spectre peer.
Example syntax to request addresses from a remote peer is:

```Go
// Use the most recent protocol version supported by the package and the
// main bitcoin network.
pver := wire.ProtocolVersion
spectrenet := wire.Mainnet

// Create a new getaddr spectre message.
msg := wire.NewMsgGetAddr()

// Writes a spectre message msg to conn using the protocol version
// pver, and the spectre network spectrenet. The return is a possible
// error.
err := wire.WriteMessage(conn, msg, pver, spectrenet)
if err != nil {
	// Log and handle the error
}
```
