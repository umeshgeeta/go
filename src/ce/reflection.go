// MIT License
// Author: Umesh Patil, Neosemantix, Inc.

package main

// Reference:	https://gist.github.com/drewolson/4771479

// Simple exercise which explains how reflection works in GO how you can use it
// to read well-known configuration structure out of some other uber config
// structures.

import (
	"../../util"
	"fmt"
	"reflect"
)

type SomeCfgSet struct {
	SomeField   string
	LogSettings util.LoggingCfg
}

type Foo struct {
	FirstName string `tag_name:"tag 1"`
	LastName  string `tag_name:"tag 2"`
	Age       int    `tag_name:"tag 3"`
}

func (f *Foo) reflect() {
	val := reflect.ValueOf(f).Elem()
	reflectValue(val)
}

func reflectValue(val reflect.Value) {
	fmt.Printf("Total num of fields: %d\n", val.NumField())
	found := false
	var sv interface{} = nil
	for i := 0; !found && i < val.NumField(); i++ {
		valueField := val.Field(i)
		typeVal := val.Type()
		typeField := typeVal.Field(i)
		tag := typeField.Tag
		fmt.Printf("Type: %v,\t Field Name: %s,\t Field Type: %v,\t Field Value: %v,\t Tag Value: %s\n",
			typeVal, typeField.Name, typeField.Type, valueField.Interface(), tag.Get("tag_name"))
		if typeField.Type.String() == "util.LoggingCfg" {
			fmt.Println("Found!")
			sv = valueField.Interface()
		}
	}
	fmt.Printf("Found value as: %v\n", sv)
}

func main() {
	f := &Foo{
		FirstName: "Drew",
		LastName:  "Olson",
		Age:       30,
	}
	f.reflect()

	lc := &util.LoggingCfg{
		LogFileName:  "test.log",
		MaxSizeInMb:  5,
		Backups:      5,
		AgeInDays:    30,
		Compress:     false,
		LogOnConsole: true,
		DebugLog:     true,
	}

	reflectValue(reflect.ValueOf(lc).Elem())

	scs := &SomeCfgSet{
		SomeField:   "Some uber config",
		LogSettings: *lc,
	}

	reflectValue(reflect.ValueOf(scs).Elem())
}
