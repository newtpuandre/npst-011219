package main

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"log"
	"strconv"
	"strings"
)

var (
	hash     = "42f82ae6e57626768c5f525f03085decfdc5c6fe"
	alphabet = [26]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	numbers  = [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	symbols  = [12]string{"*", "@", "!", "#", "%", "&", "(", ")", "^", "~", "{", "}"}
)

func main() {

	var output []string
	output = make([]string, 4)

	//Symbols first
	for _, symbol := range symbols {
		output[0] = symbol

		for _, number := range numbers {
			numberString := strconv.Itoa(number)
			output[1] = numberString

			for _, upperCase := range alphabet {
				output[2] = strings.ToUpper(upperCase)

				for _, lowercase := range alphabet {
					output[3] = lowercase

					checkHash(output)
				}

			}

		}

	}

}

func checkHash(output []string) {
	actualOutput := strings.Join(output, "")

	h := sha1.New()
	h.Write([]byte(actualOutput))
	bs := h.Sum(nil)

	shasum := hex.EncodeToString(bs)
	if shasum == hash {
		fmt.Println("Found password for hash: " + hash + " | Password: " + actualOutput)
		log.Fatal("done.")
	}
}
