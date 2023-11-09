package gui

import (
	"fmt"
	"os"
	"strings"

	"github.com/mxkrsv/etu-oop-2023/task3/matrix"
	"github.com/mxkrsv/etu-oop-2023/task3/numbers"

	"github.com/diamondburned/gotk4/pkg/gio/v2"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
)

type Application[n numbers.StdlibNumeric, N numbers.CustomNumeric[n, N]] struct {
	matrix matrix.Matrix[n, N]
	gtkApp *gtk.Application
}

func NewApplication[n numbers.StdlibNumeric, N numbers.CustomNumeric[n, N]]() Application[n, N] {
	a := Application[n, N]{
		matrix: matrix.Matrix[n, N]{},
		gtkApp: gtk.NewApplication(
			"com.github.mxkrsv.etu-oop-2023.task3.application.gui",
			gio.ApplicationFlagsNone,
		),
	}

	a.gtkApp.ConnectActivate(func() { a.activate() })

	return a
}

func (a Application[n, N]) activate() {
	window := gtk.NewApplicationWindow(a.gtkApp)
	window.SetTitle("Matrix test flight")
	textView := gtk.NewTextView()
	box1 := gtk.NewBox(gtk.OrientationVertical, 5)
	box1.Append(textView)
	window.SetChild(box1)
	window.SetDefaultSize(400, 300)

	box2 := gtk.NewBox(gtk.OrientationHorizontal, 5)

	btnReadMatrix := gtk.NewButtonWithLabel("Read")
	box2.Append(btnReadMatrix)

	btnTranspose := gtk.NewButtonWithLabel("Transpose")
	box2.Append(btnTranspose)

	btnDet := gtk.NewButtonWithLabel("Det")
	box2.Append(btnDet)

	btnRank := gtk.NewButtonWithLabel("Rank")
	box2.Append(btnRank)

	box1.Append(box2)

	box1.SetMarginTop(10)
	box1.SetMarginBottom(10)
	box1.SetMarginStart(10)
	box1.SetMarginEnd(10)

	statusLabel := gtk.NewLabel("")
	box1.Append(statusLabel)

	buffer := textView.Buffer()
	btnReadMatrix.ConnectClicked(func() {
		s := buffer.Text(buffer.IterAtOffset(0), buffer.EndIter(), false)
		err := a.matrix.Read(strings.NewReader(s))
		if err != nil {
			statusLabel.SetLabel(err.Error())
		}
		statusLabel.SetLabel("matrix read successfully")
	})

	btnTranspose.ConnectClicked(func() {
		err := a.matrix.Transpose()
		if err != nil {
			panic(err)
		}

		s, err := a.matrix.String()
		if err != nil {
			panic(err)
		}

		buffer.SetText(s)
		statusLabel.SetLabel("matrix transposed successfully")
	})

	btnDet.ConnectClicked(func() {
		det, err := a.matrix.Det()
		if err != nil {
			panic(err)
		}

		statusLabel.SetLabel(fmt.Sprintf("determinant: %v", det))
	})

	btnRank.ConnectClicked(func() {
		rank, err := a.matrix.Rank()
		if err != nil {
			panic(err)
		}

		statusLabel.SetLabel(fmt.Sprintf("rank: %v", rank))
	})

	window.Show()
}

func (a Application[n, N]) Run() {
	a.gtkApp.Run(os.Args)
}
