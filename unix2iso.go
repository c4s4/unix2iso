package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// Version as printed with -version option
var Version = "UNKNOWN"

const (
	// Help as printed with -help option
	Help = `unix2iso [-h] [-v] [-r] unix
Convert UNIX timestamps to ISO 8601 format:
-h     To print this help
-v     To print version
-r     To print human readable format
unix   UNIX timestamp to convert`
)

func main() {
	help := flag.Bool("h", false, "Print help")
	version := flag.Bool("v", false, "Print version")
	human := flag.Bool("r", false, "Print human readable format")
	flag.Parse()
	if *help {
		fmt.Println(Help)
		os.Exit(0)
	}
	if *version {
		fmt.Println(Version)
		os.Exit(0)
	}
	if flag.NArg() != 1 {
		fmt.Println("ERROR missing UNIX timestamp argument")
		fmt.Println(Help)
		os.Exit(1)
	}
	unix, err := strconv.ParseInt(flag.Arg(0), 10, 64)
	if err != nil {
		fmt.Printf("ERROR invalid UNIX timestamp: %v\n", err)
		fmt.Println(Help)
		os.Exit(1)
	}
	iso := Unix2iso(unix, *human)
	fmt.Println(iso)
}

func Unix2iso(unix int64, human bool) string {
	iso := time.Unix(unix, 0).UTC().Format("2006-01-02T15:04:05Z")
	if human {
		iso = strings.Replace(iso, "T", " ", 1)
		iso = strings.Replace(iso, "Z", " UTC", 1)
	}
	return iso
}
