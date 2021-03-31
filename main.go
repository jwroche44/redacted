package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

const REPLACEMENT_STR string = "XXXX"

func main() {
	// Get the phrases to be redacted
	var phrases []string

	// 0th arg = executable file name
	// 1->(n-1) are the phrases
	// last argument should be the file name to be processed
	phrases = append(phrases, os.Args[1:len(os.Args)-1]...)

	// Sanitize the phrases (remove trailing commas, check for invalid phrases)
	for i, phrase := range phrases {
		// Check for trailing comma
		if phrase[len(phrase)-1] == ',' {
			phrase = phrase[:len(phrase)-1] // Strip the last character
		}

		// Make sure a phrase of all X's didn't get passed in
		isValidPhrase := false
		for i := range phrase {
			if phrase[i] != REPLACEMENT_STR[0] {
				isValidPhrase = true
				break
			}
		}

		if !isValidPhrase {
			log.Fatalf("Invalid replacement requested: \"%s\"", phrase)
		}

		phrases[i] = phrase
	}

	// Get the file contents that should be passed in as an argument
	file_contents := os.Args[len(os.Args)-1]

	fmt.Println(SanitizeFile(phrases, file_contents))
}

// SanitizeFile will take keys and text as input and redact any keys from the text. The resulting data will be returned
func SanitizeFile(keys []string, input string) string {
	output := input

	// Call Sanitize for every key and save the result
	for _, key := range keys {
		output = Sanitize(strings.ToUpper(key), strings.ToLower(key), output)
	}

	// output should now be sanitized
	return output
}

// Sanitize will remove any instances of key (case agnostic) from input
func Sanitize(keyU string, keyL string, input string) string {
	var sb strings.Builder

	// Loop over the input string
	for i := 0; i <= len(input)-len(keyL); i++ {
		if keyL[0] == input[i] || keyU[0] == input[i] {
			keystart := i
			keyend := i

			// Loop over the input to validate if we found the key
			for j := 1; j < len(keyL); j++ {
				if keyL[j] == input[i+j] || keyU[j] == input[i+j] {
					keyend = i + j
				} else {
					break
				}
			}

			// Check if we found the key
			if keyend-keystart == len(keyL)-1 {

				// Key found, build the output string
				sb.WriteString(input[0:keystart])
				sb.WriteString(REPLACEMENT_STR)

				// Call sanitize on the remaining substring
				sb.WriteString(Sanitize(keyU, keyL, input[keyend+1:]))

				return sb.String()
			}
		}
	}

	// if we get here, the key did not exist in the string so just return the original string
	return input
}
