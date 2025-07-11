package messages

import (
	"fmt"
	"os"

	"github.com/redseverity/forcepath/utils/text"
)

func Success(msg string) {
	fmt.Print(text.Bold, text.Green)
	fmt.Print("{ ", msg, " }", text.Reset, "\n\n")
}

func Error(msg string) {
	fmt.Fprint(os.Stderr, text.Bold, text.Red)
	fmt.Fprint(os.Stderr, "{ ", msg, " }", text.Reset, "\n\n")
}

func Warning(msg string) {
	fmt.Print(text.Yellow, msg, text.Reset, "\n\n")
}

func SuccessInputDetail(label string, param string) {
	fmt.Print(text.Bold, text.Green, "$ ")
	fmt.Print(text.Cyan, label, " ", text.Reset, param, "\n")
}

func ErrorInputDetail(label string, param string) {
	fmt.Fprint(os.Stderr, text.Bold, text.Red, "$ ")
	fmt.Fprint(os.Stderr, text.Cyan, label, " ", text.Reset, param, "\n\n")
}

func SuccessRequest(label string, param string) {
	fmt.Print(text.Bold, text.Green, "● ")
	fmt.Print(text.Cyan, label, " ", text.Reset, param, "\n")
}

func ErrorRequest(label string, param string) {
	fmt.Fprint(os.Stderr, text.Bold, text.Red, "● ")
	fmt.Fprint(os.Stderr, text.Cyan, label, " ", text.Reset, param, "\n")
}

func Exit1() {
	fmt.Print(text.Bold, text.Red)
	os.Exit(1)
}

func Exit0() {
	fmt.Print(text.Bold, text.Red)
	os.Exit(0)
}
