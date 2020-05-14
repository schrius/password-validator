# Validate Password

Password Validator validate the following rules:
1. Have an 8 character minimum
2. AT LEAST 64 character maximum
3. Allow all ASCII characters and spaces
4. Not be a common password

Validator accept password from STDIN in newline delimited format and print invalid password to the command line.

### Install
Run go command to get all dependencies
```bash
go get -u ./...
```

### Build
To build an excutable run:
```bash
go build -o ./validator ./password_validator.go
```

## Validate Example
```bash
cat input_passwords.txt | ./password_validator weak_password_list.txt
mom -> Error: Too Short
password1 -> Error: Too Common
*** -> Error: Invalid Charaters
```

## MakeFile
Makefile contain command to get package, test, and build the program in an easy way
 ```bash
make
```