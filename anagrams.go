package anagrams

import (
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

func ReadSystemWords() ([]string, error) {
	path := "/usr/share/dict/words"
	contents, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
		return nil, err
	}
	return strings.Split(string(contents), "\n"), nil
}

//AnagramMap holds construct of sorted characters of anagrams to all the possible words in the 'mapping' field
type AnagramMap struct {
	mapping map[string][]string
}

//AnagramOfWord takes a word and returns a separate anagram of it
func (a *AnagramMap) AnagramOfWord(word string) string {
	wordKey := SortWord(word)
	list := a.mapping[wordKey]
	if len(list) <= 1 {
		return word
	} else {
		t := time.Now()
		r := rand.New(rand.NewSource(t.UnixNano()))
		for {
			w := list[r.Intn(len(list))]
			if w != word {
				return w
			}
		}
	}
}

func (a *AnagramMap) AnagramSentence(sent []string) []string {
	var ret []string
	for _, s := range sent {
		ret = append(ret, a.AnagramOfWord(s))
	}
	return ret
}

func AnagramList(words []string) map[string][]string {
	anagrams := make(map[string][]string)
	for _, w := range words {
		wordKey := SortWord(w)
		anagrams[wordKey] = append(anagrams[wordKey], w)
	}
	return anagrams
}
