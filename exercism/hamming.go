package hamming

import "errors"

func Distance(a, b string) (hammingDistance int, err error) {
    if len(a) != len(b) {
        return 0, errors.New("Sizes not equal")
    }

    for i := range a {
        if (a[i] != b[i]){
            hammingDistance++
        }
    }

    return
}

