// MIT License
// Author: Umesh Patil, Neosemantix, Inc.

package main

import (
	"container/list"
	"errors"
	"fmt"
	"math"
	"strings"
)

type tuple struct {
	smallerChar rune
	biggerChar  rune
}

func (t tuple) strRep() string {
	return "(" + string(t.smallerChar) + ", " + string(t.biggerChar) + ")"
}

type orderedChars struct {
	charList *list.List
}

func newOc() *orderedChars {
	result := new(orderedChars)
	result.charList = list.New()
	return result
}

func newOcForTuple(t tuple) *orderedChars {
	result := newOc()
	first := result.charList.PushFront(t.smallerChar)
	result.charList.InsertAfter(t.biggerChar, first)
	return result
}

func (oc *orderedChars) strRep() string {
	result := ""
	mark := oc.charList.Front()
	for mark != nil {
		result = result + string(mark.Value.(rune))
		mark = mark.Next()
	}
	result = result + ""
	return result
}

func (oc *orderedChars) frontChar() rune {
	return oc.charList.Front().Value.(rune)
}

func (oc *orderedChars) backChar() rune {
	return oc.charList.Back().Value.(rune)
}

// Characters from front or back.
// If 'fromFront' is false; characters come from back.
// If there are not sufficient characters; will be returned as many as available
func (oc *orderedChars) getChars(howMany int, fromFront bool) []rune {
	var result []rune
	i := 0
	if fromFront {
		mark := oc.charList.Front()
		for mark != nil && i < howMany {
			result = append(result, mark.Value.(rune))
			mark = mark.Next()
			i++
		}
	} else {
		mark := oc.charList.Back()
		for mark != nil && i < howMany {
			result = append(result, mark.Value.(rune))
			mark = mark.Prev()
			i++
		}
		l := len(result)
		result2 := make([]rune, l)
		for j := 0; j < l; j++ {
			result2[l-j-1] = result[j]
		}
		result = result2
	}
	return result
}

func (oc *orderedChars) isEmpty() bool {
	if oc.charList.Len() > 0 {
		return true
	} else {
		return false
	}
}

func (oc *orderedChars) insertAtStart(v interface{}) {
	start := oc.charList.Front()
	oc.charList.InsertBefore(v, start)
}

func (oc *orderedChars) appendElement(v interface{}) {
	end := oc.charList.Back()
	if end != nil {
		oc.charList.InsertAfter(v, end)
	} else {
		// it is an empty list...
		oc.charList.PushFront(v)
	}
}

// Found value, second returned value:
//					-1:	contradiction
//					0:	no smaller or bigger char of tuple was found
//					1:	only smaller of tuple was found
//					2:	only bigger of tuple was found
//					3:	both smaller and bigger were found
//					11:	Smaller appended at the start
//					12: Bigger appended at the start
func (oc *orderedChars) consume(t tuple) (*orderedChars, int, error) {
	found := -1
	if t.biggerChar == oc.frontChar() {
		oc.insertAtStart(t.smallerChar)
		found = 11
		return nil, found, nil
	} else if t.smallerChar == oc.backChar() {
		oc.appendElement(t.biggerChar)
		found = 12
		return nil, found, nil
	} else {
		return oc.digest(t)
	}
}

func (oc *orderedChars) merge(t *orderedChars) bool {
	result := false
	if t.backChar() == oc.frontChar() {
		tt := t.clone()
		le := tt.charList.Back()
		tt.charList.Remove(le)
		oc.charList.PushFrontList(tt.charList)
		result = true
	} else if strings.EqualFold(string(t.getChars(2, false)), string(oc.getChars(2, true))) {
		tt := t.clone()
		le := tt.charList.Back()
		le2 := le.Prev()
		tt.charList.Remove(le)
		tt.charList.Remove(le2)
		oc.charList.PushFrontList(tt.charList)
		result = true
	} else if t.frontChar() == oc.backChar() {
		tt := t.clone()
		le := tt.charList.Front()
		tt.charList.Remove(le)
		oc.charList.PushBackList(tt.charList)
		result = true
	} else if strings.EqualFold(string(t.getChars(2, true)), string(oc.getChars(2, false))) {
		tt := t.clone()
		le := tt.charList.Front()
		le2 := le.Next()
		tt.charList.Remove(le)
		tt.charList.Remove(le2)
		oc.charList.PushBackList(tt.charList)
		result = true
	}
	return result
}

