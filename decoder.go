package main

import (
	"bufio"
	"bytes"
	"crypto/sha256"
	"encoding/base64"
	"io"
	"log"
)

type encryptedDecoder struct {
	buf *bufio.Reader
	key [32]byte
}

func NewEncryptedDecoder(p string, r io.Reader) encryptedDecoder {
	ed := encryptedDecoder{
		buf: bufio.NewReader(r),
		key: sha256.Sum256([]byte(p)),
	}

	// Test for magic bytes
	header, err := ed.buf.ReadBytes([]byte(":")[0])
	if err != nil {
		log.Fatal(err)
	}

	if !bytes.Equal(header, []byte("b64:")) {
		log.Fatal("Invalid input file, magic bytes not found.")
	}

	return ed
}

func (ed encryptedDecoder) Read(p []byte) (int, error) {
	// Read base64 chunk (delimited by ;)
	s, err := ed.buf.ReadString([]byte(";")[0])
	if err != nil {
		if err == io.EOF {
			return 0, err
		}
		log.Fatal(err)
	}

	// Decode base64 chunk
	ciphertext, err := base64.StdEncoding.DecodeString(s[:len(s)-1])
	if err != nil {
		log.Fatal(err)
	}

	// Decrypt ciphertext
	plaintext := decrypt(ciphertext, ed.key[:])

	// Copy plaintext to read buffer
	n := copy(p, plaintext)

	return n, nil
}
