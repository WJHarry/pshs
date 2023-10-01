package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/gdamore/tcell"
	"github.com/gdamore/tcell/encoding"
)

// 0: normal 1: regxep
var mode int8 = 0

func showHistory(fullHistoryList []string) {
	encoding.Register()
	screen, err := tcell.NewScreen()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	screen.Init()
	screen.EnableMouse()

	screen.Clear()
	drawSearch(screen)
	screen.Show()

	for i, history := range fullHistoryList {
		fullHistoryList[i] = history + "\t"
	}

	historyList := fullHistoryList
	drawHistory(screen, historyList, 0)
	highLightLine(screen, 0, historyList[0], 0)

	offset := 0
	currentLine := 0

	searchText := ""
	selectHistory := ""

	for {
		screen.Show()
		ev := screen.PollEvent()

		switch ev := ev.(type) {
		case *tcell.EventResize:
			screen.Sync()
			screen.Clear()
			drawSearch(screen)
			drawHistory(screen, historyList[offset:], offset)
			highLightLine(screen, currentLine, historyList[currentLine], offset)
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyDown || ev.Key() == tcell.KeyCtrlN {
				if currentLine == len(historyList) - 1 {
					continue
				}
				_, height := screen.Size()
				if currentLine - offset == height - 6 {
					offset++
					screen.Clear()
					drawSearch(screen)
					drawHistory(screen, historyList[offset:], offset)
				}
				cancelHighLightLine(screen, currentLine - offset, historyList[currentLine], offset)
				currentLine++
				highLightLine(screen, currentLine - offset, historyList[currentLine], offset)
				screen.Size()
			} else if ev.Key() == tcell.KeyUp || ev.Key() == tcell.KeyCtrlP {
				if currentLine == 0 {
					continue
				}
				if currentLine == offset {
					offset--
					screen.Clear()
					drawSearch(screen)
					drawHistory(screen, historyList[offset:], offset)
				}
				cancelHighLightLine(screen, currentLine - offset, historyList[currentLine], offset)
				currentLine--
				highLightLine(screen, currentLine - offset, historyList[currentLine], offset)
			} else if ev.Key() == tcell.KeyEnter {
				selectHistory = historyList[currentLine]
				goto exit
			} else if ev.Key() == tcell.KeyCtrlC {
				goto exit
			} else if ev.Key() == tcell.KeyRune {
				if len(searchText) > 100 {
					continue
				}
				searchText += string(ev.Rune())
				historyList, _ := searchHistory(searchText, fullHistoryList)

				screen.Clear()
				offset = 0
				currentLine = 0
				drawSearch(screen)
				drawText(screen, 1, 1, searchStyle, searchText)

				if len(historyList) > 0 {
					drawHistory(screen, historyList[offset:], offset)
					highLightLine(screen, currentLine, historyList[currentLine], offset)
				}
			} else if ev.Key() == tcell.KeyBS {
				if len(searchText) == 0 {
					continue
				}
				if len(searchText) == 1 {
					searchText = ""
					historyList = fullHistoryList
				} else {
					searchText = searchText[:len(searchText)-1]
					historyList, _ = searchHistory(searchText, fullHistoryList)
				}
				screen.Clear()
				offset = 0
				currentLine = 0
				drawSearch(screen)
				drawText(screen, 1, 1, searchStyle, searchText)
				if len(historyList) > 0 {
					drawHistory(screen, historyList[offset:], offset)
					highLightLine(screen, currentLine, historyList[currentLine], offset)
				}
			} else if ev.Key() == tcell.KeyCtrlR {
				screen.Clear()
				if mode == 0 {
					mode = 1
				} else {
					mode = 0
				}
				historyList, _ := searchHistory(searchText, fullHistoryList)
				offset = 0
				currentLine = 0
				drawSearch(screen)
				drawText(screen, 1, 1, searchStyle, searchText)
				if len(historyList) > 0 {
					drawHistory(screen, historyList[offset:], offset)
					highLightLine(screen, currentLine, historyList[currentLine], offset)
				}
			}
		}
	}
exit:
	screen.SetStyle(tcell.StyleDefault.Foreground(tcell.ColorDefault).Background(tcell.ColorDefault))
	screen.Fini()

	fmt.Println(selectHistory)
}


func searchHistory(searchText string, fullHistoryList []string) ([]string, bool) {
	if mode == 0 {
		return filter(fullHistoryList, func(history string) bool {
			return strings.Contains(history, searchText)
		}), true
	} else {
		re, err := regexp.Compile(searchText)
		if err != nil {
			return nil, false
		}

		var result []string
		for _, history := range fullHistoryList {
			if re.MatchString(history) {
				result = append(result, history)
			}
		}
		return result, true
	}
}
