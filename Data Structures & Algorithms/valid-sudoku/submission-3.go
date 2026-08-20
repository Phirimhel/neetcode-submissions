import (
	"sync"
	"context"
)

func isValidSudoku(board [][]byte) bool {

	wg := sync.WaitGroup{}
	ch := make(chan struct{})
	ctx, cancel := context.WithCancel(context.Background())
	
	wg.Add(len(board) * 2)
	for i := range board {
		go func() {
			defer wg.Done()
			select {
					case <-ctx.Done():
					return
					default:
				}
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
		}()
	}

	for i := range board {
		go func() {
			defer wg.Done()
			select {
					case <-ctx.Done():
					return
					default:
				}
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
		}()
	}

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


	

	for range ch {
		cancel()
		return false
	}

	return true
	}

