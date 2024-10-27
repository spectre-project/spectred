# DEPRECATED

**IMPORTANT NOTICE:**
The Go version of Spectre has been replaced by the stable Rust version. We strongly recommend switching to the new Rust version as soon as possible.

**Link to the Rust version:** [https://github.com/spectre-project/rusty-spectre](https://github.com/spectre-project/rusty-spectre)

**PLEASE NOTE:**

- Bugs and feature requests for the Go version will no longer be addressed. If you encounter any issues, please reproduce them in the Rust version and report them in the appropriate repository.
- Any pull requests or issues opened in this repository will be closed without consideration, except those related to `spectrewallet`, which is still being maintained. For all other cases, please use the [Rust implementation](https://github.com/spectre-project/rusty-spectre).

# Spectred

[![Build Status](https://github.com/spectre-project/spectred/actions/workflows/tests.yaml/badge.svg)](https://github.com/spectre-project/spectred/actions/workflows/tests.yaml)
[![GitHub release](https://img.shields.io/github/v/release/spectre-project/spectred.svg)](https://github.com/spectre-project/spectred/releases)
[![GitHub license](https://img.shields.io/github/license/spectre-project/spectred.svg)](https://github.com/spectre-project/spectred/blob/main/LICENSE)
[![GitHub downloads](https://img.shields.io/github/downloads/spectre-project/spectred/total.svg)](https://github.com/spectre-project/spectred/releases)
[![Join the Spectre Discord Server](https://img.shields.io/discord/1233113243741061240.svg?label=&logo=discord&logoColor=ffffff&color=5865F2)](https://discord.com/invite/FZPYpwszcF)

Spectred was the reference full node Spectre implementation written in
Go (golang). It is a [DAG](https://en.wikipedia.org/wiki/Directed_acyclic_graph)
as a proof-of-work cryptocurrency with instant confirmations and
sub-second block times. It is based on [the PHANTOM protocol](https://eprint.iacr.org/2018/104.pdf), a
generalization of Nakamoto consensus.

## Overview

Spectre is a fork of [Kaspa](https://github.com/kaspanet/kaspad)
introducing CPU-only mining algorithm [SpectreX](https://github.com/spectre-project/go-spectrex).

SpectreX is based on [AstroBWTv3](https://github.com/deroproject/derohe/tree/main/astrobwt/astrobwtv3)
and proof-of-work calculation is done in the following steps:

* Step 1: SHA-3
* Step 2: AstroBWTv3
* Step 3: HeavyHash

Spectre will add full non-disclosable privacy and anonymous
transactions in future implemented with the GhostFACE protocol
build by a team of anonymous crypto algorithm researchers and
engineers. Simple and plain goal:

* PHANTOM Protocol + GhostDAG + GhostFACE = Spectre

Spectre will become a ghostchain; nothing more, nothing less. Design
decisions have been made already and more details about the GhostFACE
protocol will be released at a later stage. Sneak peak: It will use
[Pedersen Commitments](https://github.com/threehook/go-pedersen-commitment)
as it allows perfect integration with the Spectre UTXO model and
allows perfect hiding. ElGamal will be used for TX signature signing
as it has a superior TPS (transactions per second) performance. Any PRs
are welcome and can be made with anonymous accounts. No pre-mine, no
shit, pure privacy is a hit!

## Comparison

Why another fork? Kaspa is great but we love privacy, Monero and DERO
are great but we love speed! So lets join the cool things from both.
We decided to take Kaspa as codebase, quick comparison:

| Feature                      | Spectre  | Kaspa      | Monero  | DERO       |
| ---------------------------- | -------- | ---------- | ------- | ---------- |
| PoW Algorithm                | SpectreX | kHeavyHash | RandomX | AstroBWTv3 |
| Balance Encryption           | Future   | No         | Yes     | Yes        |
| Transaction Encryption       | Future   | No         | Yes     | Yes        |
| Message Encyrption           | Future   | No         | No      | Yes        |
| Untraceable Transactions     | Future   | No         | Yes     | Yes        |
| Untraceable Mining           | Yes      | No         | No      | Yes        |
| Built-in multicore CPU-miner | Yes      | No         | Yes     | Yes        |
| High BPS                     | Yes      | Yes        | No      | No         |
| High TPS                     | Yes      | Yes        | No      | No         |

Untraceable Mining is already achieved with AstroBWTv3 and a multicore
miner is already being shipped with Spectre, working on ARM/x86. There
is already a proof-of-concept Rust [AstroBWT](https://github.com/Slixe/astrobwt)
implementation currently under review and investigation to merge it
into Spectre Rust codebase. We leave it up to the community to build
an highly optimized CPU-miner.

## Mathematics

We love numbers, you will find a lot of mathematical constants in the
source code, in the genesis hash, genesis payload, genesis merkle hash
and more. Mathematical constants like [Pi](https://en.wikipedia.org/wiki/Pi),
[E](<https://en.wikipedia.org/wiki/E_(mathematical_constant)>) and
several prime numbers used as starting values for nonce or difficulty.
The first released version is `0.3.14`, the famous Pi divided by 10.

## Installation

### Install from Binaries

Pre-compiled binaries for Linux `x86_64`, Windows `x64` and macOS `x64`
as universal binary can be downloaded at: [https://github.com/spectre-project/spectred/releases](https://github.com/spectre-project/spectred/releases)

### Build from Source

Go 1.23 or later is required. Install Go according to the installation
instructions at [http://golang.org/doc/install](http://golang.org/doc/install).
Ensure Go was installed properly and is a supported version:

```bash
go version
```

Run the following commands to obtain and install spectred including
all dependencies:

```bash
git clone https://github.com/spectre-project/spectred
cd spectred
go install . ./cmd/...
```

Spectred (and utilities) should now be installed in
`$(go env GOPATH)/bin`. If you did not already add the `bin` directory
to your system path during Go installation, you are encouraged to do
so now.

### Getting Started

Spectred has several configuration options available to tweak how it
runs, but all of the basic operations work with zero configuration.

```bash
spectred
```

## Mining

The built-in `spectreminer` is a very simple unoptimized multi-core
miner. Without specifying amount of cores it spawns one mining worker
as the original Kaspa builtin miner. You can specify amount of workers
based on number of cores with the command line option `--workers`.

## Discord

Join our [Discord](https://discord.spectre-network.org/) server and
discuss with us. Don't forget: We love privacy!

## Issue Tracker

The [integrated github issue tracker](https://github.com/spectre-project/spectred/issues)
is used for this project.

## License

Spectred is licensed under the copyfree [ISC License](https://choosealicense.com/licenses/isc/).
