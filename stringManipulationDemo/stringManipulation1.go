package main

import (
	"fmt"
	str "strings"
)

var pln = fmt.Println

func stringManipulationsDemo() {
	pln("stringManipulationsDemo")

	pln("Count:", str.Count("foobar", "o"))
	pln("Count:", str.Count("foobar", "")) //before and after each rune

	cutDemo := func(s, sep string) {
		before, after, found := str.Cut(s, sep)
		pln("Cut:", "found = ", found, "before = ", before, "after = ", after)
	}
	cutDemo("foobar", "bar")
	cutDemo("foobar", "foo")
	cutDemo("foobar", "oo")
	cutDemo("foobar", "baz")

	pln("Fields: ", str.Fields(" foo		bar baz             "))

	pln("Compare:", str.Compare("foobar", "bar"))
	pln("Compare:", str.Compare("foobar", "foobar"))

	pln("Contains:", str.Contains("foobar", "bar"))

	pln("HasPrefix:", str.HasPrefix("foobar", "foo"))

	pln("HasSuffix:", str.HasSuffix("foobar", "bar"))

	pln("Index:", str.Index("foobar", "bar"))
	pln("Index:", str.Index("foobar", "foo"))
	pln("Index:", str.Index("foobar", "boo"))

	pln("Join:", str.Join([]string{"foo", "bar", "baz"}, "|"))
	pln("Join:", str.Join([]string{"foo", "bar", "baz"}, "&&"))
	pln("Join:", str.Join([]string{"foo", "bar", "baz"}, ""))

	pln("Repeat:", str.Repeat("foo", 0))
	pln("Repeat:", str.Repeat("foo", 1))
	pln("Repeat:", str.Repeat("foo", 2))
	pln("Repeat:", str.Repeat("f", 2))

	pln("Replace:", str.Replace("foo foo foo", "foo", "bar", 2))
	pln("Replace:", str.Replace("foo foo foo", "foo", "bar", -1))
	pln("Replace:", str.Replace("banana", "ana", "aka", -1))

	pln("Split:", str.Split("foo,bar,baz", ","))
	pln("Split:", str.Split(" foo,bar,baz", ""))
	pln("Split:", str.Split("", ","))
	pln("Split:", str.Split("", ""))
	pln("Split:", str.Split("foo bar foo baz foo bat", "foo "))

	pln("Trim:", str.Trim(",|&,foo|bar&baz|&-=", ",|&=-"))

	pln("ToUpper:", str.ToUpper("FooBar"))
	pln("ToLower:", str.ToLower("FooBar"))
}

func main() {
	stringManipulationsDemo()
}
