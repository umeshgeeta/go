// MIT License
// Author: Umesh Patil, Neosemantix, Inc.

package smi

type FbAuthCred struct {
	ObjectId string
	UserName string
	Password string
}

type TwitterCred struct {
	User     string
	Password string
}

type CommonContent struct {
	Caption  string
	ImageUrl string
	Link     string
}

// Social Media Content - one unit
type PayloadToPublish struct {
	FbAuthCred  FbAuthCred
	TwitterCred TwitterCred
	Content     CommonContent
}

type FbPayload struct {
	FbAuthCred FbAuthCred
	Content    CommonContent
}

type TwitterPayload struct {
	TwitterCred TwitterCred
	Content     CommonContent
}

func NewPayloadToPublish(fbc FbAuthCred, content CommonContent) *PayloadToPublish {
	var emptyTwitterCred TwitterCred
	sm := PayloadToPublish{fbc, emptyTwitterCred, content}
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
