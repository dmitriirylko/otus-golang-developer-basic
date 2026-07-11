package move

// func TestMoveCtors(t *testing.T) {
// 	m := NewMove(board.Square{X: 0, Y: 0}, board.Square{X: 1, Y: 1})
// 	assert.Equal(t, m.Flag, FlagQuiet)
// 	assert.True(t, m.isQuiet())
// 	assert.False(t, m.isCapture())
// 	m1 := m.WithCapture()
// 	assert.Equal(t, m1.Flag, FlagCapture)
// 	assert.False(t, m1.isQuiet())
// 	assert.True(t, m1.isCapture())
// }

// func TestMoveToSan(t *testing.T) {
// 	// TODO: увеличить количество тест-кейсов
// 	cases := []struct {
// 		name        string
// 		from        board.Square
// 		to          board.Square
// 		pieceType   board.PieceType
// 		withCapture bool
// 		expected    string
// 	}{
// 		{"1", board.Square{X: 0, Y: 1}, board.Square{X: 0, Y: 2}, board.Pawn, false, "a3"},
// 		{"2", board.Square{X: 6, Y: 6}, board.Square{X: 4, Y: 4}, board.Bishop, false, "Be5"},
// 		{"3", board.Square{X: 0, Y: 7}, board.Square{X: 3, Y: 7}, board.Rook, true, "Rxd8"},
// 		{"4", board.Square{X: 6, Y: 4}, board.Square{X: 5, Y: 5}, board.Pawn, true, "gxf6"},
// 	}

// 	for _, tc := range cases {
// 		t.Run(tc.name, func(t *testing.T) {
// 			piece := board.NewPiece(tc.pieceType, board.White)
// 			move := NewMove(tc.from, tc.to)
// 			if tc.withCapture == true {
// 				move = move.WithCapture()
// 			}
// 			san := move.ToSan(piece, nil)
// 			assert.Equal(t, san, tc.expected)
// 		})
// 	}
// }
