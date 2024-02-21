# Description
This program verifies decoding the certificate key pair with go-pkcs12 library version 0.2.1

# Pre-requisites
This is a golang program, and will need golang runtime to be installed in the system.
Follow the instructions mentioned in this document for go installation: https://go.dev/doc/install

# Executing the code
To run the program, 2 arguments are needed, absolute path to the certificate and password to decode the certificate.

Following is the command:
```
go run main.go "<absolute path to certificate file>" "<password>"
```

Example:
```
go run main.go "/Users/ABC/certFolder/certificate.p12" "passwordToFile"
```