func (oc *orderedChars) digest(t tuple) (*orderedChars, int, error) {
	found := -1
	var result *orderedChars = nil
	var err error
	mark := oc.charList.Front()
	var sc, bc *list.Element = nil, nil
	for mark != nil {
		if mark.Value.(rune) == t.smallerChar {
			if bc != nil {
				// we got bigger character before the smaller character
				err = errors.New(fmt.Sprintf("Contrdiction for tuple %v", t))
			}
			sc = mark
		} else if mark.Value.(rune) == t.biggerChar {
			bc = mark
		}
		mark = mark.Next()
	}
	if sc != nil {
		if bc != nil {
			// we found both chars of tuple in the OC, nothing to do
			found = 3
		} else {
			result = oc.cloneSlice(oc.charList.Front(), sc)
			//result.appendElement(bc)
			result.appendElement(t.biggerChar)
			found = 1
		}
	} else {
		if bc != nil {
			result = oc.cloneSlice(bc, oc.charList.Back())
			//result.insertAtStart(sc.Value)
			result.insertAtStart(t.smallerChar)
			found = 2
		} else {
			// both sc and bc were not found, we return
			found = 0
		}
	}
	return result, found, err
}

func (oc orderedChars) cloneSlice(start *list.Element, end *list.Element) *orderedChars {
	result := newOc()
	mark := start
	for mark != end {
		result.appendElement(mark.Value)
		mark = mark.Next()
	}
	// last element is not added, we want to have it added as well
	result.charList.PushBack(mark.Value)
	return result
}

func (oc orderedChars) clone() *orderedChars {
	return oc.cloneSlice(oc.charList.Front(), oc.charList.Back())
}

func (oc orderedChars) size() int {
	return oc.charList.Len()
}

type ocList struct {
	ocs *list.List
}

func newOcl() *ocList {
	result := new(ocList)
	result.ocs = list.New()
	return result
}

func (ocl *ocList) strRep() string {
	result := "{"
	mark := ocl.ocs.Front()
	for mark != nil {
		oc := mark.Value.(*orderedChars)
		result = result + oc.strRep() + " "
		mark = mark.Next()
	}
	result = result + "}"
	return result
}

func (ocl *ocList) insertOc(oc *orderedChars) {
	mark := ocl.ocs.Front()
	found := false
	for !found && mark != nil {
		if (mark.Value.(*orderedChars)).size() > oc.size() {
			mark = mark.Next()
		} else {
			found = true
		}
	}
	if mark == nil {
		ocl.ocs.PushBack(oc) // empty case as well as end of the list
	} else {
		ocl.ocs.InsertBefore(oc, mark)
	}
}

func (ocl *ocList) sort() *ocList {
	result := newOcl()
	mark := ocl.ocs.Front()
	for mark != nil {
		result.insertOc(mark.Value.(*orderedChars))
		mark = mark.Next()
	}
	return result
}

func (ocl *ocList) add(extra *ocList) {
	mark := extra.ocs.Front()
	for mark != nil {
		ocl.insertOc(mark.Value.(*orderedChars))
		mark = mark.Next()
	}
}

func (ocl *ocList) remove(ss *list.Element) *orderedChars {
	return ocl.ocs.Remove(ss).(*orderedChars)
}

