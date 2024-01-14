package typer

import (
	"github.com/go-vgo/robotgo"
)

func Typer(str string, amount int) {
	for i := 0; i < amount; i++ {
		robotgo.TypeStr(str)
		robotgo.KeySleep = 10
		robotgo.KeyTap("enter")
	}
}
