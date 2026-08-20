import "sync"


func isValidSudoku(board [][]byte) bool {

	ch := make(chan struct{})
	wg := sync.WaitGroup{}

	wg.Add(2) 
	go func() {
		defer wg.Done()
		for i := range board {
		
			m := make(map[byte]struct{})
			for j := range board[i] {

				if board[i][j] == '.' {
					continue 
				}

				if _, exists := m[board[i][j]]; exists {
					ch <- struct{}{}
					return 
				}

				m[board[i][j]] = struct{}{}
		}
	
	}
		}()

	go func() {
	defer wg.Done()
	for i := range board {
		
			m := make(map[byte]struct{})
			for j := range board {
				

				if board[j][i] == '.' {
					continue 
				}

				if _, exists := m[board[j][i]]; exists {
					ch <- struct{}{}
					return 
				}

			m[board[j][i]] = struct{}{}
		}	
		
	}
	}()

	
	go func() {
		wg.Wait()
		close(ch)
	}()


	var sb strings.Builder
	sBoard := make([]string, 9)

	sb.Grow(9)

	j := 3
	sBoardIdx := 0
	for i := 0; i < 9; i++ {
		sb.Write(board[i][:j])
		i++
		sb.Write(board[i][:j])
		i++
		sb.Write(board[i][:j])
		sBoard[sBoardIdx] = sb.String()
		sb.Reset()
		sBoardIdx++
	}

	j = 6
	for i := 0; i < 9; i++ {
		sb.Write(board[i][j-3 : j])
		i++
		sb.Write(board[i][j-3 : j])
		i++
		sb.Write(board[i][j-3 : j])
		sBoard[sBoardIdx] = sb.String()
		sb.Reset()
		sBoardIdx++
	}

	j = 9
	for i := 0; i < 9; i++ {
		sb.Write(board[i][j-3 : j])
		i++
		sb.Write(board[i][j-3 : j])
		i++
		sb.Write(board[i][j-3 : j])
		sBoard[sBoardIdx] = sb.String()
		sb.Reset()
		sBoardIdx++
	}

	for i := range sBoard {
		m := make(map[byte]struct{})
		for j := range sBoard[i] {

			if sBoard[i][j] == '.' {
				continue
			}

			if _, exists := m[sBoard[i][j]]; exists {
				return false
			}

			m[sBoard[i][j]] = struct{}{}
		}
	}


	

	_, ok := <-ch 
	if ok {
		return false
	}

	return true
	}

