package validate

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"unicode"
)

// Error message for invalid password
type Error struct {
	Password string
	Messsage string
}

func (e *Error) Error() string {
	return fmt.Sprintf("%s -> Error: %s",
		e.Password, e.Messsage)
}

// ValidLength
// validate password length
// minimum 8 and maximum is 64
func ValidLength(password string) bool {
	return len(password) >= 8 && len(password) <= 64
}

// ValidLetter
// validate password if characters are letters, marks, numbers, punctuation, symbols, and the ASCII space character
func ValidLetter(password string) bool {
	for _, char := range password {
		if !unicode.IsPrint(char) {
			return false
		}
	}
	return true
}

// IsWeakPassword
// check password if it is weak by searching the weak password fine
// this should use for single password validation
func IsWeakPassword(password, path string) bool {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() == password {
			return true
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return false
}

// LoadWeakPasswordList load strings from file and store as a map collection
func LoadWeakPasswordList(path string) map[string]bool {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	weakPasswordList := make(map[string]bool)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		weakPasswordList[scanner.Text()] = true
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return weakPasswordList
}

// Validate all requirment length, printable letter and weak password
func Validate(password string, weaklist map[string]bool) error {
	if !ValidLength(password) {
		return &Error{
			Password: password,
			Messsage: "Invalid length: minimum 8 to maximum 64"}
	} else if !ValidLetter(password) {
		return &Error{
			Password: password,
			Messsage: "Invalid Charaters"}
	} else if weaklist != nil {
		if _, ok := weaklist[password]; ok {
			return &Error{
				Password: password,
				Messsage: "Too Common"}
		} else {
			return nil
		}
	} else {
		return nil
	}
}
