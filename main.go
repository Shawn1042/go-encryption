package main

import (
    "fmt"
    "strings"
)

const alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

// shiftLetters shifts the letters of the alphabet by the given key.
func shiftLetters(key int) string {
    return alphabet[key:] + alphabet[:key]
}

// transform applies a character-wise transformation defined by the mapping function.
func transform(text string, mapper func(rune) rune) string {
    var result strings.Builder
    for _, r := range text {
        result.WriteRune(mapper(r))
    }
    return result.String()
}

func encrypt(key int, plainText string) string {
    shifted := shiftLetters(key)
    return transform(plainText, func(r rune) rune {
        if pos := strings.IndexRune(alphabet, r); pos != -1 {
            return rune(shifted[pos])
        }
        return r
    })
}

func decrypt(key int, encryptedText string) string {
    shifted := shiftLetters(key)
    return transform(encryptedText, func(r rune) rune {
        if pos := strings.IndexRune(shifted, r); pos != -1 {
            return rune(alphabet[pos])
        }
        return r
    })
}

func main() {
    plainText := "I LOVE CODING!!"
    fmt.Println("Plain Text:", plainText)
    encrypted := encrypt(5, plainText)
    fmt.Println("Encrypted text:", encrypted)
    decrypted := decrypt(5, encrypted)
    fmt.Println("Decrypted:", decrypted)
}
