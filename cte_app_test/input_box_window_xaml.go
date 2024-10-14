package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type InputBoxWindow struct {
	window      fyne.Window
	txtValor    *widget.Entry
	btnOk       *widget.Button
	btnCancelar *widget.Button
}

func NewInputBoxWindow() *InputBoxWindow {
	w := &InputBoxWindow{}
	w.createUI()
	return w
}

func (w *InputBoxWindow) createUI() {
	myApp := app.New()
	w.window = myApp.NewWindow("Input Box")

	w.txtValor = widget.NewEntry()
	w.btnOk = widget.NewButton("OK", w.btnOkClick)
	w.btnCancelar = widget.NewButton("Cancel", w.btnCancelarClick)

	content := container.NewVBox(
		w.txtValor,
		container.NewHBox(w.btnOk, w.btnCancelar),
	)

	w.window.SetContent(content)
	w.window.Resize(fyne.NewSize(300, 100))

	w.txtValor.OnSubmitted = func(s string) {
		w.btnOk.SetFocus()
	}

	w.window.Canvas().SetOnTypedKey(func(ke *fyne.KeyEvent) {
		if ke.Name == fyne.KeyEnter {
			w.btnOk.SetFocus()
		}
	})

	w.window.SetOnShown(func() {
		w.txtValor.SetFocus()
	})
}

func (w *InputBoxWindow) btnCancelarClick() {
	w.window.Close()
}

func (w *InputBoxWindow) btnOkClick() {
	w.window.Close()
}

func (w *InputBoxWindow) Show() {
	w.window.ShowAndRun()
}

func main() {
	inputBox := NewInputBoxWindow()
	inputBox.Show()
}