func (ocl *ocList) consumeATuple(t tuple) (*ocList, error) {
	var err error
	extraOcl := newOcl()
	subsumed := false
	notfound := true
	mark := ocl.ocs.Front()
	for mark != nil {
		anOc := mark.Value.(*orderedChars)
		oc, found, err := anOc.consume(t)
		if err != nil {
			break
		}
		if oc != nil {
			extraOcl.ocs.PushBack(oc)
		}
		switch found {
		case 0:
			notfound = notfound && true
		case 3:
			subsumed = subsumed || true
		default:
			notfound = notfound && false
		}
		mark = mark.Next()
	}
	if notfound && !subsumed {
		extraOcl.ocs.PushBack(newOcForTuple(t))
	}
	return extraOcl, err
}

func (ocl *ocList) consume(ts []tuple, consolidate bool) {
	start := 0
	if ocl.ocs.Len() == 0 {
		// nothing is added so far, we simply absorb all tuples
		t := ts[0]
		oc := newOcForTuple(t)
		ocl.insertOc(oc)
		start = 1
	}
	for _, t := range ts[start:] {
		extra, err := ocl.consumeATuple(t)
		if err != nil {
			return
		} else {
			ocl.add(extra)
		}
		if consolidate {
			//fmt.Printf("Before consolidate: %s\n", ocl.strRep())
			ocl.consolidate()
		}
	}
}

func (ocl *ocList) removeSubsumed() {
	mark := ocl.ocs.Back()
	for mark != nil {
		if ocl.canBeSubsumed(mark.Value.(*orderedChars)) {
			//ocl.remove(mark.Value.(*orderedChars))
			ocl.remove(mark)
		}
		mark = mark.Prev()
	}
}

func (ocl *ocList) consolidate() int {
	startSize := ocl.ocs.Len()
	// first remove those which are subsumed
	ocl.removeSubsumed()
	var toBeDeleted []*list.Element
	front := ocl.ocs.Front()
	start := ocl.ocs.Back()
	for start != nil && start != front {
		mark := start.Prev()
		merged := false
		for mark != nil {
			// we merge it in any as many orderedChars as applicable
			merged = mark.Value.(*orderedChars).merge(start.Value.(*orderedChars))
			mark = mark.Prev()
		}
		if merged {
			toBeDeleted = append(toBeDeleted, start)
		}
		start = mark
	}
	for tbd := range toBeDeleted {
		ocl.remove(toBeDeleted[tbd])
	}
	// now that many orderedChars are merged, we run removal of subsumed again
	ocl.removeSubsumed()
	// sort again
	ocl = ocl.sort()
	endSize := ocl.ocs.Len()
	return startSize - endSize
}

func (ocl *ocList) canBeSubsumed(oc *orderedChars) bool {
	result := false
	length := oc.size()
	if length > 0 {
		mark := ocl.ocs.Back()
		for mark != nil && !result {
			if mark.Value != oc && mark.Value.(*orderedChars).size() >= length {
				result = firstContainedInSecond(oc, mark.Value.(*orderedChars))
			}
			// else we really not need to check things
			mark = mark.Prev()
		}
	} else {
		result = true
	}
	return result
}

func generateTuples(smaller rune, biggers []rune) []tuple {
	var tlist []tuple
	for _, v := range biggers {
		t := new(tuple)
		t.smallerChar = smaller
		t.biggerChar = v
		tlist = append(tlist, *t)
	}
	return tlist
}

func compare(w1 string, w2 string) error {
	matching := true
	i := 0
	for matching {
		if i < len(w1) && i < len(w2) && w1[i] == w2[i] {
			i++
		} else {
			matching = false
		}
	}
	if !matching && i < len(w1) && i < len(w2) {
		firstChar := rune(w1[i])
		secondChar := rune(w2[i])

		// We are add the error condition before accepting the incoming tuple.
		// If the reverse of the incoming tuple already exists, then we have an error.
		if isContradictory(firstChar, secondChar) {
			return errors.New("contradictory input")
		}

		arr := greaterList[firstChar]
		if arr != nil {
			greaterList[firstChar] = append(arr, secondChar)
		} else {
			arr = make([]rune, 1)
			arr[0] = secondChar
			greaterList[firstChar] = arr
		}
	}
	return nil
}

