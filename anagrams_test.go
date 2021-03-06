package anagrams

import (
	"fmt"
	"testing"
)

func TestLengthChk(t *testing.T) {
	x := "hihi"
	y := "hihii"
	z := "ihih"

	result := lenCheck(x, y)
	if result == true {
		t.Errorf("Strings of different length but true returned")
	}

	result = lenCheck(x, z)
	if result == false {
		t.Errorf("Strings of same length but false returned..")
	}
}

func TestRuneSliceSort(t *testing.T) {
	x := "hihi"
	y := "hhii"

	result := sortWord(x)
	if result != y {
		t.Errorf("Word wasn't sorted properly: %s", result)
	}
}

func TestRuneSliceSort2(t *testing.T) {
	x := "watcatw"
	y := "aacttww"
	result := sortWord(x)

	if result != y {
		t.Errorf("Word wasn't sorted properly: %s", result)
	}
}

func TestAnagram1(t *testing.T) {
	x := "heater"
	y := "reheat"

	result := anagram1(x, y)
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

func TestNewAnagramMap(t *testing.T) {
	var testtable = []struct {
		wordpath string
		works    bool
	}{
		{
			wordpath: "",
			works:    true,
		},
		{
			wordpath: "/usr/share/dict/words",
			works:    true,
		},
		{
			wordpath: "/dev/null",
			works:    false,
		},
	}

	for _, tt := range testtable {
		am, err := NewAnagramMap(tt.wordpath)
		if err != nil && tt.works {
			t.Errorf("error[%v] was not excpected for '%s'", err, tt.wordpath)
		}
		if am == nil && !tt.works {
			t.Errorf("anagrammap should have been created for %s", tt.wordpath)
		}
	}
}

func TestAnagramList(t *testing.T) {
	words, err := ReadSystemWords()
	if err != nil {
		t.Log("No error reading word list")
	}
	anagrams := anagramList(words)
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
	anagrams := anagramList(words)

	AM := &AnagramMap{anagrams}
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

	//fmt.Printf("%#v\n", AM.Mapping["bnorw"])
	sentence := [...]string{"The", "quick", "brown", "fox", "jumps", "over", "the", "lazy", "dog"}
	s := sentence[:]
	anasent := AM.AnagramSentence(s)
	if anasent[len(anasent)-1] != "god" {
		fmt.Printf("returned sentence: %#v\n", anasent)
		t.Errorf("'god' should be replaced with a 'dog'")
	}
}

func TestAnagramMapCaps(t *testing.T) {
	words := []string{"HihI", "America!", "fatcamp", "WHO", "how"}
	anagrams := anagramList(words)
	AM := AnagramMap{Mapping: anagrams}

	if len(AM.Mapping["how"]) != 2 {
		t.Errorf("'WHO' not correctly lowercased")
	}
	if AM.Mapping["aacfmpt"]["fatcamp"] != true {
		t.Errorf("Basic anagram mapping failed: %#v", AM.Mapping)
	}
}
