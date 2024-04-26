# spectreminer

`spectreminer` is a CPU-based miner for `spectred`.

## Requirements

Go 1.19 or later.

## Build from Source

* Install Go according to the installation instructions here:
  http://golang.org/doc/install

* Ensure Go was installed properly and is a supported version:

```bash
go version
```

* Run the following commands to obtain and install `spectred`
  including all dependencies:

```bash
git clone https://github.com/spectre-project/spectred
cd spectred/cmd/spectreminer
go install .
```

* `spectreminer` should now be installed in `$(go env GOPATH)/bin`.
  If you did not already add the bin directory to your system path
  during Go installation, you are encouraged to do so now.
  
## Usage

The full `spectreminer` configuration options can be seen with:

```bash
spectreminer --help
```

But the minimum configuration needed to run it is:

```bash
spectreminer --miningaddr=<YOUR_MINING_ADDRESS>
```