func isContradictory(fc rune, sc rune) bool {
	result := false
	crtChars := greaterList[sc]
	if crtChars != nil && len(crtChars) > 0 {
		for _, v := range crtChars {
			if v == fc {
				result = true
			}
		}
	}
	return result
}

// First returned value is first gen tuples while the second one is remaining
func getFirstGenTuples(allTuples []tuple) ([]tuple, []tuple) {
	var firstGen, others []tuple
	alienAlphabet.reset()
	for t := range allTuples {
		if alienAlphabet.isNewTuple(allTuples[t]) {
			firstGen = append(firstGen, allTuples[t])
			alienAlphabet.tupleConsumed(allTuples[t])
		} else {
			others = append(others, allTuples[t])
		}
	}
	return firstGen, others
}

// We assume the alphabet is frozen before this call, meaning no more any new
// letters are to be encountered as well as explit call to freeze the alphabet
// is invoked (that sets the indices of alphabet characters in a fixed order).
func firstContainedInSecond(oc1 *orderedChars, oc2 *orderedChars) bool {
	result := false
	oc1bm := alienAlphabet.bitmap(oc1)
	if (oc1bm & alienAlphabet.bitmap(oc2)) == oc1bm {
		// this implies that every bit which is '1' in oc1; it is '1' in oc2 as well
		result = true
	}
	return result
}

var greaterList map[rune][]rune
var ocLists *ocList
var alienAlphabet *alphabet
var isolatedLetters map[rune]bool

func reset() {
	greaterList = nil
	ocLists = nil
	alienAlphabet = nil
	isolatedLetters = nil
}

func findContainedWordPairs(words []string) map[string]string {
	result := make(map[string]uint32)
	abt := newAlphabet()
	for _, w := range words {
		abt.populate(w)
	}
	abt.freezeLetters()
	done := make([]bool, len(words))
	for i, start := range words {
		if !done[i] {
			for j, next := range words[i+1:] {
				if strings.HasPrefix(next, start) {
					remaining := next[len(start):]
					if len(remaining) > 0 {
						// when start key is absent, es is "" empty string
						es, found := result[start]
						if !found {
							result[start] = abt.bitmapForStr(remaining)
						} else {
							result[start] = es | abt.bitmapForStr(remaining)
						}
						done[j] = true
					}
					// else nothing to add present
				}
				// else we skip it
			}
			done[i] = true
		}
		// else it is already done
	}
	result2 := make(map[string]string)
	for k, v := range result {
		result2[k] = abt.strFromBitmap(v)
	}
	return result2
}

