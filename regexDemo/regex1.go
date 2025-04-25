package main

import (
	"fmt"
	"regexp"
)

var pln = fmt.Println
var pf = fmt.Printf

func reMatchStringDemo(re *regexp.Regexp) {
	pln("MatchString: ", re.MatchString("foobar"))
	pln("MatchString: ", re.MatchString("foobaz"))
	pln("MatchString: ", re.MatchString("foobazbar"))
	pln("MatchString: ", re.MatchString(""))
}

func reFindDemo(re *regexp.Regexp) {
	s := "123-foobar-456-foobar"
	f := re.Find([]byte(s))
	//Find returns byte slice
	pf("Find: %q\n", f)
	//FindAll returns slice of byte slices
	fa := re.FindAll([]byte(s), -1)
	pf("FindAll: %q\n", fa)
}

func reMatchDemo(re *regexp.Regexp) {
	s := "foobar"
	pln("Match: ", re.Match([]byte(s)))
	pln("Match: ", re.Match([]byte("foobaz")))
}

func reFindStringDemo(re *regexp.Regexp) {
	s := "123-foobar-456-foobar"
	f := re.FindString(s)
	//returns byte slice
	pf("FindString: %q\n", f)
	//returns slice of strings
	fa := re.FindAllString(s, -1)
	pf("FindAllString: %q\n", fa)
}

func reFindSubmatchDemo(re *regexp.Regexp) {
	s := "123-foobazbar-456-foobar"
	f := re.FindSubmatch([]byte(s))
	//returns byte slice
	pf("FindSubmatch: %q\n", f)
	//returns slice of byte slices
	fa := re.FindAllSubmatch([]byte(s), -1)
	pf("FindAllSubmatch: %q\n", fa)
}

func reFindIndexDemo(re *regexp.Regexp) {
	s := "123-foobazbar-456-foobar"
	f := re.FindIndex([]byte(s))
	//returns byte slice
	pf("FindIndex: %d\n", f)
	//returns slice of byte slices
	fa := re.FindAllIndex([]byte(s), -1)
	pln("FindAllIndex:", fa)
}

func reDemo() {
	re, err := regexp.Compile("foo([a-z])*bar")
	if err != nil {
		pln("error in compiling regular expression")
		return
	}
	reMatchStringDemo(re)
	reMatchDemo(re)
	/*
		16 Find based method given by below regular expression:
		Find(All)?(String)?(Submatch)?(Index)?
	*/
	reFindDemo(re)
	reFindStringDemo(re)
	reFindSubmatchDemo(re)
	reFindIndexDemo(re)
}

func main() {
	reDemo()
}
