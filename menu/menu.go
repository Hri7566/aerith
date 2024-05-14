package menu

import (
	"aerith/screen"

	rg "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type ButtonCallback func()

type IButton struct {
	Rectangle rl.Rectangle
	Text      string
	Callback  ButtonCallback
}

var buttons []IButton

func Init() {
	buttons = append(buttons, IButton{
		Rectangle: rl.NewRectangle(24, 24, 120, 30),
		Text:      "Render",
		Callback: func() {
			screen.Mode = "render"
		},
	})
}

func Update(dt float32) {

}

func Draw() {
	style := uint(rg.GetStyle(rg.DEFAULT, rg.BACKGROUND_COLOR))
	bgColor := rl.GetColor(style)
	rl.ClearBackground(bgColor)

	for _, button := range buttons {
		if rg.Button(button.Rectangle, button.Text) {
			button.Callback()
		}
	}
}
