# Unix Ti ISO

Tool to parse a Unix timestamp and print it as an ISO date.

## Installation

### Unix users (Linux and MacOSX)

Unix users may download and install latest *unix2iso* release with command:

```bash
sh -c "$(curl https://sweetohm.net/dist/unix2iso/install)"
```

If *curl* is not installed on you system, you might run:

```bash
sh -c "$(wget -O - https://sweetohm.net/dist/unix2iso/install)"
```

**Note:** Some directories are protected, even as *root*, on **MacOSX** (since *El Capitan* release), thus you can't install *unix2iso* in */usr/bin* for instance.

### Binary package

Otherwise, you can download latest binary archive at <https://github.com/c4s4/unix2iso/releases>. Unzip the archive, put the binary of your platform somewhere in your *PATH* and rename it *unix2iso*.

## Usage

To convert Unix timestamp *1609459200* to ISO date, run:

```bash
$ unix2iso 1609459200
2021-01-01T00:00:00Z
```

To print a more human-readable format, use the `-r` flag:

```bash
$ unix2iso -r 1609459200
2021-01-01 00:00:00 UTC
```

You can get help with `-h` flag:

```bash
$ unix2iso -h
unix2iso [-h] [-v] [-r] unix
Convert UNIX timestamps to ISO 8601 format:
-h     To print this help
-v     To print version
-r     To print human readable format
unix   UNIX timestamp to convert
```

And version with `-v` flag:

```bash
$ unix2iso -v
1.0.0
```

*Enjoy!*
