package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"log"
)

func encrypt(p []byte, k []byte) []byte {
	block, err := aes.NewCipher(k)
	if err != nil {
		log.Fatal(err)
	}

	ciphertext := make([]byte, aes.BlockSize+len(p))
	iv := ciphertext[:aes.BlockSize]

	if _, err = io.ReadFull(rand.Reader, iv); err != nil {
		log.Fatal(err)
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[aes.BlockSize:], p)

	return ciphertext
}

func decrypt(c []byte, k []byte) []byte {
	block, err := aes.NewCipher(k)
	if err != nil {
		log.Fatal(err)
	}

	if len(c) < aes.BlockSize {
		log.Fatal("Ciphertext is too short.")
	}

	iv := c[:aes.BlockSize]
	c = c[aes.BlockSize:]
	plaintext := make([]byte, len(c))

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(plaintext, c)

	return plaintext
}
