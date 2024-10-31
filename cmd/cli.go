package main

import (
	"flag"
	"fmt"
	"os"
)

func cli() ([]uint, []string) {
	sha224 := flag.Bool("sha224", false, "Enable SHA-224 Hash Output")
	sha256 := flag.Bool("sha256", false, "Enable SHA-256 Hash Output (Default if None Specified)")
	sha384 := flag.Bool("sha384", false, "Enable SHA-384 Hash Output")
	sha512 := flag.Bool("sha512", false, "Enable SHA-512 Hash Output")
	all := flag.Bool("all", false, "Enable All SHA Algorithm Hash Outputs")

	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage: %s [options] <file1> <file2> ...\n", os.Args[0])
		fmt.Println("\nOptions:")
		flag.PrintDefaults()
		fmt.Println("\nPositional Arguments:")
		fmt.Println("  <file1> <file2> ...   List of filenames to hash")
		fmt.Println()
	}

	flag.Parse()

	var bits []uint

	if !*sha224 && !*sha256 && !*sha384 && !*sha512 {
		*sha256 = true
	}

	if *all {
		*sha224 = true
		*sha256 = true
		*sha384 = true
		*sha512 = true
	}

	if *sha224 {
		bits = append(bits, 224)
	}
	if *sha256 {
		bits = append(bits, 256)
	}
	if *sha384 {
		bits = append(bits, 384)
	}
	if *sha512 {
		bits = append(bits, 512)
	}

	filenames := flag.Args()
	if len(filenames) == 0 {
		flag.Usage()
		os.Exit(0)
	}

	return bits, flag.Args()
}
