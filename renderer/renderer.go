package renderer

import (
	"aerith/screen"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var showFPS bool = false

func Update(dt float32) {
	if rl.IsKeyPressed(rl.KeyF) {
		showFPS = !showFPS
	}

	if rl.IsKeyPressed(rl.KeyQ) {
		screen.Mode = "menu"
	}
}

func Draw() {
	rl.ClearBackground(rl.Black)

	if showFPS {
		rl.DrawFPS(0, screen.Height-26)
	}
}
