package anagrams

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"sort"
	"strings"
	"time"
	"unicode/utf8"
)

//RuneSlice type attaches functions to allow sorting of runes
type RuneSlice []rune

func (rs RuneSlice) Len() int {
	return utf8.RuneCountInString(string(rs))
}

func (rs RuneSlice) Less(i, j int) bool {
	return rs[i] < rs[j]
}

func (rs RuneSlice) Swap(i, j int) {
	rs[i], rs[j] = rs[j], rs[i]
}

//LenCheck returns true if strings are of equal length, else false
func LenCheck(x string, y string) bool {
	if len(x) == len(y) {
		return true
	}
	return false
}

//SortWord takes a string and returns its characters sorted as a string
func SortWord(word string) string {
	w := []rune(word)
	rs := RuneSlice(w)
	sort.Sort(rs)
	return string(rs)
}

//Anagram1 sorts the characters of two given strings and returns true if they match
func Anagram1(x string, y string) bool {
	if LenCheck(x, y) {
		//Continue anagram check
		//TODO: map characters
		amapX := SortWord(x)
		amapY := SortWord(y)
		if amapX == amapY {
			return true
		}
	}
	return false
}

//ReadSystemWords parses the system word list(/usr/share/dict/words) and returns
//an array of strings.
func ReadSystemWords() ([]string, error) {
	path := "/usr/share/dict/words"
	contents, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
		return nil, err
	}
	return strings.Split(string(contents), "\n"), nil
}

//AnagramMap holds construct of sorted characters of anagrams to all the
//possible words in the 'Mapping' field.  Mapping field holds a 'set' map to
//allow fast lookup of unique words possible.
type AnagramMap struct {
	Mapping map[string]map[string]bool
}

//AnagramOfWord takes a word and returns a separate anagram of it
func (a *AnagramMap) AnagramOfWord(word string) string {
	word = strings.ToLower(word)
	wordKey := SortWord(word)
	wordMap := a.Mapping[wordKey]
	if len(wordMap) <= 1 {
		return word
	} else {
		uniqueWords := make([]string, 0)
		for k, _ := range wordMap {
			if k != word {
				uniqueWords = append(uniqueWords, k)
			}
		}

		t := time.Now()
		r := rand.New(rand.NewSource(t.UnixNano()))
		w := uniqueWords[r.Intn(len(uniqueWords))]
		return w
	}
}

//AnagramSentence takes a list of words and replaces each word with
//anagram of the word if possible.
func (a *AnagramMap) AnagramSentence(sent []string) []string {
	var ret []string
	fmt.Println(sent)
	for _, s := range sent {
		fmt.Println(s)
		ret = append(ret, a.AnagramOfWord(s))
	}
	return ret
}

//AnagramList is an ill named function which creates the mapping of character
//posibilities to their possible anagrams.
func AnagramList(words []string) map[string]map[string]bool {
	anagrams := make(map[string]map[string]bool)
	for _, w := range words {
		w = strings.ToLower(w)
		wordKey := SortWord(w)
		if anagrams[wordKey] != nil {
			anagrams[wordKey][w] = true
		} else {
			anagrams[wordKey] = make(map[string]bool)
			anagrams[wordKey][w] = true
		}
	}
	return anagrams
}
