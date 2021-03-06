package main


import "fmt"


type Color bool

const WHITE Color = true
const BLACK Color = false

type PieceType uint8

// it holds: WHITE_PIECE % 8 == PIECE == BLACK_PIECE % 8
const (
    NO_PIECE PieceType = iota
    KING // = 1
    QUEEN
    ROOK
    BISHOP
    KNIGHT
    PAWN // = 6
    _
    _
    WHITE_KING // = 9
    WHITE_QUEEN
    WHITE_ROOK
    WHITE_BISHOP
    WHITE_KNIGHT
    WHITE_PAWN // = 14
    _
    _
    BLACK_KING // = 17
    BLACK_QUEEN
    BLACK_ROOK
    BLACK_BISHOP
    BLACK_KNIGHT
    BLACK_PAWN // = 22
)

// typeOf takes a some piece and checks for piece type
func (pieceType PieceType) is(otherType PieceType) bool {
    return pieceType % 8 == otherType
}

func (pieceType PieceType) color() Color {
    s := pieceType / 8
    if s == 1 {
        return WHITE
    } else if s == 2 {
        return BLACK
    } else if pieceType == NO_PIECE {
        panic("Need a piece to check color")
    } else {
        panic("Need a white or a black piece to check color")
    }
}

var PIECES = []PieceType{
    WHITE_KING, WHITE_QUEEN, WHITE_ROOK, WHITE_BISHOP, WHITE_KNIGHT, WHITE_PAWN,
    BLACK_KING, BLACK_QUEEN, BLACK_ROOK, BLACK_BISHOP, BLACK_KNIGHT, BLACK_PAWN,
}
var WHITE_PIECES = []PieceType{
    WHITE_KING, WHITE_QUEEN, WHITE_ROOK, WHITE_BISHOP, WHITE_KNIGHT, WHITE_PAWN,
}
var BLACK_PIECES = []PieceType{
    BLACK_KING, BLACK_QUEEN, BLACK_ROOK, BLACK_BISHOP, BLACK_KNIGHT, BLACK_PAWN,
}

const FEN_WHITE_KING = "K"
const FEN_WHITE_QUEEN = "Q"
const FEN_WHITE_ROOK = "R"
const FEN_WHITE_BISHOP = "B"
const FEN_WHITE_KNIGHT = "N"
const FEN_WHITE_PAWN = "P"
const FEN_BLACK_KING = "k"
const FEN_BLACK_QUEEN = "q"
const FEN_BLACK_ROOK = "r"
const FEN_BLACK_BISHOP = "b"
const FEN_BLACK_KNIGHT = "n"
const FEN_BLACK_PAWN = "p"
const FEN_EMPTY = "x"


const ALGEBRAIC_WHITE_KING = "K"
const ALGEBRAIC_WHITE_QUEEN = "Q"
const ALGEBRAIC_WHITE_ROOK = "R"
const ALGEBRAIC_WHITE_BISHOP = "B"
const ALGEBRAIC_WHITE_KNIGHT = "N"
const ALGEBRAIC_WHITE_PAWN = ""
const ALGEBRAIC_BLACK_KING = "K"
const ALGEBRAIC_BLACK_QUEEN = "Q"
const ALGEBRAIC_BLACK_ROOK = "R"
const ALGEBRAIC_BLACK_BISHOP = "B"
const ALGEBRAIC_BLACK_KNIGHT = "N"
const ALGEBRAIC_BLACK_PAWN = ""

const HTML_WHITE_KING = "&#9812;"
const HTML_WHITE_QUEEN = "&#9813;"
const HTML_WHITE_ROOK = "&#9814;"
const HTML_WHITE_BISHOP = "&#9815;"
const HTML_WHITE_KNIGHT = "&#9816;"
const HTML_WHITE_PAWN = "" //"&#9817;"
const HTML_BLACK_KING = "&#9818;"
const HTML_BLACK_QUEEN = "&#9819;"
const HTML_BLACK_ROOK = "&#9820;"
const HTML_BLACK_BISHOP = "&#9821;"
const HTML_BLACK_KNIGHT = "&#9822;"
const HTML_BLACK_PAWN = "" //"&#9823;"

// U+2654
const UTF8_WHITE_KING = "♔"
const UTF8_WHITE_QUEEN = "♕"
const UTF8_WHITE_ROOK = "♖"
const UTF8_WHITE_BISHOP = "♗"
const UTF8_WHITE_KNIGHT = "♘"
const UTF8_WHITE_PAWN = "" //"♙"
const UTF8_BLACK_KING = "♚"
const UTF8_BLACK_QUEEN = "♛"
const UTF8_BLACK_ROOK = "♜"
const UTF8_BLACK_BISHOP = "♝"
const UTF8_BLACK_KNIGHT = "♞"
const UTF8_BLACK_PAWN = "" //"♟"


