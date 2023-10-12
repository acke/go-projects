package chessboard

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File []bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map [string] File


// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	numberOfOccupiedSquares := 0

    for _, squareOccupied := range cb[file] {
        if squareOccupied {
            numberOfOccupiedSquares++
        }
    }

    return numberOfOccupiedSquares
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	numberInRank := 0

    if rank > 8 || rank < 1 {
        return 0
    }

    for _, f := range cb {
        if f[rank - 1] {
            	numberInRank++
        }
    }

    return numberInRank
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
    numberOfSquaresOnBoard := 0

    for _, f := range cb {
        numberOfSquaresOnBoard += len(f)
    }

    return numberOfSquaresOnBoard
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	totalOccupied := 0

    for _, file := range cb {
        for _, square := range file {
            if square {
                totalOccupied++
            }
        }
    }

    return totalOccupied
}

