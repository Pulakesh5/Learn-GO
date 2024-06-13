package chessboard

// import "fmt"
// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type File []bool
type Chessboard map[string] File 
// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
    count := 0
	for _, cell := range cb[file] {
		if(cell) {
            count++
        }
    }
    return count
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	if(rank<1 || rank>8) {
        return 0
    } else {
		count := 0
    	for _, file := range cb {
            if(file[rank-1]) {
                count++
            }
        }
        return count
    }   
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	countSquare :=0 
    for _, file := range cb {
        countSquare+=len(file)
    }
    return countSquare
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	countOccupied := 0
    for _, file := range cb {
		for _, cell := range file {
            if(cell) {
                countOccupied++
            }
        }
    }
    return countOccupied
}
