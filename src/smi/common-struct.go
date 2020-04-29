/*
 * MIT License
 * Copyright (c) 2020. Neosemantix, Inc.
 * Author: Umesh Patil
 */

package smi

// Social Media Content - one unit
type SmContent struct {
	Message string
	Link    string
}

func NewSmContent(msg string, url string) *SmContent {
	sm := SmContent{msg, url}
	return &sm
}
