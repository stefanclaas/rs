package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
	"strconv"
)

func generateRandomString(length int) string {
	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
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
	if len(os.Args) != 3 {
		fmt.Println("Usage: random [Number of characters] [Number of strings]")
		return
	}

	length, err1 := strconv.Atoi(os.Args[1])
	count, err2 := strconv.Atoi(os.Args[2])

	if err1 != nil || err2 != nil || length <= 0 || count <= 0 {
		fmt.Println("Invalid Input. Please enter valid integers for the number of characters and strings.")
		return
	}

	for i := 0; i < count; i++ {
		randomString := generateRandomString(length)
		fmt.Println(randomString)
	}
}

