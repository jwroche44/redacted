package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/jwroche44/redacted/sanitizer"
)

func main() {
	// Parse the inputs
	keyf, inf, outf, err := processInputs()
	if err != nil {
		log.Fatal(err)
	}

	// Make a sanitizer that we can use for this
	var s *sanitizer.Sanitizer

	// Set the keywords
	if err = s.LoadKeywordsFromFile(keyf); err != nil {
		log.Fatal(err)
	}

	// Run the Sanitizer on the input file
	if err = s.SanitizeFile(inf, outf); err != nil {
		log.Fatal(err)
	}

	log.Printf("Sanitization completed successfuly.")
}

// processInputs will take the command line arguments as input, validate and return their values
func processInputs() (keyFile string, inFile string, outFile string, err error) {
	keyFile = *flag.String("k", "", "The path to the file containing keywords for redacting")
	inFile = *flag.String("i", "", "The path to the input file to be sanitized")
	outFile = *flag.String("o", "./output.txt", "The path where the sanitized file should be placed")

	flag.Parse()

	// Err should be nil unless something is wrong
	err = nil

	if keyFile == "" || inFile == "" {
		err = fmt.Errorf("The key and input file parameters must be specified")
		return
	}

	// Validate the required files exist
	if !validateFile(keyFile) {
		err = fmt.Errorf("Unknown key file specified: %s", keyFile)
	} else if !validateFile(inFile) {
		err = fmt.Errorf("Unknown input file specified: %s", inFile)
	}

	return
}

// validateFile will check if a file at the specified path exists
func validateFile(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}

	return true
}
