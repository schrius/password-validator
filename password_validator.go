package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/schrius/password-validator/validate"
)

// Password validator
// validate the following rules
// 1. Have an 8 character minimum
// 2. AT LEAST 64 character maximum
// 3. Allow all ASCII characters and spaces (unicode optional)
// 3. Not be a common password
// expmale intput and output
// cat input_passwords.txt | ./password_validator weak_password_list.txt
// mom -> Error: Too Short
// password1 -> Error: Too Common
// *** -> Error: Invalid Charaters

// validate password
// push result to channel
func validatePassword(password string, weakList map[string]bool, wg *sync.WaitGroup, ch chan error) {
	defer wg.Done()
	ch <- validate.Validate(password, weakList)
}

// Each vailidattion run in separated goroutines concurrently
func main() {
	wg := new(sync.WaitGroup)
	passwordSize := 0
	scanner := bufio.NewScanner(os.Stdin)
	weakList := make(map[string]bool)
	ch := make(chan error)

	// load file if weak password list is provided
	if len(os.Args) > 1 {
		weakList = validate.LoadWeakPasswordList(os.Args[1])
	}

	// check each string if it meet the requirement
	for scanner.Scan() {
		password := scanner.Text()
		wg.Add(1)
		passwordSize++
		go validatePassword(password, weakList, wg, ch)
	}

	// Retrieve Error from channel and display the message
	for i := 0; i < passwordSize; i++ {
		err := <-ch
		if err != nil {
			fmt.Println(err)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}

	wg.Wait()
}
