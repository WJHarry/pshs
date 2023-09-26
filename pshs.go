package main

func main() {
	historyList := findHistory()
	reverse(historyList)
	if len(historyList) > 999 {
		historyList = historyList[:999]
	}
	showHistory(historyList)
}
