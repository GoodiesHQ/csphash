package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"hash"
	"io"
	"os"
)

type algo func() hash.Hash

var hashers = map[uint]algo{
	224: sha256.New224,
	256: sha256.New,
	384: sha512.New384,
	512: sha512.New,
}

func hashFile(filename string, bits uint) (string, error) {
	hasher, found := hashers[bits]
	if !found {
		return "", fmt.Errorf("invalid bit size '%d'", bits)
	}

	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}

	h := hasher()
	if _, err := io.Copy(h, file); err != nil {
		return "", err
	}

	data := h.Sum(nil)
	return fmt.Sprintf("sha%d-%s", bits, base64.StdEncoding.EncodeToString(data)), nil
}
