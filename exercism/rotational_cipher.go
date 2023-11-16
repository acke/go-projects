package rotationalcipher

import (
    "strings"
    "unicode"
)

func RotationalCipher(plain string, shiftKey int) (deciphered string) {
    lowercase := "abcdefghijklmnopqrstuvwxyz"
    uppercase := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

    for _, r := range plain {
        switch{
            case !unicode.IsLetter(r) :
                deciphered = deciphered + string(r)
            case unicode.IsLower(r):
                deciphered = deciphered + DecipherLetter(r, shiftKey, lowercase)
            case unicode.IsUpper(r):
                deciphered = deciphered + DecipherLetter(r, shiftKey, uppercase)
        }
    }
    return
}

func DecipherLetter (letter rune, shiftKey int, letters string) string {
    cipheredKeyIndex := strings.Index(letters, string(letter))

    if (shiftKey + cipheredKeyIndex >= 26){
        return string (letters[shiftKey + cipheredKeyIndex - 26])
    }

    return string (letters[shiftKey + cipheredKeyIndex])
}

