package chess

type Color int

const (
	ColorEmpty Color = iota
	White
	Black
	ColorCount
)

type PieceType int

const (
	PieceTypeEmpty = iota
	King
	Queen
	Rook
	Bishop
	Knight
	Pawn
	PieceTypeCount
)

type Piece int

const (
	PieceEmpty Piece = iota
	WhiteKing
	WhiteQueen
	WhiteRook
	WhiteBishop
	WhiteKnight
	WhitePawn
	BlackKing
	BlackQueen
	BlackRook
	BlackBishop
	BlackKnight
	BlackPawn
	PieceCount
)

func (p Piece) Rune() rune {
	if p == PieceEmpty || p == PieceCount {
		return ' '
	}
	return 0x2653 + rune(p)
}

func (pt PieceType) Piece(color Color) Piece {
	switch pt {
	case King:
		switch color {
		case White:
			return WhiteKing
		case Black:
			return BlackKing
		default:
			return PieceEmpty
		}
	case Queen:
		switch color {
		case White:
			return WhiteQueen
		case Black:
			return BlackQueen
		default:
			return PieceEmpty
		}
	case Rook:
		switch color {
		case White:
			return WhiteRook
		case Black:
			return BlackRook
		default:
			return PieceEmpty
		}
	case Bishop:
		switch color {
		case White:
			return WhiteBishop
		case Black:
			return BlackBishop
		default:
			return PieceEmpty
		}
	case Knight:
		switch color {
		case White:
			return WhiteKnight
		case Black:
			return BlackKnight
		default:
			return PieceEmpty
		}
	case Pawn:
		switch color {
		case White:
			return WhitePawn
		case Black:
			return BlackPawn
		default:
			return PieceEmpty
		}
	default:
		return PieceEmpty
	}
}

func PosColor2Piece(pos int, color Color, isPawn bool) Piece {
	if pos > 7 {
		return PieceEmpty
	}

	if isPawn {
		return PieceType(Pawn).Piece(color)
	}

	switch {
	case pos == 0 || pos == 7:
		return PieceType(Rook).Piece(color)
	case pos == 1 || pos == 6:
		return PieceType(Knight).Piece(color)
	case pos == 2 || pos == 5:
		return PieceType(Bishop).Piece(color)
	case pos == 3:
		return PieceType(Queen).Piece(color)
	case pos == 4:
		return PieceType(King).Piece(color)
	default:
		return PieceEmpty
	}
}
