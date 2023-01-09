package words

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/rs/zerolog/log"
)

var dictionary map[string][]string

func alphabetize(word string) string {
	s := strings.Split(word, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func buildDictionary() {
	dictionary := make(map[string][]string)
	txtwords := readWords("words/words.txt")
	log.Info().Int("txtwords_count", len(txtwords)).Msg("txtwords slice size")
	for _, word := range txtwords {
		alphabetized := alphabetize(word)
		var lst []string
		if len(dictionary) > 0 && len(dictionary[alphabetized]) > 0 {
			lst = dictionary[alphabetized]
		} else {
			lst = []string{}
		}
		lst = append(lst, word)
		dictionary[alphabetized] = lst
	}
}

func readWords(fName string) []string {
	txtwords := make([]string, 0)

	file, err := os.Open(fName)
	if err != nil {
		log.Info().Err(err).Msgf("failed to open a file: %s", fName)
		return txtwords
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		txtwords = append(txtwords, scanner.Text())
	}

	return txtwords
}

func Permutations(word string) {
	buildDictionary()
	aw := alphabetize(word)
	fmt.Printf("Permutation group of %s is %s", word, dictionary[aw])
}
