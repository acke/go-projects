package atbash

import (
    "unicode"
    "unicode/utf8"
)

func Atbash(s string) (ciphered string) {
    cipher:= "zyxwvutsrqponmlkjihgfedcba"
    counter:= 0
    runeKey := 97

    for i, a := range s {
       	if unicode.IsNumber(a) {
            ciphered = ciphered + string (a)
            counter++
        }
        if unicode.IsLetter(a) {
            cipherKey := int(unicode.ToLower(a)) - runeKey
            ciphered = ciphered + string (cipher[cipherKey])
            counter++
        }
        if counter == 5 && utf8.RuneCountInString(s) > i+2 {
            ciphered = ciphered + " "
            counter = 0
        }
    }

    return ciphered
}

