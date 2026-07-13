package board

type Color uint8

const (
	NoneColor Color = iota
	White
	Black
)

type PieceType uint8

const (
	NoneType PieceType = iota
	Pawn
	Knight
	Bishop
	Rook
	Queen
	King
)

func (t PieceType) SanLetter() rune {
	switch t {
	case Pawn:
		return 'P'
	case Knight:
		return 'N'
	case Bishop:
		return 'B'
	case Rook:
		return 'R'
	case Queen:
		return 'Q'
	case King:
		return 'K'
	default:
		return '?'
	}
}

func SanLetterToPieceType(r rune) PieceType {
	switch r {
	case 'P':
		return Pawn
	case 'N':
		return Knight
	case 'B':
		return Bishop
	case 'R':
		return Rook
	case 'Q':
		return Queen
	case 'K':
		return King
	default:
		return NoneType
	}
}

type Piece struct {
	Type  PieceType
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

func (p Piece) SanLetter() rune {
	return p.Type.SanLetter()
}
