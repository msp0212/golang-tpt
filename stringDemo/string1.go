package main

import (
    "fmt"
    "unicode/utf8"
)

func main() {
    
    // mohit singh in hindi language
    const name = "मोहित सिंह"

    // return number of raw bytes in the string
    fmt.Println("name of len = ", len(name), name)

    // print raw bytes as hex value
    for i := 0; i < len(name); i++ {
        fmt.Printf("%x ", name[i])
    }
    fmt.Println()

    // rune count (may not be same as number of raw bytes for utf8 chars)
    fmt.Println("Rune count in name", utf8.RuneCountInString(name))

    // print runes in the string
    for inx, runeVal := range name {
        fmt.Printf("rune = %#U at inx = %d\n", runeVal, inx);
    }
    
    // traverse the string and convert raw bytes to rune
    for inx, size := 0, 0; inx < len(name); inx += size {
		runeVal, runeSize := utf8.DecodeRuneInString(name[inx:])
		fmt.Printf("rune = %#U of size = %d at inx = %d\n", runeVal, runeSize, inx)
		size = runeSize
	}
}
