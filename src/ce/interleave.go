/*
 * MIT License
 * Copyright (c) 2023. Neosemantix, Inc.
 * Author: Umesh Patil
 */

// MIT License
// Author: Umesh Patil, Neosemantix, Inc.

package main

type state struct {
	s3index int
	s1index int
	s2index int
}

type pastDecisions struct {
	decisions []state
}

func (ds *pastDecisions) pushState(i1, i2, i3 int) state {
	st := state{
		s3index: i3,
		s1index: i1,
		s2index: i2,
	}
	ds.decisions = append(ds.decisions, st)
	return st
}

func (ds *pastDecisions) popState() *state {
	dl := len(ds.decisions)
	if dl == 0 {
		return nil
	}
	result := ds.decisions[dl-1]
	ds.decisions = ds.decisions[:dl-1]
	return &result
}

func isInterleave1Cut(s1 string, s2 string, s3 string) bool {
	sr1 := []rune(s1)
	sr2 := []rune(s2)
	sr3 := []rune(s3)
	l1 := len(sr1)
	l2 := len(sr2)
	l3 := len(sr3)
	if l1+l2 != l3 {
		return false
	}
	if l3 == 0 { //trivially
		return true
	}
	if l1 == 0 {
		return s2 == s3
	}
	if l2 == 0 {
		return s1 == s3
	}
	dcs := make([]state, 0)
	pdcs := pastDecisions{
		decisions: dcs,
	}
	i, j, k := 0, 0, 0
	lastj := -1
	useS2 := false
	for k < l3 {
		if j < l2 && sr2[j] != sr3[k] {
			if i < l1 && sr1[i] == sr3[k] {
				// s1 & s3 are matching
				// we consume & move forward
				i++
				k++
			} else {
				lastState := pdcs.popState()
				if lastState == nil {
					// nothing to backtrack and we are not matching
					return false
				}
				i = lastState.s1index
				j = lastState.s2index
				k = lastState.s3index
				useS2 = true
			}
		} else {
			// s2 & s3 are equal
			if useS2 {
				// this is of bscktracked path
				j++
				k++
				useS2 = false
			} else {
				if i < l1 && sr1[i] == sr3[k] {
					// both are matching with s3, we pick up s1 and track it
					if lastj != j {
						pdcs.pushState(i, j, k)
						lastj = j
					}
					// and now consume s1
					i++
					k++
				} else {
					// only s2 is matching, we consume that
					j++
					k++
				}
			}
		}
	}
	return true
}
