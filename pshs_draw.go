package main

import(
	"strings"
	"strconv"
	"github.com/gdamore/tcell"
)

var searchStyle = tcell.StyleDefault.Foreground(tcell.ColorRed).Background(tcell.ColorDefault)
var modeStype = tcell.StyleDefault.Foreground(tcell.ColorYellow).Background(tcell.ColorGray)

func drawSearch(screen tcell.Screen) {
	width, _ := screen.Size()
	drawText(screen, 1, 0, searchStyle, strings.Repeat("-", width - 2))
	drawText(screen, 0, 1, searchStyle, "|")
	drawText(screen, width - 1, 1, searchStyle, "|")
	drawText(screen, 1, 2, searchStyle, strings.Repeat("-", width - 2))

	if mode == 1 {
		drawText(screen, 1, 3, modeStype, "Regexp mode")
		drawText(screen, 13, 3, searchStyle, "Ctrl-R: Normal mode  Down/Ctrl-N: Next  Up/Ctrl-P: Previous  Ctrl-C: Quit  Enter: Select and execute")
	} else {
		drawText(screen, 1, 3, modeStype, "Normal mode")
		drawText(screen, 13, 3, searchStyle, "Ctrl-R: Regexp mode  Down/Ctrl-N: Next  Up/Ctrl-P: Previous  Ctrl-C: Quit  Enter: Select and execute")
	}
}

func drawHistory(screen tcell.Screen, drawList []string, offset int) {
	defStyle := tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorPurple)

	for index, history := range drawList {
		drawText(screen, 0, index + 5, defStyle, strconv.Itoa(index + offset))
		drawText(screen, 5, index + 5, defStyle, history)
	}
}

func highLightLine(screen tcell.Screen, row int, text string, offset int) {
	highLightStyle := tcell.StyleDefault.Foreground(tcell.ColorBrown).Background(tcell.ColorYellow)
	drawText(screen, 0, row + 5, highLightStyle, strconv.Itoa(row + offset))
	drawText(screen, 5, row + 5, highLightStyle, text)
}

func cancelHighLightLine(screen tcell.Screen, row int, text string, offset int) {
	highLightStyle := tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorPurple)
	drawText(screen, 0, row + 5, highLightStyle, strconv.Itoa(row + offset))
	drawText(screen, 5, row + 5, highLightStyle, text)
}

func drawText(s tcell.Screen, x1, y1 int, style tcell.Style, text string) {
	row := y1
	col := x1
	for _, r := range []rune(text) {
		s.SetContent(col, row, r, nil, style)
		col++
	}
}
