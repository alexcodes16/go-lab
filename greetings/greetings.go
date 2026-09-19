package greetings

import (
    "errors"
    "fmt"
    "math/rand"
)

// Hello returns a random greeting for the named person.
func Hello(name string) (string, error) {
    if name == "" {
        return "", errors.New("empty name")
    }

    formats := []string{
        "Hi, %v. Welcome!",
        "Great to see you, %v!",
        "Hail, %v! Well met!",
    }

    message := fmt.Sprintf(formats[rand.Intn(len(formats))], name)
    return message, nil
}

// Hellos returns a map that associates each of the named people
// with a greeting message.
func Hellos(names []string) (map[string]string, error) {
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
