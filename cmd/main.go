package main

import (
	"fmt"
)

func main() {
	bitsizes, filenames := cli()

	for _, filename := range filenames {
		fmt.Println("File:", filename)
		for _, bitsize := range bitsizes {
			hash, err := hashFile(filename, bitsize)
			if err != nil {
				fmt.Println("Failed:", err)
				fmt.Println()
				continue
			}
			fmt.Println(hash)
		}
		fmt.Println()
	}
}
