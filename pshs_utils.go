package main

import(
	"os"
)

const PS = "ps"
const CMD = "cmd"

func reverse(s []string) {
	for i := 0; i < len(s)/2; i++ {
		j := len(s) - i - 1
		s[i], s[j] = s[j], s[i]
	}
}

func filter(slice []string, f func(string) bool) []string {
	var n []string
	for _, e := range slice {
		if f(e) {
			n = append(n, e)
		}
	}
	return n
}

func checkPsOrCmd() string {
	wtSession := os.Getenv("WT_SESSION")
	if wtSession == "" {
		return CMD
	}
	return PS
}
