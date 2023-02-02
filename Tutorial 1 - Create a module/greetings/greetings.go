package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// In Go, a function whose name starts with a capital letter can be called by a function not in the same package. This is known in Go as an exported name.
func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func Hellos(names []string) (map[string]string, error) {
	// A map to associate names with messages. In Go, you initialize a map with the following syntax: make(map[key-type]value-type)
	messages := make(map[string]string)
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}
	return messages, nil
}

// Go executes init functions automatically at program startup, after global variables have been initialized.
// Init sets initial values for variables used in the function.
func init() {
	rand.Seed(time.Now().UnixNano())
}

// starts with a lowercase letter, making it accessible only to code in its own package (in other words, it's not exported).
func randomFormat() string {
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	return formats[rand.Intn(len(formats))]
}
