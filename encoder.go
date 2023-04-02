package main

import (
	"bufio"
	"crypto/sha256"
	"encoding/base64"
	"io"
	"log"
)

type encryptedEncoder struct {
	buf *bufio.Writer
	key [32]byte
}

func NewEncryptedEncoder(p string, w io.WriteCloser) encryptedEncoder {
	ee := encryptedEncoder{
		buf: bufio.NewWriter(w),
		key: sha256.Sum256([]byte(p)),
	}

	// Initialize output file with magic bytes
	ee.buf.Write([]byte("b64:"))

	return ee
}

func (ee encryptedEncoder) Write(p []byte) (int, error) {
	// Encrypt data
	ciphertext := encrypt(p, ee.key[:])

	// Base64 encode ciphertext
	s := base64.StdEncoding.EncodeToString(ciphertext)

	// Write base64 chunk with delimiter
	if _, err := ee.buf.WriteString(s + ";"); err != nil {
		log.Fatal()
	}

	return len(p), nil
}
