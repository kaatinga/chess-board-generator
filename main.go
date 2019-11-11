package main

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

func placePiece(x, y, step int, piece image.Point) {
	draw.Draw(img, image.Rect(x, y, x+step, y+step), chessGlyphs, piece, draw.Over)
}

func main() {

	file, err := os.Create("someimage.png")
	if err != nil {
		log.Println(err)
	}
	defer file.Close()

	// определяем размер картинки и инициализируем холст
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

	// Верхний ряд
	draw.Draw(img, image.Rect(step, 0, boardSize, step), letters, image.Pt(0, 0), draw.Over)

	// Левый ряд
	for yHead := step; yHead < boardSize+1; yHead = yHead + step {
		draw.Draw(img, image.Rect(0, yHead, step, yHead+step), letters, image.Pt(yHead-step, step), draw.Over)
	}

	glyphsArray := map[string]image.Point{
		"WhiteKing":    image.Pt(0, 50),
		"WhiteQueen":   image.Pt(50, 50),
		"WhiteTower":   image.Pt(100, 50),
		"WhiteOfficer": image.Pt(150, 50),
		"WhiteHorse":   image.Pt(200, 50),
		"WhitePawn":    image.Pt(250, 50),
		"BlackKing":    image.Pt(0, 0),
		"BlackQueen":   image.Pt(50, 0),
		"BlackTower":   image.Pt(100, 0),
		"BlackOfficer": image.Pt(150, 0),
		"BlackHorse":   image.Pt(200, 0),
		"BlackPawn":    image.Pt(250, 0),
	}

	placePiece(50, 50, step, glyphsArray["WhiteTower"])
	placePiece(100, 50, step, glyphsArray["WhitePawn"])
	placePiece(250, 50, step, glyphsArray["BlackTower"])
	placePiece(100, 250, step, glyphsArray["WhiteHorse"])
	placePiece(150, 250, step, glyphsArray["WhiteKing"])
	placePiece(350, 50, step, glyphsArray["BlackKing"])
	placePiece(300, 300, step, glyphsArray["BlackQueen"])

	png.Encode(file, img)
}
