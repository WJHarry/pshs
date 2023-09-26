package main

import(
	"strings"
	"bufio"
	"os"

	ps "github.com/bhendo/go-powershell"
	"github.com/bhendo/go-powershell/backend"
)

func findHistory() []string {

	historyList := make([]string, 10)

	env := checkPsOrCmd()

	if env == PS {
		back := &backend.Local{}
		shell, err := ps.New(back)
		if err != nil {
			panic(err)
		}
		defer shell.Exit()
		historyPath, _, err := shell.Execute("(Get-PSReadlineOption).HistorySavePath")
		if err != nil {
			panic(err)
		}

		historyPath = strings.Replace(historyPath, "\\", "/", -1)
		historyPath = strings.Replace(historyPath, "\r", "", -1)
		historyPath = strings.Replace(historyPath, "\n", "", -1)
		f, err := os.Open(historyPath)
		if err != nil {
			panic(err)
		}
		defer f.Close()

		scanner := bufio.NewScanner(f)

		for scanner.Scan() {
			historyList = append(historyList, scanner.Text())
		}
	} else {
		// TODO: create a new history file
	}
	return historyList
}
