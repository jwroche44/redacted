package sanitizer

import "fmt"

// Sanitizer is the base struct for parsing files
type Sanitizer struct {
	keywords []string
	// TODO: Add in statistics?
	// - How many times each keyword was replaced?
}

// LoadKeywordsFromFile will load any keywords specified in the file into memory
func (s *Sanitizer) LoadKeywordsFromFile(path string) error {
	return fmt.Errorf("ERROR: Not Implemented")
}

// SanitizeFile will santize the file at inputPath using the configured keywords
// and save the result to the file at outputPath
func (s Sanitizer) SanitizeFile(inputPath string, outputPath string) error {
	return fmt.Errorf("ERROR: Not Implemented")
}
