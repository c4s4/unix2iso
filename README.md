# Unix Ti ISO

Tool to parse a Unix timestamp and print it as an ISO date.

## Installation

Download the binary from the releases page and place it in your PATH.

## Usage

To convert Unix timestamp *1609459200* to ISO date, run:

```bash
$ unix2iso 1609459200
2021-01-01T00:00:00Z
```

To print a more human-readable format, use the `-h` flag:

```bash
$ unix2iso -h 1609459200
2021-01-01 00:00:00 UTC
```

*Enjoy!*

