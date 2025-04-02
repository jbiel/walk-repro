package main

import (
	"github.com/tailscale/walk"
	. "github.com/tailscale/walk/declarative"
)

func main() {
	var testWindow *walk.MainWindow
	var usernameEdit, passwordEdit *walk.LineEdit

	app, err := walk.InitApp()
	if err != nil {
		panic(err)
	}

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
						Text: "Greetings from tailscale/walk!",
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
						Text:    "Cancel",
						Default: false,
						OnClicked: func() {
							testWindow.Close()
						},
					},
					PushButton{
						Text:    "Continue",
						Default: true,
						OnClicked: func() {
							testWindow.Close()
						},
					},
				},
			},
		},
	}.Create()

	app.Run()
}
