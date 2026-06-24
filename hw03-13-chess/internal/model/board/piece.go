package board

type Color uint8

const (
	White Color = iota
	Black
)

type PieceType uint8

const (
	None PieceType = iota
	Pawn
	Knight
	Bishop
	Rook
	Queen
	King
)

type Piece struct {
	Type PieceType
	Color Color
}

func (p Piece) String() string {
	switch p.Type {
	case Pawn:
		if p.Color == White {
			return "♙"
		}
		return "♟︎"
	case Knight:
		if p.Color == White {
            return "♘"
        }
        return "♞"
	case Bishop:
		if p.Color == White {
            return "♗"
        }
        return "♝"
	case Rook:
		if p.Color == White {
            return "♖"
        }
        return "♜"
	case Queen:
		if p.Color == White {
            return "♕"
        }
        return "♛"
	case King:
		if p.Color == White {
            return "♔"
        }
        return "♚"
	default:
		return "?"
	}
}

func NewPiece(t PieceType, c Color) *Piece {
	return &Piece{Type: t, Color: c}
}

func (p Piece) SanLetter() string {
	switch p.Type {
	case Knight:
        return "N"
	case Bishop:
        return "B"
	case Rook:
        return "R"
	case Queen:
        return "Q"
	case King:
        return "K"
	default:
		return ""
	}
}