func alienOrder(words []string) string {
	result := ""
	alienAlphabet = newAlphabet()
	greaterList = make(map[rune][]rune)
	isolatedLetters = make(map[rune]bool)
	w := words[0]
	alienAlphabet.populate(w)
	for _, aw := range words[1:] {
		if strings.EqualFold(w, aw) {
			// this comparision does not yield anything; we keep these letters aside
			for _, r := range w {
				isolatedLetters[rune(r)] = true
			}
		} else {
			err := compare(w, aw)
			if err != nil {
				// invalid order we exit
				return ""
			}
			alienAlphabet.populate(aw)
		}
		w = aw
	}
	if len(greaterList) > 0 {
		// at least one comparison was available
		// We need to remove isolated letters which are present in greaterList.
		for k, _ := range isolatedLetters {
			if greaterList[k] != nil {
				delete(isolatedLetters, k)
			}
		}
		fmt.Println(dumpMap(greaterList))

		// also freeze the alphabet
		alienAlphabet.freezeLetters()
		fmt.Println(fmt.Sprintf("Alien alphabet characters encountered: %s", alienAlphabet.strRep()))

		var allTuples []tuple
		ocLists = newOcl()
		for r, rl := range greaterList {
			tuples := generateTuples(r, rl)
			allTuples = append(allTuples, tuples...)
		}
		fmt.Println(dumpTupleSlice(allTuples))
		firstGenTup, remainingTup := getFirstGenTuples(allTuples)
		fmt.Printf("firstGenTup: %s\n", dumpTupleSlice(firstGenTup))
		fmt.Printf("remainingTup: %s\n", dumpTupleSlice(remainingTup))
		// all tuples are expected to be distinct, so need of consolidating is not there
		ocLists.consume(firstGenTup, false)
		fmt.Printf("After consuming firstGenTup:%s\n", ocLists.strRep())
		ocLists.consume(remainingTup, true)
		fmt.Printf("After consuming remainingTup:%s\n", ocLists.strRep())
		reduced := ocLists.consolidate()
		for reduced > 0 {
			reduced = ocLists.consolidate()
		}
		fmt.Printf("After consolidation:%s\n", ocLists.strRep())

	} else {
		// We did not get any comparision between any two letters.
		// This is Ok only if there is only one alphabet. For other cases,
		// we will have to check further things.
		if len(alienAlphabet.runeBoolMap) == 1 {
			for k, _ := range alienAlphabet.runeBoolMap {
				return string(k)
			}
		} else {
			// No relationship for any letter pair and it is not a single
			// letter corner case. We need to check the next corner case
			// where we want to check whether any word is contained in
			// another one or not. If so, those pairs are treated with
			// relations are returned.
			constrs := findContainedWordPairs(words)
			if len(constrs) > 0 {
				result = ""
				for k, v := range constrs {
					result = result + v + k
				}
				return result
			} else {
				return ""
			}
		}
	}
	// reset alphabet so as we can track those already covered and left out
	alienAlphabet.reset()
	if ocLists != nil && ocLists.ocs != nil && ocLists.ocs.Len() > 0 {
		mark := ocLists.ocs.Front()
		for mark != nil {
			anOc := mark.Value.(*orderedChars)
			result = alienAlphabet.filter(anOc) + result
			mark = mark.Next()
		}
	}
	// any remaining isolated chars
	if len(isolatedLetters) > 0 {
		for k, _ := range isolatedLetters {
			if !alienAlphabet.runeBoolMap[k] {
				result = result + string(k)
				alienAlphabet.letterConsumed(k)
			}
		}
	}
	// if anything left in the alphabet
	for k, v := range alienAlphabet.runeBoolMap {
		if !v {
			result = result + string(k)
		}
	}
	// so that unit tests can run
	reset()
	return result
}

func dumpMap(m map[rune][]rune) string {
	result := ""
	for r, s := range m {
		result = result + string(r) + ": " + string(s) + "\n"
	}
	return result
}

func dumpTupleSlice(ts []tuple) string {
	result := "{"
	for _, t := range ts {
		result = result + t.strRep()
	}
	return result + "}"
}

type alphabet struct {
	runeBoolMap            map[rune]bool
	letterPositions        map[rune]uint32
	reverseLetterPositions map[uint32]rune
}

func newAlphabet() *alphabet {
	return &alphabet{make(map[rune]bool), make(map[rune]uint32), make(map[uint32]rune)}
}

func (set *alphabet) add(i rune) bool {
	_, found := set.runeBoolMap[i]
	result := found
	if found {
		// if found, the worry about the value and not success of finding operation
		result = set.runeBoolMap[i]
	}
	set.runeBoolMap[i] = true
	//return !found //False if it existed already
	return !result
}

func (set *alphabet) get(i rune) bool {
	_, found := set.runeBoolMap[i]
	return found //true if it existed already
}

func (set *alphabet) populate(w string) {
	for _, c := range w {
		set.add(c)
	}
}

func (set *alphabet) freezeLetters() {
	i := 0
	for k, _ := range set.runeBoolMap {
		twopower := uint32(math.Pow(2, float64(i))) // each letter for one bit position
		set.letterPositions[k] = twopower
		set.reverseLetterPositions[twopower] = k
		i++
	}
}

func (set *alphabet) bitmap(word *orderedChars) uint32 {
	var result uint32 = 0
	mark := word.charList.Front()
	for mark != nil {
		u := set.letterPositions[mark.Value.(rune)]
		result = result | u //bitwise addition
		mark = mark.Next()
	}
	return result
}

