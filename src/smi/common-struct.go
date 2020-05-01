/*
 * MIT License
 * Copyright (c) 2020. Neosemantix, Inc.
 * Author: Umesh Patil
 */

package smi

type FbAuthCred struct {
	ObjectId string
	UserName string
	Password string
}

type CommonContent struct {
	Caption  string
	ImageUrl string
	Link     string
}

// Social Media Content - one unit
type PayloadToPublish struct {
	FbAuthCred FbAuthCred
	Content    CommonContent
}

func NewPayloadToPublish(fbc FbAuthCred, content CommonContent) *PayloadToPublish {
	sm := PayloadToPublish{fbc, content}
	return &sm
}

func NewFbAuthCred(oid string, user string, pword string) *FbAuthCred {
	fb := FbAuthCred{oid, user, pword}
	return &fb
}

func NewCommonContent(cap string, iurl string, link string) *CommonContent {
	cc := CommonContent{cap, iurl, link}
	return &cc
}
