package menu

import (
	"aerith/renderer"
	"aerith/rmidi"
	"aerith/screen"
	"fmt"

	rg "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type ButtonCallback func(b *IButton)

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
		Callback: func(b *IButton) {
			screen.Mode = "render"
			renderer.Init()
		},
	})

	buttons = append(buttons, IButton{
		Rectangle: rl.NewRectangle(24, float32(screen.Height)-30-24, 120, 30),
		Text:      "Quit",
		Callback: func(b *IButton) {
			screen.ShouldClose = true
		},
	})

	buttons = append(buttons, IButton{
		Rectangle: rl.NewRectangle(24, float32(screen.Height)-88, 120, 30),
		Text:      "Enable FPS",
		Callback: func(b *IButton) {
			renderer.EnableFPS = !renderer.EnableFPS

			if renderer.EnableFPS {
				b.Text = "Disable FPS"
			} else {
				b.Text = "Enable FPS"
			}
		},
	})

	buttons = append(buttons, IButton{
		Rectangle: rl.NewRectangle(24, 30+24+24, 120, 30),
		Text:      "Load MIDI",
		Callback: func(b *IButton) {
			// Load MIDI file
			// TODO Open file dialog

			file := "test.mid"
			rmidi.LoadMIDI("file")
			fmt.Println("Loaded " + file)
		},
	})

	buttons = append(buttons, IButton{
		Rectangle: rl.NewRectangle(24, 30+30+24+24, 120, 30),
		Text:      "Unload MIDI",
		Callback: func(b *IButton) {
			// Unload MIDI file
			rmidi.UnloadMIDI()
			fmt.Println("Unloaded MIDI")
		},
	})
}

func Update(dt float32) {

}

func Draw() {
	style := uint(rg.GetStyle(rg.DEFAULT, rg.BACKGROUND_COLOR))
	bgColor := rl.GetColor(style)
	rl.ClearBackground(bgColor)

	// Draw every menu button

	for i := range buttons {
		button := &buttons[i]
		if rg.Button(button.Rectangle, button.Text) {
			button.Callback(button)
		}
	}
}