var FEN_TO_PIECE = map[string]PieceType{
    FEN_WHITE_KING: WHITE_KING,
    FEN_WHITE_QUEEN: WHITE_QUEEN,
    FEN_WHITE_ROOK: WHITE_ROOK,
    FEN_WHITE_BISHOP: WHITE_BISHOP,
    FEN_WHITE_KNIGHT: WHITE_KNIGHT,
    FEN_WHITE_PAWN: WHITE_PAWN,
    FEN_BLACK_KING: BLACK_KING,
    FEN_BLACK_QUEEN: BLACK_QUEEN,
    FEN_BLACK_ROOK: BLACK_ROOK,
    FEN_BLACK_BISHOP: BLACK_BISHOP,
    FEN_BLACK_KNIGHT: BLACK_KNIGHT,
    FEN_BLACK_PAWN: BLACK_PAWN,
    FEN_EMPTY: 0,
}

var PIECE_TO_FEN = map[PieceType]string{
    WHITE_KING: FEN_WHITE_KING,
    WHITE_QUEEN: FEN_WHITE_QUEEN,
    WHITE_ROOK: FEN_WHITE_ROOK,
    WHITE_BISHOP: FEN_WHITE_BISHOP,
    WHITE_KNIGHT: FEN_WHITE_KNIGHT,
    WHITE_PAWN: FEN_WHITE_PAWN,
    BLACK_KING: FEN_BLACK_KING,
    BLACK_QUEEN: FEN_BLACK_QUEEN,
    BLACK_ROOK: FEN_BLACK_ROOK,
    BLACK_BISHOP: FEN_BLACK_BISHOP,
    BLACK_KNIGHT: FEN_BLACK_KNIGHT,
    BLACK_PAWN: FEN_BLACK_PAWN,
}

var PIECE_TO_HTML = map[PieceType]string{
    WHITE_KING: HTML_WHITE_KING,
    WHITE_QUEEN: HTML_WHITE_QUEEN,
    WHITE_ROOK: HTML_WHITE_ROOK,
    WHITE_BISHOP: HTML_WHITE_BISHOP,
    WHITE_KNIGHT: HTML_WHITE_KNIGHT,
    WHITE_PAWN: HTML_WHITE_PAWN,
    BLACK_KING: HTML_BLACK_KING,
    BLACK_QUEEN: HTML_BLACK_QUEEN,
    BLACK_ROOK: HTML_BLACK_ROOK,
    BLACK_BISHOP: HTML_BLACK_BISHOP,
    BLACK_KNIGHT: HTML_BLACK_KNIGHT,
    BLACK_PAWN: HTML_BLACK_PAWN,
}

var FEN_TO_HTML = map[string]string{
    FEN_WHITE_KING: HTML_WHITE_KING,
    FEN_WHITE_QUEEN: HTML_WHITE_QUEEN,
    FEN_WHITE_ROOK: HTML_WHITE_ROOK,
    FEN_WHITE_BISHOP: HTML_WHITE_BISHOP,
    FEN_WHITE_KNIGHT: HTML_WHITE_KNIGHT,
    FEN_WHITE_PAWN: HTML_WHITE_PAWN,
    FEN_BLACK_KING: HTML_BLACK_KING,
    FEN_BLACK_QUEEN: HTML_BLACK_QUEEN,
    FEN_BLACK_ROOK: HTML_BLACK_ROOK,
    FEN_BLACK_BISHOP: HTML_BLACK_BISHOP,
    FEN_BLACK_KNIGHT: HTML_BLACK_KNIGHT,
    FEN_BLACK_PAWN: HTML_BLACK_PAWN,
}

var FEN_TO_UTF8 = map[string]string{
    FEN_WHITE_KING: UTF8_WHITE_KING,
    FEN_WHITE_QUEEN: UTF8_WHITE_QUEEN,
    FEN_WHITE_ROOK: UTF8_WHITE_ROOK,
    FEN_WHITE_BISHOP: UTF8_WHITE_BISHOP,
    FEN_WHITE_KNIGHT: UTF8_WHITE_KNIGHT,
    FEN_WHITE_PAWN: UTF8_WHITE_PAWN,
    FEN_BLACK_KING: UTF8_BLACK_KING,
    FEN_BLACK_QUEEN: UTF8_BLACK_QUEEN,
    FEN_BLACK_ROOK: UTF8_BLACK_ROOK,
    FEN_BLACK_BISHOP: UTF8_BLACK_BISHOP,
    FEN_BLACK_KNIGHT: UTF8_BLACK_KNIGHT,
    FEN_BLACK_PAWN: UTF8_BLACK_PAWN,
}

var PIECE_TO_UTF8 = map[PieceType]string{
    WHITE_KING: UTF8_WHITE_KING,
    WHITE_QUEEN: UTF8_WHITE_QUEEN,
    WHITE_ROOK: UTF8_WHITE_ROOK,
    WHITE_BISHOP: UTF8_WHITE_BISHOP,
    WHITE_KNIGHT: UTF8_WHITE_KNIGHT,
    WHITE_PAWN: UTF8_WHITE_PAWN,
    BLACK_KING: UTF8_BLACK_KING,
    BLACK_QUEEN: UTF8_BLACK_QUEEN,
    BLACK_ROOK: UTF8_BLACK_ROOK,
    BLACK_BISHOP: UTF8_BLACK_BISHOP,
    BLACK_KNIGHT: UTF8_BLACK_KNIGHT,
    BLACK_PAWN: UTF8_BLACK_PAWN,
}

