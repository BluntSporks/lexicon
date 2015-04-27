package lexicon

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
)

// DefaultDataPath gets the lexicon word list file location from the LEXICON_DATA environment variable.
func DefaultDataPath() string {
	dir := os.Getenv("LEXICON_DATA")
	if dir == "" {
		log.Fatal("Set LEXICON_DATA variable to directory of lexicon data files")
	}
	return dir
}

// CountChars counts all the characters in a language lexicon.
func CountChars(words map[string]bool) map[rune]int {
	chars := make(map[rune]int)
	for word := range words {
		for _, ch := range word {
			chars[ch]++
		}
	}
	return chars
}

// LoadAllLanguages loads all the language files.
func LoadAllLanguages(langDir string) map[string]map[string]bool {
	// Load the languages.
	langFiles, err := ioutil.ReadDir(langDir)
	if err != nil {
		log.Fatal(err)
	}
	langWords := make(map[string]map[string]bool)
	for _, langFile := range langFiles {
		name := langFile.Name()
		path := path.Join(langDir, name)
		langWords[name] = LoadLanguage(path)
	}
	return langWords
}

// LoadLanguage loads a language file.
func LoadLanguage(langFile string) map[string]bool {
	// Open file.
	handle, err := os.Open(langFile)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()

	// Scan file line by line.
	words := make(map[string]bool)
	scanner := bufio.NewScanner(handle)
	for scanner.Scan() {
		line := scanner.Text()
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
		word := strings.TrimSpace(strings.ToLower(line))
		words[word] = true
	}
	return words
}
