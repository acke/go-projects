package railfence

import "fmt"

func Encode(message string, rails int) string {

    rail := 0
    movingDown := true
    fence := make([]string, rails)

    for _, letter := range message {
        fmt.Println("Letter:", string(letter))
        fence[rail] += string(letter)
        fmt.Println(fence)

        switch {
            case rail == 0:
                movingDown = true
                rail++
            case rail < rails && movingDown:
                fmt.Println(fence[rail])
                if rail == (rails - 1) {
                    rail--
                    movingDown = false
                    continue
                }
                rail++
        	case rail < rails && !movingDown:
                rail--
        }
    }

	var result string
	for _, slice := range fence {
		result += slice
	}
	return result
}



func Decode(message string, rails int) string {
	panic("Please implement the Decode function")
}