var PIECE_TO_ALGEBRAIC = map[PieceType]string{
    WHITE_KING: ALGEBRAIC_WHITE_KING,
    WHITE_QUEEN: ALGEBRAIC_WHITE_QUEEN,
    WHITE_ROOK: ALGEBRAIC_WHITE_ROOK,
    WHITE_BISHOP: ALGEBRAIC_WHITE_BISHOP,
    WHITE_KNIGHT: ALGEBRAIC_WHITE_KNIGHT,
    WHITE_PAWN: ALGEBRAIC_WHITE_PAWN,
    BLACK_KING: ALGEBRAIC_BLACK_KING,
    BLACK_QUEEN: ALGEBRAIC_BLACK_QUEEN,
    BLACK_ROOK: ALGEBRAIC_BLACK_ROOK,
    BLACK_BISHOP: ALGEBRAIC_BLACK_BISHOP,
    BLACK_KNIGHT: ALGEBRAIC_BLACK_KNIGHT,
    BLACK_PAWN: ALGEBRAIC_BLACK_PAWN,
}

type CastlingType uint8
func (c CastlingType) is(otherType CastlingType) bool {
    return c == otherType
}

const (
    NO_CASTLING CastlingType = iota
    WHITE_CASTLING_SHORT
    WHITE_CASTLING_LONG
    BLACK_CASTLING_SHORT
    BLACK_CASTLING_LONG
)

const WHITE_CASTLING_SHORT_STRING = "0-0"
const BLACK_CASTLING_SHORT_STRING = "0-0"
const WHITE_CASTLING_LONG_STRING  = "0-0-0"
const BLACK_CASTLING_LONG_STRING  = "0-0-0"

var CASTLING_TO_STRING = map[CastlingType]string{
    WHITE_CASTLING_SHORT: WHITE_CASTLING_SHORT_STRING,
    WHITE_CASTLING_LONG:  WHITE_CASTLING_LONG_STRING,
    BLACK_CASTLING_SHORT: BLACK_CASTLING_SHORT_STRING,
    BLACK_CASTLING_LONG:  BLACK_CASTLING_LONG_STRING,
}

func (c CastlingType) String() string {
    return CASTLING_TO_STRING[c]
}

const A1 = 1 << 0
const B1 = 1 << 1
const C1 = 1 << 2
const D1 = 1 << 3
const E1 = 1 << 4
const F1 = 1 << 5
const G1 = 1 << 6
const H1 = 1 << 7
const A2 = 1 << 8
const B2 = 1 << 9
const C2 = 1 << 10
const D2 = 1 << 11
const E2 = 1 << 12
const F2 = 1 << 13
const G2 = 1 << 14
const H2 = 1 << 15
const A3 = 1 << 16
const B3 = 1 << 17
const C3 = 1 << 18
const D3 = 1 << 19
const E3 = 1 << 20
const F3 = 1 << 21
const G3 = 1 << 22
const H3 = 1 << 23
const A4 = 1 << 24
const B4 = 1 << 25
const C4 = 1 << 26
const D4 = 1 << 27
const E4 = 1 << 28
const F4 = 1 << 29
const G4 = 1 << 30
const H4 = 1 << 31
const A5 = 1 << 32
const B5 = 1 << 33
const C5 = 1 << 34
const D5 = 1 << 35
const E5 = 1 << 36
const F5 = 1 << 37
const G5 = 1 << 38
const H5 = 1 << 39
const A6 = 1 << 40
const B6 = 1 << 41
const C6 = 1 << 42
const D6 = 1 << 43
const E6 = 1 << 44
const F6 = 1 << 45
const G6 = 1 << 46
const H6 = 1 << 47
const A7 = 1 << 48
const B7 = 1 << 49
const C7 = 1 << 50
const D7 = 1 << 51
const E7 = 1 << 52
const F7 = 1 << 53
const G7 = 1 << 54
const H7 = 1 << 55
const A8 = 1 << 56
const B8 = 1 << 57
const C8 = 1 << 58
const D8 = 1 << 59
const E8 = 1 << 60
const F8 = 1 << 61
const G8 = 1 << 62
const H8 = 1 << 63


type PromotionKey struct {
    color Color
    uciPromotion string
}

var PROMOTION_TO_PIECE = map[PromotionKey]PieceType {
    PromotionKey{WHITE, "q"} : WHITE_QUEEN,
    PromotionKey{WHITE, "r"} : WHITE_ROOK,
    PromotionKey{WHITE, "b"} : WHITE_BISHOP,
    PromotionKey{WHITE, "n"} : WHITE_KNIGHT,
    PromotionKey{BLACK, "q"} : BLACK_QUEEN,
    PromotionKey{BLACK, "r"} : BLACK_ROOK,
    PromotionKey{BLACK, "b"} : BLACK_BISHOP,
    PromotionKey{BLACK, "n"} : BLACK_KNIGHT,
}
