package main

func isValidSudoku(board [][]byte) bool {
	var count int
	for i := 0; i < 9; i += 3 {
		for j := 0; j < 9; j += 3 {
			res, empty := checkBox(board, i, j)
			if !res {
				return false
			}
			if empty {
				count++
			}
		}
	}
	if count == 9 {
		return true
	}
	if count > 0 {
		return false
	}
	return true
}

func checkBox(board [][]byte, i, j int) (bool, bool) {
	boardMap := make(map[byte]byte)
	var m int = 0
	for p := i; m < 3; p++ {
		m++
		var n int = 0
		for k := j; n < 3; k++ {
			n++
			if string(board[p][k]) != "." {
				if _, found := boardMap[board[p][k]]; found {
					return false, false
				} else {
					boardMap[board[p][k]]++
				}
			}
		}

	}
	if len(boardMap) == 0 {
		return true, true
	}
	return true, false
}
