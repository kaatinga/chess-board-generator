package DrawChessBoard

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"os"
)

var img *image.RGBA
var chessGlyphs, letters image.Image

type piece struct {
	coordinates image.Point
	name        string
}

var WhiteKing piece
var WhiteQueen piece
var WhiteTower piece
var WhiteOfficer piece
var WhiteHorse piece
var WhitePawn piece
var BlackKing piece
var BlackQueen piece
var BlackTower piece
var BlackOfficer piece
var BlackHorse piece
var BlackPawn piece

var grey color.Color = color.RGBA{125, 125, 125, 255}
var lightGrey color.Color = color.RGBA{230, 230, 230, 255}
var boardSize = 451

func drawLine(x1, stepX, x2, y1, stepY, y2 int, colour *color.Color) {
	for x := x1; x < x2+1; x = x + stepX {
		for y := y1; y < y2+1; y = y + stepY {
			img.Set(x, y, *colour)
		}
	}
}

func PlacePiece(x, y, step int, piece2place piece) {
	draw.Draw(img, image.Rect(x, y, x+step, y+step), chessGlyphs, piece2place.coordinates, draw.Over)
}

// Init() initialises the chess board. Do it first of all
func Init() {

	// определяем размер картинки и инициализируем доску
	img = image.NewRGBA(image.Rect(0, 0, boardSize, boardSize))

	draw.Draw(img, img.Bounds(), image.Transparent, image.ZP, draw.Src)
	var step = 50
	drawLine(0, step, boardSize, 0, 1, boardSize, &grey)
	drawLine(0, 1, boardSize, 0, step, boardSize, &grey)

	for xHead := step << 1; xHead < boardSize+1; xHead = xHead + step<<1 {
		for yHead := step; yHead < boardSize+1; yHead = yHead + step<<1 {
			drawLine(xHead+1, 1, xHead+step-1, yHead+1, 1, yHead+step-1, &lightGrey)
		}
	}
	for yHead := step << 1; yHead < boardSize+1; yHead = yHead + step<<1 {
		for xHead := step; xHead < boardSize+1; xHead = xHead + step<<1 {
			drawLine(xHead+1, 1, xHead+step-1, yHead+1, 1, yHead+step-1, &lightGrey)
		}
	}

	// Загружаем фигурки
	chessGlyphsSrc, err := os.Open("chess_glyphs.png")
	if err != nil {
		log.Println(err)
	}
	defer chessGlyphsSrc.Close()

	// Загружаем буковки и цифорки
	lettersSrc, err := os.Open("letters.png")
	if err != nil {
		log.Println(err)
	}
	defer lettersSrc.Close()

	chessGlyphs, _, _ = image.Decode(chessGlyphsSrc)
	letters, _, _ = image.Decode(lettersSrc)

	// Рисуем верхний ряд
	draw.Draw(img, image.Rect(step, 0, boardSize, step), letters, image.Pt(0, 0), draw.Over)

	// Рисуем левый ряд
	for yHead := step; yHead < boardSize+1; yHead = yHead + step {
		draw.Draw(img, image.Rect(0, yHead, step, yHead+step), letters, image.Pt(yHead-step, step), draw.Over)
	}

	// Формируем массив фигурок для работы
	WhiteKing = piece{image.Pt(0, 50), "White King"}
	WhiteQueen = piece{image.Pt(50, 50), "White Queen"}
	WhiteTower = piece{image.Pt(100, 50), "White Tower"}
	WhiteOfficer = piece{image.Pt(150, 50), "White Officer"}
	WhiteHorse = piece{image.Pt(200, 50), "White Horse"}
	WhitePawn = piece{image.Pt(250, 50), "White Pawn"}
	BlackKing = piece{image.Pt(0, 0), "Black King"}
	BlackQueen = piece{image.Pt(50, 0), "Black Queen"}
	BlackTower = piece{image.Pt(100, 0), "Black Tower"}
	BlackOfficer = piece{image.Pt(150, 0), "Black Officer"}
	BlackHorse = piece{image.Pt(200, 0), "Black Horse"}
	BlackPawn = piece{image.Pt(250, 0), "Black Pawn"}
}

// Save() saves a PNG file, point out a file name
func Save(fileName string) {

	file, err := os.Create(fileName)
	if err != nil {
		log.Println(err)
	}
	defer file.Close()

	err = png.Encode(file, img)
	if err != nil {
		log.Println(err)
	}
}
