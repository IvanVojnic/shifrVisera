package main

import (
	"fmt"
	"math"
	"os"
)

func getAlphabet() []rune {
	var alphabet = []rune{'а', 'б', 'в', 'г', 'д', 'е', 'ё', 'ж', 'з', 'и', 'й', 'к', 'л', 'м', 'н', 'о', 'п', 'р', 'с', 'т', 'у', 'ф', 'х', 'ц', 'ч', 'ш', 'щ', 'ъ', 'ы', 'ь', 'э', 'ю', 'я'}
	return alphabet
}

var AlphMap = map[rune]int{
	'а': 1,
	'б': 2,
	'в': 3,
	'г': 4,
	'д': 5,
	'е': 6,
	'ж': 7,
	'з': 8,
	'и': 9,
	'й': 10,
	'к': 11,
	'л': 12,
	'м': 13,
	'н': 14,
	'о': 15,
	'п': 16,
	'р': 17,
	'с': 18,
	'т': 19,
	'у': 20,
	'ф': 21,
	'х': 22,
	'ц': 23,
	'ч': 24,
	'ш': 25,
	'щ': 26,
	'ъ': 27,
	'ы': 28,
	'ь': 29,
	'э': 30,
	'ю': 31,
	'я': 32,
	' ': 33,
}

func main() {
	word := "шифртекст"
	key := "иею"
	shifr := ShifrVisera(word, key)
	fmt.Println(shifr)

	fmt.Println(len(shifr))

	fmt.Println(string(shifr))

	fmt.Print("Insert word: ")
	fmt.Fscan(os.Stdin, &word)

	fmt.Print("Insert key: ")
	fmt.Fscan(os.Stdin, &key)
	rootLen := math.Sqrt(float64(len(word)))
	for {
		if math.Round(rootLen) != float64(len(word)) {
			break
		}
		word = word + " "
	}
	wordIndex := getArrayOfIndex(word)
	rootLenArr := math.Round(float64(len(wordIndex)))
	var
	keyIndex := getArrayOfIndex(key)
}

func getArrayOfIndex(word string) []int {
	var arr []int
	for _, symbol := range word {
		arr = append(arr, AlphMap[symbol])
	}
	return arr
}

func getIndexOfCurrentSymbol(symbol rune) int {
	alph := getAlphabet()
	for i := 0; i < len(alph); i++ {
		if alph[i] == symbol {
			return i
		}
	}
	return -1
}

func ShifrVisera(word string, key string) []rune {
	wordArr := []rune(word)
	keyArr := []rune(key)
	shifrArr := []rune("")
	keyIndex := 0
	for _, symbol := range wordArr {
		wordArrIndex := getIndexOfCurrentSymbol(symbol)
		keyArrIndex := getIndexOfCurrentSymbol(keyArr[keyIndex])
		newIndex := (wordArrIndex + keyArrIndex) % 33
		shifrSymbol := getAlphabet()[newIndex]
		shifrArr = append(shifrArr, shifrSymbol)
		keyIndex++
		if keyIndex == len(keyArr) {
			keyIndex = 0
		}
	}
	return shifrArr
}
