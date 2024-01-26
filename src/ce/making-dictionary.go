
package main

import ( 
  "fmt"
  "sort"
)

var dataset = []string{
  "alerting", "altering", "integral", "relating", "triangle",
  "post", "pots", "spot", "stop", "tops", // "opst"
  // "dormitory", "dirty room",
  // "stone age", "stage one",
}

var dictionary map[string][]string

func main() {
  
  buildDictionary();

  // loop to serve the words...

}

func buildDictionary(){
  totalWords := len(dataset)
  dictionary = make(map[string][]string)
  i := 0
  for i < totalWords {
    orderedWord := getAlphabeticalOrder(dataset[i])
    anagrams, ok := dictionary[orderedWord]
    if !ok {
      anagrams := make([]string)
    }
    anagrams = append(anagrams, dataset[i])
    dictionary[orderedWord] = anagrams
  }
}

func findAnagrams(word string)[]string {
  return dictionary[word]
}


func getAlphabeticalOrder(word string) string {
  var r []rune
  for _, runeValue := range word {
    r = append(r, runeValue)
  }
  sort.Sort(r)
  return string(r)
}

func isAnagram(input string, word string)bool {
  wordMap := make(map[byte]int)
  ba := []byte(word)
  wordLen := len(ba)
  i := 0
  for i < wordLen {
    b := ba[i]
    count, ok := wordMap[b]
    if !ok {
      wordMap[b] = 1
    } else {
      wordMap[b] = count+1
    }
    i++
  }
  
  ia := []byte(input)
  inputLen := len(ia)
  
  if inputLen != wordLen {
    return false
  }
  
  i = 0
  for i < inputLen {
    a := ia[i]
    count, ok := wordMap[a]
    if ok {
      if count == 1 {
        delete(wordMap, a)
      } else {
        wordMap[a] = count-1
      }
    } else {
      return false
    }
    i++
  }
  
  if len(wordMap) == 0 {
    return true
  }
  return false
  
}

