// pqcgate is the CLI used in CI to sign and verify image digests.
package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: pqcgate <keygen|sign|verify>")
		os.Exit(1)
	}
	switch os.Args[1] {
	case "keygen", "sign", "verify":
		fmt.Printf("%s: not implemented yet\n", os.Args[1])
	default:
		fmt.Println("unknown command:", os.Args[1])
		os.Exit(1)
	}
}
