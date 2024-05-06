package services

import (
	"io"
	"log"
	"net/http"
	"regexp"
	"strings"
)

func getString(url string) string {
	resp, err := http.Get(url)
	notErr(err)
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	notErr(err)
	return string(body)
}
func getWordsSlice(str string) []string {
	cleanStr := regexp.MustCompile(`[,.\n]`).ReplaceAllString(str, "")
	words := strings.Fields(cleanStr)
	return words
}

func countWords(words []string) map[string]int {
	wordFreq := make(map[string]int)
	for _, word := range words {
		word = strings.ToLower(word)
		wordFreq[word]++
	}
	return wordFreq
}
func notErr(err error) {
	if err != nil {
		log.Fatal(err)
		return
	}
}
