package main

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

func main() {
	var testWindow *walk.MainWindow
	var usernameEdit, passwordEdit *walk.LineEdit

	MainWindow{
		AssignTo: &testWindow,
		Title:    "Test title",
		Size:     Size{Width: 500, Height: 300},
		Layout:   VBox{},
		Children: []Widget{
			Composite{
				Layout: HBox{},
				Children: []Widget{
					Label{
						Text: "Greetings from lxn/walk!",
						Font: Font{PointSize: 14, Bold: true},
					},
					HSpacer{},
				},
			},
			Label{
				Text: "Enter login",
			},
			VSpacer{Size: 10},

			// Credentials Section
			Composite{
				Layout: Grid{Columns: 2},
				Children: []Widget{
					Label{Text: "Username:"},
					LineEdit{AssignTo: &usernameEdit},
					Label{Text: "Password:"},
					LineEdit{AssignTo: &passwordEdit, PasswordMode: true},
				},
			},
			VSpacer{},

			// Button Section
			Composite{
				Layout: HBox{},
				Children: []Widget{
					HSpacer{},
					PushButton{
						Text: "Cancel",
						OnClicked: func() {
							testWindow.Close()
						},
					},
					PushButton{
						Text: "Continue",
						OnClicked: func() {
							testWindow.Close()
						},
					},
				},
			},
		},
	}.Run()
}
