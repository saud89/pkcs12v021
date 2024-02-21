package main

import (
	"fmt"
	"log"
	"os"

	"software.sslmate.com/src/go-pkcs12"
)

func main() {
	argsToMain := os.Args[1:]
	if len(argsToMain) == 0 {
		log.Fatal("please provide the path of the file to read")
	}

	fileContent, err := os.ReadFile(argsToMain[0])
	if err != nil {
		log.Fatal("error while reading the file, err: ", err)

	}
	password := argsToMain[1]

	privateKey, cert, cacerts, err := pkcs12.DecodeChain(fileContent, password)
	if err != nil {
		log.Fatal("error while decoding the certificates, err: ", err)
	}
	fmt.Printf("private key is : %+v\n", privateKey)
	fmt.Printf("cert is : %+v\n", cert)
	fmt.Printf("cacerts is : %+v\n", cacerts)
}
