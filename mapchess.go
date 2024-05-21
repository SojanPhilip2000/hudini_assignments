package chessboard

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File []bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
    count:=0
    //ivide file "a", "b", ...entu venelum aava
    for _,square := range cb[file]{
        if square{
            count++
        }
    }
    return count
	panic("Please implement CountInFile()")
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
    count:=0
    if rank > 0 && rank <= 8{
        for _,square1 := range(cb){
            if square1[rank-1]{
                count++
            }
    	}
    }
    return count
	panic("Please implement CountInRank()")
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
    return 64
	panic("Please implement CountAll()")
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
count:=0
 for _,square := range(cb){
      for _,occupied := range(square){
            if occupied{
                count++
            }
    	}
}
         return count
}
