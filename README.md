[![Go Report Card](https://goreportcard.com/badge/github.com/schrius/password-validator)](https://goreportcard.com/report/github.com/schrius/password-validator)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
![Go](https://github.com/schrius/password-validator/workflows/Go/badge.svg)
# Validate Password

Password Validator validate the following rules:
1. Have an 8 character minimum
2. AT LEAST 64 character maximum
3. Allow all ASCII characters and spaces
4. Not be a common password

Validator accept password from STDIN in newline delimited format and print invalid password to the command line.

## Prerequisites
go > 1.12 is required

### Install
Run go command to get all dependencies
```bash
go get -u ./...
```

### Build
To build an excutable run:
```bash
go build -o ./password_validator ./password_validator.go
```

## Validate Example
```bash
cat input_passwords.txt | ./password_validator weak_password_list.txt
mom -> Error: Too Short
password1 -> Error: Too Common
*** -> Error: Invalid Charaters
```

## Quick Start
Makefile contain command to get package, test, and build the program in an easy way
 ```bash
make
```