func (set *alphabet) bitmapForStr(word string) uint32 {
	var result uint32 = 0
	for _, r := range word {
		u := set.letterPositions[r]
		result = result | u
	}
	return result
}

func (set *alphabet) strFromBitmap(u uint32) string {
	result := ""
	for tp, r := range set.reverseLetterPositions {
		if (tp & u) == tp {
			// Bitwise AND of tp with u returns back u
			// meaning both are 1 bits which implies
			// letter associated with this power of two is present
			result = result + string(r)
		}
	}
	return result
}

func (set *alphabet) filter(word *orderedChars) string {
	result := ""
	mark := word.charList.Front()
	for mark != nil {
		if !set.runeBoolMap[mark.Value.(rune)] {
			result = result + string(mark.Value.(rune))
			set.runeBoolMap[mark.Value.(rune)] = true
		}
		mark = mark.Next()
	}
	return result
}

func (set *alphabet) strRep() string {
	result := "{"
	for k, v := range set.runeBoolMap {
		if v {
			result = result + string(k) + ", "
		}
	}
	result = result + "} {"
	if len(set.letterPositions) > 0 {
		for k, v := range set.letterPositions {
			result = result + "(" + fmt.Sprintf("%v = %d),", string(k), v)
		}
	}
	result = result + "}"
	return result
}

func (set *alphabet) reset() {
	for k, _ := range set.runeBoolMap {
		set.runeBoolMap[k] = false
	}
}

func (set *alphabet) isNewTuple(t tuple) bool {
	return !set.runeBoolMap[t.smallerChar] && !set.runeBoolMap[t.biggerChar]
}

func (set *alphabet) tupleConsumed(t tuple) {
	//set.runeBoolMap[t.smallerChar] = true
	//set.runeBoolMap[t.biggerChar] = true
	set.letterConsumed(t.smallerChar)
	set.letterConsumed(t.biggerChar)
}

func (set *alphabet) letterConsumed(r rune) {
	set.runeBoolMap[r] = true
}

func main() {

	input := []string{"ze", "yf", "xd", "wd", "vd", "ua", "tt", "sz", "rd", "qd", "pz", "op", "nw", "mt", "ln", "ko", "jm", "il", "ho", "gk", "fa", "ed", "dg", "ct", "bb", "ba"}
	fmt.Println(alienOrder(input)) // expected:		"zyxwvutsrqponmlkjihgfedcba"
	fmt.Println()
	reset()

	//input = []string{"wrt","wrtkj","wrtkjd"}
	//fmt.Println(alienOrder(input))		// expected:		djkrtw
	//fmt.Println()
	//reset()
	//
	//input = []string{"wrt","wrtkj"}
	//fmt.Println(alienOrder(input))		// expected:		jkrtw
	//fmt.Println()
	//reset()
	//
	//input = []string{"abc","ab"}
	//fmt.Println(alienOrder(input))		// expected:		 empty
	//fmt.Println()
	//reset()
	//
	//input = []string{"wrt", "wrf", "er", "ett", "rftt"}
	//fmt.Println(alienOrder(input))		// expected: wertf
	//fmt.Println()
	//reset()
	//
	//input = []string{"ac","ab","b"}
	//fmt.Println(alienOrder(input))		// expected: acb
	//fmt.Println()
	//reset()
	//
	//input = []string{"z","x"}
	//fmt.Println(alienOrder(input))		// expected: zx
	//fmt.Println()
	//reset()
	//
	//input = []string{"z","x","z"}
	//fmt.Println(alienOrder(input))		// expected: 		empty
	//fmt.Println()
	//reset()
	//
	//input = []string{"z","z"}
	//fmt.Println(alienOrder(input))		// expected: z
	//fmt.Println()
	//reset()
	//
	//input = []string{"zy","zx"}
	//fmt.Println(alienOrder(input))		// expected yxz
	//fmt.Println()
	//reset()
}
