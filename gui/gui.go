package gui

import (
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/stenstromen/textstorm/typer"
)

func Gui(app *fyne.App) {
	w := (*app).NewWindow("TextStorm")
	w.Resize(fyne.NewSize(500, 500))
	w.CenterOnScreen()

	message := widget.NewMultiLineEntry()
	message.SetMinRowsVisible(10)
	message.SetPlaceHolder("Enter your message here")

	amount := widget.NewEntry()
	amount.SetPlaceHolder("# of times to send message")

	countdownLabel := widget.NewLabel("")

	requestPermission := widget.NewButton("Request Keyboard Permission", func() {
		keyboardPermission()
	})

	padding := container.NewVBox(
		widget.NewLabel(""),
		widget.NewLabel(""),
		widget.NewLabel(""),
		widget.NewLabel(""),
	)

	send := widget.NewButton("Send", func() {
		textType(message.Text, amount.Text, countdownLabel)
	})

	w.SetContent(container.NewVBox(
		countdownLabel,
		message,
		amount,
		requestPermission,
		padding,
		send,
	))

	w.Show()
}

func textType(text, amount string, countdownLabel *widget.Label) {
	amountInt, err := strconv.Atoi(amount)
	if err != nil {
		fyne.LogError("Error converting amount to integer", err)
		return
	}

	go func() {
		for i := 3; i > 0; i-- {
			countdownLabel.SetText("TextStorm in: " + strconv.Itoa(i))
			time.Sleep(1 * time.Second)
		}
		countdownLabel.SetText("")
		typer.Typer(text, amountInt)
	}()
}

func keyboardPermission() {
	typer.Typer("test", 1)
}
