package main

import (
	"flag"
	"io"
	"log"
	"os"
)

func encode(inputFilename string, outputFilename string, password string) {
	inputFile, err := os.OpenFile(inputFilename, os.O_RDONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer inputFile.Close()

	outputFile, err := os.OpenFile(outputFilename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer outputFile.Close()

	encoder := NewEncryptedEncoder(password, outputFile)

	if _, err := io.Copy(encoder, inputFile); err != nil {
		log.Fatal(err)
	}
}

func decode(inputFilename string, outputFilename string, password string) {
	inputFile, err := os.OpenFile(inputFilename, os.O_RDONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer inputFile.Close()

	outputFile, err := os.OpenFile(outputFilename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer outputFile.Close()

	decoder := NewEncryptedDecoder(password, inputFile)

	if _, err := io.Copy(outputFile, decoder); err != nil {
		log.Fatal(err)
	}
}

func main() {
	var mode string
	var password string
	var inputFilename string
	var outputFilename string

	log.SetFlags(0)

	flag.StringVar(&mode, "mode", "", "[encode, decode]")
	flag.StringVar(&inputFilename, "input", "", "input filename")
	flag.StringVar(&outputFilename, "output", "", "output filename")
	flag.StringVar(&password, "password", "", "encryption password")
	flag.Parse()

	if flag.NFlag() != 4 {
		flag.Usage()
		os.Exit(1)
	}

	switch mode {
	case "encode":
		encode(inputFilename, outputFilename, password)
	case "decode":
		decode(inputFilename, outputFilename, password)
	default:
		log.Fatal("Invalid mode provided")
	}
}
