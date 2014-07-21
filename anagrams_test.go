package anagrams

import (
	"fmt"
	"testing"
)

func TestLengthChk(t *testing.T) {
	x := "hihi"
	y := "hihii"
	z := "ihih"

	result := LenCheck(x, y)
	if result == true {
		t.Errorf("Strings of different length but true returned")
	}

	result = LenCheck(x, z)
	if result == false {
		t.Errorf("Strings of same length but false returned..")
	}
}

func TestRuneSliceSort(t *testing.T) {
	x := "hihi"
	y := "hhii"

	result := SortWord(x)
	if result != y {
		t.Errorf("Word wasn't sorted properly: %s", result)
	}
}

func TestRuneSliceSort2(t *testing.T) {
	x := "watcatw"
	y := "aacttww"
	result := SortWord(x)

	if result != y {
		t.Errorf("Word wasn't sorted properly: %s", result)
	}
}

func TestAnagram1(t *testing.T) {
	x := "heater"
	y := "reheat"

	result := Anagram1(x, y)
	if result == false {
		t.Error("Anagram not detected correctly")
	}
}

func TestReadWords(t *testing.T) {
	words, err := ReadSystemWords()
	if err != nil {
		t.Log("No error reading word list")
	}
	t.Log(words[:5])
	if len(words) < 90000 {
		t.Errorf("Words not read from dictionary successfully")
	}
}

func TestAnagramList(t *testing.T) {
	words, err := ReadSystemWords()
	if err != nil {
		t.Log("No error reading word list")
	}
	anagrams := AnagramList(words)
	if len(anagrams) < 5000 {
		t.Log(anagrams["acr"])
		t.Errorf("Number of anagram combinations dubiously low for number of words..")
	}
	/*
		for k, v := range anagrams {
			if len(v) > 2 {
				fmt.Println(k, v)
			}
		}
		fmt.Println(anagrams["aflt"])
	*/
}

func TestAnagramMap(t *testing.T) {
	words, err := ReadSystemWords()
	if err != nil {
		t.Log("No error reading word list")
	}
	anagrams := AnagramList(words)

	AM := &AnagramMap{mapping: anagrams}
	word := "ropes"
	ana := AM.AnagramOfWord(word)
	if ana == word {
		t.Log(ana, word)
		t.Errorf("Words should not be equivalent!")
	}
	//fmt.Println(ana)

	word = "jjjjjjjj"
	ana = AM.AnagramOfWord(word)
	if ana != word {
		fmt.Printf("returned word: %s\n", ana)
		t.Errorf("There is no word that is all 'j's...")
	}

	sentence := [...]string{"The", "quick", "brown", "fox", "jumps", "over", "the", "lazy", "dog"}
	s := sentence[:]
	anasent := AM.AnagramSentence(s)
	if anasent[len(anasent)-1] != "god" {
		fmt.Printf("returned sentence: %#v\n", anasent)
		t.Errorf("'god' should be replaced with a 'dog'")
	}
}
