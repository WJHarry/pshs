package main

import(
	"strings"
	"strconv"
	"github.com/gdamore/tcell"
)

func drawSearch(screen tcell.Screen) {
	searchStyle := tcell.StyleDefault.Foreground(tcell.ColorPink).Background(tcell.ColorDefault)
	width, _ := screen.Size()
	drawText(screen, 1, 0, searchStyle, strings.Repeat("-", width - 2))
	drawText(screen, 0, 1, searchStyle, "|")
	drawText(screen, width - 1, 1, searchStyle, "|")
	drawText(screen, 1, 2, searchStyle, strings.Repeat("-", width - 2))
}

func drawHistory(screen tcell.Screen, drawList []string, offset int) {
	defStyle := tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorPurple)

	for index, history := range drawList {
		drawText(screen, 0, index + 4, defStyle, strconv.Itoa(index + offset))
		drawText(screen, 5, index + 4, defStyle, history)
	}
}

func highLightLine(screen tcell.Screen, row int, text string, offset int) {
	highLightStyle := tcell.StyleDefault.Foreground(tcell.ColorBrown).Background(tcell.ColorYellow)
	drawText(screen, 0, row + 4, highLightStyle, strconv.Itoa(row + offset))
	drawText(screen, 5, row + 4, highLightStyle, text)
}

func cancelHighLightLine(screen tcell.Screen, row int, text string, offset int) {
	highLightStyle := tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorPurple)
	drawText(screen, 0, row + 4, highLightStyle, strconv.Itoa(row + offset))
	drawText(screen, 5, row + 4, highLightStyle, text)
}

func drawText(s tcell.Screen, x1, y1 int, style tcell.Style, text string) {
	row := y1
	col := x1
	for _, r := range []rune(text) {
		s.SetContent(col, row, r, nil, style)
		col++
	}
}
