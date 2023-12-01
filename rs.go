package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"flag"
)

var (
	lowercase = flag.Bool("l", false, "generate lowercase strings")
	numeric = flag.Bool("n", false, "generate numeric strings")
	alpha = flag.Bool("a", false, "generate alphabetic strings")
	mixed = flag.Bool("m", false, "generate mixed strings (letters/numbers)")
	blank = flag.Bool("b", false, "add a blank line between strings")
)

func generateRandomString(length int) string {
	charset := "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	if *lowercase {
		charset = "abcdefghijklmnopqrstuvwxyz"
	}
	if *numeric {
		charset = "0123456789"
	}
	if *alpha {
		if *lowercase {
			charset = "abcdefghijklmnopqrstuvwxyz"
		} else {
			charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		}
	}
	if *mixed {
		if *lowercase {
			charset = "abcdefghijklmnopqrstuvwxyz0123456789"
		} else {
			charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
		}
	}

	result := make([]byte, length)
	max := big.NewInt(int64(len(charset)))

	for i := 0; i < length; i++ {
		n, err := rand.Int(rand.Reader, max)
		if err != nil {
			panic(err)
		}
		result[i] = charset[n.Int64()]
	}

	return string(result)
}

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options] <Number of characters> <Number of strings>]\n", os.Args[0])
		fmt.Fprintln(os.Stderr, "Options:")
		flag.PrintDefaults()
	}

	flag.Parse()

	if len(flag.Args()) != 2 {
		flag.Usage()
		return
	}

	length, err1 := strconv.Atoi(flag.Args()[0])
	count, err2 := strconv.Atoi(flag.Args()[1])

	if err1 != nil || err2 != nil || length <= 0 || count <= 0 {
		fmt.Println("Invalid Input. Please enter valid integers for the number of characters and strings.")
		return
	}

	for i := 0; i < count; i++ {
		randomString := generateRandomString(length)
		fmt.Println(randomString)
		if *blank && i != count-1 {
			fmt.Println()
		}
	}
}

